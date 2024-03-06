package nestjs

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

func (lang *NestJS) genJestTest(
	args language.GenerateArgs,
	nestjsConfig *NestjsConfig,
	project *Project,
	jestSources []string,
	tsSources []string,
) ([]*rule.Rule, []interface{}) {
	generatedRules := make([]*rule.Rule, 0)
	generatedImports := make([]interface{}, 0)

	// Ignore if subdirectory that is not a project
	if project == nil {
		return generatedRules, generatedImports
	}

	// Gen ts_project for test
	generatedTsRules, generatedTsImports := lang.genTsProjectForTest(args, nestjsConfig, project, jestSources)
	generatedRules = append(generatedRules, generatedTsRules...)
	generatedImports = append(generatedImports, generatedTsImports...)

	// Add each test as an individual rule
	for _, baseName := range jestSources {
		filePath := path.Join(args.Dir, baseName)
		ruleName := project.Name + "_test"
		r := rule.NewRule(
			getKind(args.Config, "jest_test"),
			ruleName,
		)
		r.SetAttr("data", []string{baseName})
		relativePart := path.Dir(baseName)
		imports, jestTestCount := readFileAndParse(filePath, relativePart)
		imports.set[baseName] = true

		// TODO: Add snapshot support
		var collectedSnapshots []string
		snapshotFile, err := os.Stat(path.Join(args.Dir, "__snapshots__", baseName+".snap"))
		if err == nil && snapshotFile.Mode().IsRegular() {
			collectedSnapshots = append(
				collectedSnapshots,
				path.Join("__snapshots__", baseName+".snap"),
			)
		}

		addJestAttributes(args, nestjsConfig, ruleName, r, jestTestCount, collectedSnapshots)

		generatedRules = append(generatedRules, r)
		generatedImports = append(generatedImports, imports)
	}

	return generatedRules, generatedImports
}

func addJestAttributes(
	args language.GenerateArgs,
	config *NestjsConfig,
	baseName string,
	r *rule.Rule,
	jestTestCount int,
	collectedSnapshots []string,
) {
	if config.JestConfigRelativePath == "" && config.jestConfigPath == "" && !config.Quiet {
		log.Print(
			Warn(
				"[%s/%s] no config for jest_test, use gazelle:js_jest_config directive",
				args.Rel,
				baseName,
			),
		)
	}

	if config.jestConfigPath == config.PackageJSONPath {
		r.SetAttr("config", fmt.Sprintf("//:%s_package_json", config.RootPkg))
	} else {
		r.SetAttr("config", config.JestConfigRelativePath)
	}

	if config.DefaultNpmLabel != "" {
		r.SetAttr("node_modules", config.DefaultNpmLabel)
	}

	if config.JestTestsPerShard > 0 {
		shardCount := int(math.Ceil(float64(jestTestCount) / float64(config.JestTestsPerShard)))
		if shardCount > 1 {
			r.SetAttr("shard_count", shardCount)
		}
	}
	if config.JestSize != "" {
		r.SetAttr("size", config.JestSize)
	}
	if len(config.Visibility.Labels) > 0 {
		r.SetAttr("visibility", config.Visibility.Labels)
	}
	if len(collectedSnapshots) > 0 {
		r.SetAttr("snapshots", collectedSnapshots)
	} else {
		r.DelAttr("snapshots")
	}
}
