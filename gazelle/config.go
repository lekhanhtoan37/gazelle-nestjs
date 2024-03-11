package nestjs

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
	"github.com/bazelbuild/buildtools/labels"
)

type NestjsConfig struct {
	Enabled          bool
	RootTsConfigFile string
	PackageFile      string
	NpmDependencies  struct {
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}
	LookupTypes            bool
	ImportAliases          []struct{ From, To string }
	ImportAliasPattern     *regexp.Regexp
	Visibility             Visibility
	CollectBarrels         bool
	CollectWebAssets       bool
	CollectAllAssets       bool
	CollectedAssets        map[string]bool
	CollectAll             bool
	CollectAllRoot         string
	CollectAllSources      map[string]bool
	Fix                    bool
	WebAssetSuffixes       map[string]bool
	Quiet                  bool
	Verbose                bool
	DefaultNpmLabel        string
	JestConfigRelativePath string
	JestTestsPerShard      int
	JestSize               string

	// Nestjs
	RootPkg             string
	Root                string
	CollectWithTsConfig bool
	PackageByDir        map[string]*Project
	SourcePerProject    map[string]map[string]bool
	Monorepo            bool
	NestCliPath         string
	PackageJSONPath     string
	IsNestjs            bool
	CompilerOptions     *CompilerOptions
	Transpiler          NestTranspiler
	RootTsConfig        *TsConfig
	PnpmLockFilePath    string

	// Jest
	jestConfigPath string
	jestConfig     *JestConfig

	IsIgnoreE2E bool
}

type Visibility struct {
	Labels []string
}

func (v *Visibility) String() string {
	return fmt.Sprintf("%v", v.Labels)
}

func (v *Visibility) Set(value string) error {
	v.Labels = append(v.Labels, value)
	return nil
}

type NestjsConfigs map[string]*NestjsConfig

const DefaultPackageJsonPath = "package.json"

func NewNestjsConfig() *NestjsConfig {
	root := ""
	return &NestjsConfig{
		Enabled:     true,
		PackageFile: DefaultPackageJsonPath,
		NpmDependencies: struct {
			Dependencies    map[string]string "json:\"dependencies\""
			DevDependencies map[string]string "json:\"devDependencies\""
		}{
			Dependencies:    make(map[string]string),
			DevDependencies: make(map[string]string),
		},
		LookupTypes:        true,
		ImportAliases:      []struct{ From, To string }{},
		ImportAliasPattern: regexp.MustCompile("$^"),
		Visibility: Visibility{
			Labels: []string{},
		},
		CollectBarrels:         false,
		CollectWebAssets:       false,
		CollectAllAssets:       false,
		CollectedAssets:        make(map[string]bool),
		CollectAll:             false,
		CollectAllRoot:         "",
		CollectAllSources:      make(map[string]bool),
		Fix:                    false,
		RootPkg:                root,
		Root:                   root,
		WebAssetSuffixes:       make(map[string]bool),
		Quiet:                  false,
		Verbose:                false,
		DefaultNpmLabel:        "//:node_modules",
		JestTestsPerShard:      -1,
		JestConfigRelativePath: "",

		// New
		NestCliPath:         fmt.Sprintf("%v/nest-cli.json", root),
		CollectWithTsConfig: true,
		PackageByDir:        make(map[string]*Project),
		IsNestjs:            false,
		SourcePerProject:    make(map[string]map[string]bool),
		RootTsConfigFile:    "tsconfig.json",
		PnpmLockFilePath:    "pnpm-lock.yaml",

		IsIgnoreE2E: true,
	}
}

func newJsConfigsWithRootConfig() NestjsConfigs {
	rootConfig := NewNestjsConfig()
	rootConfig.CollectedAssets = make(map[string]bool)
	return NestjsConfigs{
		"": rootConfig,
	}
}

func (nestjsConfig *NestjsConfig) parseRootTsConfig() {
	if nestjsConfig.RootTsConfigFile == "" {
		return
	}

	tsConfigPath := path.Join(nestjsConfig.Root, nestjsConfig.RootTsConfigFile)
	data, err := os.ReadFile(tsConfigPath)
	if err != nil {
		log.Printf(Err("Read tsconfig for project failed: %v, err: %v \n", nestjsConfig.Root, err))
		return
	}

	var tsConfig *TsConfig
	err = json.Unmarshal(data, &tsConfig)
	if err != nil {
		log.Printf(Err("Parse tsconfig for project failed: %v, err: %v \n", nestjsConfig.Root, err))
		return
	}

	nestjsConfig.RootTsConfig = tsConfig
}

func (nestjsConfig *NestjsConfig) parseJestConfig() {
	// jest in package.json
	if nestjsConfig.jestConfigPath == nestjsConfig.PackageJSONPath {
		data, err := os.ReadFile(nestjsConfig.jestConfigPath)
		if err != nil {
			log.Printf(Err("failed to open %s: %v\n", nestjsConfig.jestConfigPath, err))
			return
		}

		var wrappedJestConfig *struct {
			Jest JestConfig "json:\"jest\""
		}
		err = json.Unmarshal(data, &wrappedJestConfig)
		if err != nil {
			log.Printf(Err("failed to parse %s: %v\n", nestjsConfig.jestConfigPath, err))
			return
		}
		nestjsConfig.jestConfig = &wrappedJestConfig.Jest
	} else {
		// TODO: Support jest.config.js OR jest.config.ts
	}
}

func (nestjsConfig *NestjsConfig) parsePackageJSON(
	bazelGazelleConfig *config.Config,
	ruleFile *rule.File,
	directive rule.Directive,
) {
	values := strings.Split(directive.Value, " ")
	if len(values) != 2 {
		log.Fatalf(
			Err(
				"failed to read directive %s: %s, expected 2 values",
				directive.Key,
				directive.Value,
			),
		)
	}
	nestjsConfig.PackageFile = values[0]
	npmLabel := values[1]
	if strings.HasPrefix(npmLabel, ":") {
		npmLabel = labels.ParseRelative(npmLabel, ruleFile.Pkg).Format()
	}
	if !strings.HasSuffix(npmLabel, ":") && !strings.HasSuffix(npmLabel, "/") {
		npmLabel += "/"
	}

	rootDir := bazelGazelleConfig.ReadBuildFilesDir
	if bazelGazelleConfig.ReadBuildFilesDir == "" {
		rootDir = path.Join(bazelGazelleConfig.RepoRoot, ruleFile.Pkg)
	}

	// Read package.json
	packageJSONPath := path.Join(rootDir, nestjsConfig.PackageFile)
	data, err := os.ReadFile(packageJSONPath)
	if err != nil {
		log.Fatalf(Err("failed to open %s: %v", directive.Value, err))
	}

	// Read dependencies from file
	newDeps := struct {
		Dependencies    map[string]string "json:\"dependencies\""
		DevDependencies map[string]string "json:\"devDependencies\""
	}{
		Dependencies:    make(map[string]string),
		DevDependencies: make(map[string]string),
	}
	if err := json.Unmarshal(data, &newDeps); err != nil {
		log.Fatalf(Err("failed to parse %s: %v", directive.Value, err))
	}

	// Store npmLabel in dependencies
	for k := range newDeps.Dependencies {
		nestjsConfig.NpmDependencies.Dependencies[k] = npmLabel
	}
	for k := range newDeps.DevDependencies {
		nestjsConfig.NpmDependencies.DevDependencies[k] = npmLabel
	}

	nestjsConfig.PackageJSONPath = packageJSONPath
	nestjsConfig.jestConfigPath = nestjsConfig.PackageJSONPath
	nestjsConfig.PnpmLockFilePath = path.Join(rootDir, nestjsConfig.PnpmLockFilePath)
}

func (nestjsConfig *NestjsConfig) parseNestCliJSON(
	bazelGazelleConfig config.Config,
	ruleFile *rule.File,
) {
	// log.Printf("parseNestCliJSON: %+v", ruleFile)
	var root string
	if bazelGazelleConfig.ReadBuildFilesDir != "" {
		root = path.Join(bazelGazelleConfig.ReadBuildFilesDir, ruleFile.Pkg)
		nestjsConfig.Root = root
	} else {
		root = path.Join(bazelGazelleConfig.RepoRoot, ruleFile.Pkg)
		nestjsConfig.Root = root
	}

	data, err := os.ReadFile(path.Join(root, nestjsConfig.NestCliPath))
	if err != nil {
		log.Print(Warn("Read nest-cli.json failed: ", err))
		return
	}

	var nestCliConfig *NestCliConfig
	err = json.Unmarshal(data, &nestCliConfig)
	if err != nil {
		log.Print(Warn("Parse nest-cli failed: ", err))
		return
	}

	nestjsConfig.Monorepo = nestCliConfig.Monorepo

	if nestCliConfig.Projects == nil {
		log.Fatal("Parse nest-cli failed: projects not found")
	}

	for projectName, project := range nestCliConfig.Projects {
		// nestjs structure of monorepo
		projectTypeDir := "apps"
		visibility := Visibility{Labels: []string{}}
		if project.Type == "library" {
			projectTypeDir = "libs"
			visibility = Visibility{Labels: []string{"//visibility:public"}}
		} else if project.Type == "application" {
			projectTypeDir = "apps"
		}

		dirPath := path.Join(nestjsConfig.RootPkg, projectTypeDir, projectName)

		// Set project name + relative path
		nestCliConfig.Projects[projectName].Name = projectName
		nestCliConfig.Projects[projectName].Rel = dirPath

		if nestCliConfig.Projects[projectName].CompilerOptions.TsConfigPath != "" {
			nestCliConfig.Projects[projectName].TsConfigRel = nestCliConfig.Projects[projectName].CompilerOptions.TsConfigPath
			nestCliConfig.Projects[projectName].TsConfigPath = path.Join(
				nestjsConfig.RootPkg,
				nestCliConfig.Projects[projectName].TsConfigRel,
			)
			nestCliConfig.Projects[projectName].ParseTsConfig(root)
		}

		// Set project in both root and child
		nestjsConfig.PackageByDir[dirPath] = project
		if nestjsConfig.SourcePerProject[dirPath] == nil {
			nestjsConfig.SourcePerProject[dirPath] = make(map[string]bool)
		}
		nestjsConfig.Visibility = visibility
	}

	nestjsConfig.NestCliPath = path.Join(root, nestjsConfig.NestCliPath)
	nestjsConfig.CompilerOptions = nestCliConfig.CompilerOptions
	nestjsConfig.IsNestjs = true

	// Transpiler
	if nestjsConfig.CompilerOptions.Webpack {
		nestjsConfig.Transpiler = webpack

		// TODO: Support webpack
		log.Fatalf("Webpack is not supported yet\n")
	} else {
		if _, err := os.Stat(path.Join(nestjsConfig.Root, ".swcrc")); err == nil {
			nestjsConfig.Transpiler = swc
		} else {
			nestjsConfig.Transpiler = tsc
		}
	}
}

// RegisterFlags registers command-line flags used by the extension. This
// method is called once with the root configuration when Gazelle
// starts. RegisterFlags may set an initial values in Config.Exts. When flags
// are set, they should modify these values.
func (lang *NestJS) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
	c.Exts[languageName] = newJsConfigsWithRootConfig()
}

// CheckFlags validates the configuration after command line flags are parsed.
// This is called once with the root configuration when Gazelle starts.
// CheckFlags may set default values in flags or make implied changes.
func (lang *NestJS) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	return nil
}

// KnownDirectives returns a list of directive keys that this Configurer can
// interpret. Gazelle prints errors for directives that are not recognized by
// any Configurer.
func (*NestJS) KnownDirectives() []string {
	return []string{
		"js_extension",
		"js_root",
		"js_lookup_types",
		"js_fix",
		"js_package_file",
		"js_import_alias",
		"js_visibility",
		"js_collect_barrels",
		"js_aggregate_modules",
		"js_collect_web_assets",
		"js_aggregate_web_assets",
		"js_collect_all_assets",
		"js_aggregate_all_assets",
		"js_collect_all",
		"js_jest_test_per_shard",
		"js_jest_size",
		"js_jest_config",
		"js_web_asset",
		"js_quiet",
		"js_verbose",
		"js_default_npm_label",
	}
}

// ParentForPackage returns the parent Config for the given Bazel package.
func (c *NestjsConfigs) ParentForPackage(pkg string) *NestjsConfig {
	dir := filepath.Dir(pkg)
	if dir == "." {
		dir = ""
	}
	parent := (map[string]*NestjsConfig)(*c)[dir]
	return parent
}

// Configure modifies the configuration using directives and other information
// extracted from a build file. Configure is called in each directory.
//
// c is the configuration for the current directory. It starts out as a copy
// of the configuration for the parent directory.
//
// rel is the slash-separated relative path from the repository root to
// the current directory. It is "" for the root directory itself.
//
// f is the build file for the current directory or nil if there is no
// existing build file.
func (*NestJS) Configure(c *config.Config, rel string, f *rule.File) {
	// rel := _rel
	// if _rel == "" {
	// 	rel = "."
	// }
	// Create the root config.
	if _, exists := c.Exts[languageName]; !exists {
		c.Exts[languageName] = newJsConfigsWithRootConfig()
	}

	nestjsConfigs := c.Exts[languageName].(NestjsConfigs)

	nestjsConfig, exists := nestjsConfigs[rel]
	if !exists {
		parent := nestjsConfigs.ParentForPackage(rel)
		nestjsConfig = parent.NewChild()
		nestjsConfigs[rel] = nestjsConfig
	}

	// Read directives from existing file
	if f == nil {
		return
	}

	for _, directive := range f.Directives {

		switch directive.Key {

		case "js_extension":
			switch directive.Value {
			case "enabled":
				nestjsConfig.Enabled = true
			case "disabled":
				nestjsConfig.Enabled = false
			default:
				log.Fatalf(
					Err(
						"failed to read directive %s: %s, only \"enabled\", and \"disabled\" are valid",
						directive.Key,
						directive.Value,
					),
				)
			}

		case "js_lookup_types":
			nestjsConfig.LookupTypes = readBoolDirective(directive)

		case "js_fix":
			nestjsConfig.Fix = readBoolDirective(directive)

		case "js_package_file":
			nestjsConfig.parsePackageJSON(c, f, directive)

		case "js_import_alias":
			vals := strings.SplitN(directive.Value, " ", 2)
			nestjsConfig.ImportAliases = append(
				nestjsConfig.ImportAliases,
				struct{ From, To string }{From: vals[0], To: strings.TrimSpace(vals[1])},
			)

			// Regenerate ImportAliasPattern
			keyPatterns := make([]string, 0, len(nestjsConfig.ImportAliases))
			for _, alias := range nestjsConfig.ImportAliases {
				keyPatterns = append(
					keyPatterns,
					fmt.Sprintf("(^%s)", regexp.QuoteMeta(alias.From)),
				)
			}

			var err error
			if nestjsConfig.ImportAliasPattern, err = regexp.Compile(strings.Join(keyPatterns, "|")); err != nil {
				log.Fatalf(Err("failed to parse %s: %v", directive.Value, err))
			}

		case "js_visibility":
			nestjsConfig.Visibility.Set(directive.Value)
		case "js_default_npm_label":
			nestjsConfig.DefaultNpmLabel = directive.Value
			if !strings.HasSuffix(nestjsConfig.DefaultNpmLabel, ":") &&
				!strings.HasSuffix(nestjsConfig.DefaultNpmLabel, "/") {
				nestjsConfig.DefaultNpmLabel += "/"
			}

		case "js_root":
			jSRoot, err := filepath.Rel(".", f.Pkg)
			if err != nil {
				log.Fatalf(Err("failed to read directive %s: %v", directive.Key, err))
			} else {
				nestjsConfig.Root = jSRoot
				nestjsConfig.CollectedAssets = make(map[string]bool)
			}

		case "js_collect_barrels":
			nestjsConfig.CollectBarrels = readBoolDirective(directive)

		case "js_aggregate_modules":
			nestjsConfig.CollectBarrels = readBoolDirective(directive)

		case "js_collect_web_assets":
			nestjsConfig.CollectWebAssets = readBoolDirective(directive)

		case "js_aggregate_web_assets":
			nestjsConfig.CollectWebAssets = readBoolDirective(directive)

		case "js_collect_all_assets":
			nestjsConfig.CollectAllAssets = readBoolDirective(directive)

		case "js_aggregate_all_assets":
			nestjsConfig.CollectAllAssets = readBoolDirective(directive)

		case "js_collect_all":
			collectRoot, err := filepath.Rel(".", f.Pkg)
			if err != nil {
				log.Fatalf(Err("failed to read directive %s: %v", directive.Key, err))
			} else {
				nestjsConfig.CollectAllRoot = collectRoot
				nestjsConfig.CollectAll = true
				nestjsConfig.CollectAllSources = make(map[string]bool)
			}

		case "js_jest_config":
			nestjsConfig.JestConfigRelativePath = labels.ParseRelative(directive.Value, f.Pkg).
				Format()
			nestjsConfig.jestConfigPath = path.Join(f.Pkg, nestjsConfig.JestConfigRelativePath)

		case "js_jest_test_per_shard":
			nestjsConfig.JestTestsPerShard = readIntDirective(directive)

		case "js_jest_size":
			nestjsConfig.JestSize = directive.Value

		case "js_web_asset":
			vals := strings.SplitN(directive.Value, " ", 2)
			suffixes := vals[0]
			status := false
			if len(vals) > 1 {
				val, err := strconv.ParseBool(directive.Value)
				if err != nil {
					log.Fatalf(Err("failed to read directive %s: %v", directive.Key, err))
				}
				status = val
			}
			for _, suffix := range strings.Split(suffixes, ",") {
				nestjsConfig.WebAssetSuffixes[suffix] = status
			}

		case "js_quiet":
			nestjsConfig.Quiet = readBoolDirective(directive)
			if nestjsConfig.Quiet {
				nestjsConfig.Verbose = false
			}

		case "js_verbose":
			nestjsConfig.Verbose = readBoolDirective(directive)
			if nestjsConfig.Verbose {
				nestjsConfig.Quiet = false
			}
		case "nest_cli_path":
			nestjsConfig.NestCliPath = directive.Value
		case "nest_ts_config":
			nestjsConfig.RootTsConfigFile = directive.Value
		}
	}

	if nestjsConfig.RootPkg != rel {
		log.Printf("nestjsConfig.RootPkg != rel, %v != %v", nestjsConfig.RootPkg, rel)
		return
	}

	var root string
	if c.ReadBuildFilesDir != "" {
		root = path.Join(c.ReadBuildFilesDir, f.Pkg)
	} else {
		root = path.Join(c.RepoRoot, f.Pkg)
	}

	nestjsConfig.Root = root

	if rel == nestjsConfig.RootPkg {
		nestjsConfig.parseRootTsConfig()
	}

	nestjsConfig.parseNestCliJSON(*c, f)
	nestjsConfig.parseJestConfig()
}

var jsTestExtensions = []string{
	".test.js",
	".test.jsx",
}

var tsTestExtensions = []string{
	".test.ts",
	".test.tsx",
}

var tsExtensions = []string{
	".ts",
	".tsx",
}

var jsExtensions = []string{
	".js",
	".jsx",
}

var jsTestExtensionsPattern *regexp.Regexp
var tsTestExtensionsPattern *regexp.Regexp
var tsExtensionsPattern *regexp.Regexp
var jsExtensionsPattern *regexp.Regexp

func init() { tsTestExtensionsPattern = extensionPattern(tsTestExtensions) }
func init() { jsTestExtensionsPattern = extensionPattern(jsTestExtensions) }
func init() { tsExtensionsPattern = extensionPattern(tsExtensions) }
func init() { jsExtensionsPattern = extensionPattern(jsExtensions) }

func extensionPattern(extensions []string) *regexp.Regexp {
	escaped := make([]string, len(extensions))
	for i := range extensions {
		escaped[i] = fmt.Sprintf("(%s$)", regexp.QuoteMeta(extensions[i]))
	}
	return regexp.MustCompile(strings.Join(escaped, "|"))
}

var indexFilePattern *regexp.Regexp
var trimExtPattern *regexp.Regexp

func init() {
	escaped := make([]string, len(tsExtensions)+len(jsExtensions))
	for i, ext := range append(tsExtensions, jsExtensions...) {
		escaped[i] = regexp.QuoteMeta(ext)
	}
	indexFilePattern = regexp.MustCompile(
		fmt.Sprintf(`(index)(%s)$`,
			strings.Join(escaped, "|"),
		),
	)
	trimExtPattern = regexp.MustCompile(
		fmt.Sprintf(`(\S+)(%s)$`,
			strings.Join(escaped, "|"),
		),
	)
}

func trimExt(baseName string) string {
	matches := trimExtPattern.FindStringSubmatch(baseName)
	if len(matches) > 0 {
		return matches[1]
	}
	return baseName
}

func isBarrelFile(baseName string) bool {
	return indexFilePattern.MatchString(baseName)
}

func readBoolDirective(directive rule.Directive) bool {
	if directive.Value == "" {
		return true
	} else {
		val, err := strconv.ParseBool(directive.Value)
		if err != nil {
			log.Fatalf(Err("failed to read directive %s: %v", directive.Key, err))
		}
		return val
	}
}

func readIntDirective(directive rule.Directive) int {
	if directive.Value == "" {
		return -1
	} else {
		val, err := strconv.ParseInt(directive.Value, 10, 32)
		if err != nil {
			log.Fatalf(Err("failed to read directive %s: %v", directive.Key, err))
		}
		return int(val)
	}
}

func (parent *NestjsConfig) NewChild() *NestjsConfig {

	child := NewNestjsConfig()

	child.Enabled = parent.Enabled

	child.PackageFile = parent.PackageFile

	// copy maps
	child.NpmDependencies = struct {
		Dependencies    map[string]string "json:\"dependencies\""
		DevDependencies map[string]string "json:\"devDependencies\""
	}{
		Dependencies:    make(map[string]string),
		DevDependencies: make(map[string]string),
	}
	for k, v := range parent.NpmDependencies.Dependencies {
		child.NpmDependencies.Dependencies[k] = v
	}
	for k, v := range parent.NpmDependencies.DevDependencies {
		child.NpmDependencies.DevDependencies[k] = v
	}

	child.LookupTypes = parent.LookupTypes
	child.ImportAliases = parent.ImportAliases
	child.ImportAliases = make([]struct{ From, To string }, len(parent.ImportAliases)) // copy slice
	for i := range parent.ImportAliases {
		child.ImportAliases[i] = parent.ImportAliases[i]
	}
	child.ImportAliasPattern = parent.ImportAliasPattern // Regenerated on change to ImportAliases

	child.Visibility = Visibility{
		Labels: make([]string, len(parent.Visibility.Labels)), // copy slice
	}
	for i := range parent.Visibility.Labels {
		child.Visibility.Labels[i] = parent.Visibility.Labels[i]
	}
	child.CollectBarrels = parent.CollectBarrels
	child.CollectWebAssets = parent.CollectWebAssets
	child.CollectAllAssets = parent.CollectAllAssets
	child.CollectedAssets = parent.CollectedAssets // Reinitialized on change to JSRoot

	child.CollectAll = parent.CollectAll
	child.CollectAllRoot = parent.CollectAllRoot

	child.SourcePerProject = parent.SourcePerProject // copy map
	child.CollectAllSources = make(
		map[string]bool,
	) // Copy reference, reinitialized on change to CollectAll

	child.JestTestsPerShard = parent.JestTestsPerShard
	child.JestSize = parent.JestSize
	child.JestConfigRelativePath = parent.JestConfigRelativePath

	child.Root = parent.Root
	child.RootPkg = parent.RootPkg
	child.WebAssetSuffixes = make(map[string]bool) // copy map
	for k, v := range parent.WebAssetSuffixes {
		child.WebAssetSuffixes[k] = v
	}
	child.Quiet = parent.Quiet
	child.Verbose = parent.Verbose
	child.DefaultNpmLabel = parent.DefaultNpmLabel
	child.IsNestjs = parent.IsNestjs
	child.PackageByDir = parent.PackageByDir

	child.jestConfigPath = parent.jestConfigPath
	child.jestConfig = parent.jestConfig
	child.IsIgnoreE2E = parent.IsIgnoreE2E
	child.PackageJSONPath = parent.PackageJSONPath
	child.PnpmLockFilePath = parent.PnpmLockFilePath

	child.RootTsConfigFile = parent.RootTsConfigFile
	child.Transpiler = parent.Transpiler
	child.RootTsConfig = parent.RootTsConfig

	return child
}
