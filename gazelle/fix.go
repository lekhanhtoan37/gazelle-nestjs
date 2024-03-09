package nestjs

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

// Fix repairs deprecated usage of language-specific rules in f. This is
// called before the file is indexed. Unless c.ShouldFix is true, fixes
// that delete or rename rules should not be performed.Í
func (*NestJS) Fix(c *config.Config, f *rule.File) {

	nestJSConfigs := c.Exts[languageName].(NestjsConfigs)
	nestJSConfig := nestJSConfigs[f.Pkg]

	if c.ShouldFix || nestJSConfig.Fix {
		for _, r := range f.Rules {
			// delete deprecated js_import rule
			if r.Kind() == "js_import" {
				r.Delete()
			}
			// delete deprecated ts_library rule
			if r.Kind() == "ts_library" {
				r.Delete()
			}
			// delete deprecated ts_definition rule
			if r.Kind() == "ts_definition" {
				r.Delete()
			}
		}
		for _, l := range f.Loads {

			if l.Has("js_import") {
				l.Remove("js_import")
			}
			if l.Has("ts_library") {
				l.Remove("ts_library")
			}
			if l.Has("ts_definition") {
				l.Remove("ts_definition")
			}
		}
	}

	if nestJSConfig.RootTsConfig == nil {
		return
	}

	internalPkg := make(map[string]bool)
	for pkg := range nestJSConfig.RootTsConfig.CompilerOptions.Paths {
		if strings.HasSuffix(pkg, "*") {
			continue
		}

		internalPkg[pkg] = true
	}

	err := updateRootPackageJSON(nestJSConfig.PackageJSONPath, internalPkg)
	if err != nil {
		// handle error
	}

	for _, project := range nestJSConfig.PackageByDir {
		if project.Type == "application" {
			continue
		}

		fullpath := filepath.Join(nestJSConfig.Root, project.Root, "package.json")
		for pkg := range internalPkg {
			paths := strings.Split(pkg, "/")
			if len(paths) > 2 && len(paths) <= 0 {
				log.Printf(Err("Invalid internal package: %s", pkg))
				continue
			}

			if paths[1] != project.Name {
				continue
			}

			path := filepath.Join(project.SourceRoot, project.EntryFile)
			relativePath, err := filepath.Rel(project.Root, path)
			if err != nil {
				// handle error
				continue
			}
			prefix := ".." + string(os.PathSeparator)
			isChildPkg := !strings.HasPrefix(relativePath, prefix) && relativePath != ".."

			if !isChildPkg {
				continue
			}

			packageJSON := newPackageJSON(pkg, relativePath)
			err = fixPackageJSON(fullpath, packageJSON)
			if err != nil {
				// handle error
				continue
			}
		}
	}

	// internal_nestjs.RegeneratePnpmLockFile(nestJSConfig.PnpmLockFilePath)
}
