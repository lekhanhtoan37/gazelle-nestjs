package nestjs

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
	bzl "github.com/bazelbuild/buildtools/build"
)

// func (lang *NestJS) genTsProjectV2(args language.GenerateArgs, nestjsConfig *NestjsConfig) string {
// 	return languageName
// }

func (lang *NestJS) genTsProject(
	args language.GenerateArgs,
	nestjsConfig *NestjsConfig,
	project *Project,
	tsSources []string,
	// testSources []string,
) ([]*rule.Rule, []interface{}) {
	var imports []imports

	generatedRules := make([]*rule.Rule, 0)
	generatedImports := make([]interface{}, 0)

	if nestjsConfig == nil || len(tsSources) <= 0 {
		return generatedRules, generatedImports
	}

	for _, baseName := range tsSources {
		filePath := path.Join(args.Dir, baseName)
		relativePart := ""
		if nestjsConfig.CollectWithTsConfig {
			relativePart = path.Dir(baseName)
		}
		imps, _ := readFileAndParse(filePath, relativePart)
		imports = append(imports, *imps)
	}

	if nestjsConfig.CollectWithTsConfig {
		// add as a folder
		// for _, existingRule := range lang.readExistingRules(args, false) {
		// 	// Look for existing rules with the same name, but different kind
		// 	if existingRule.Name() == name &&
		// 		existingRule.Kind() != getKind(args.Config, "ts_project") &&
		// 		existingRule.Kind() != getKind(args.Config, "js_library") {
		// 		name = name + "_ts"
		// 	}
		// }

		moduleImports := flattenImports(imports)

		// Use lists to make a rule
		ruleName := project.Name
		if project.Type == "library" {
			ruleName = ruleName + "_lib"
		}
		moduleRule := rule.NewRule(getKind(args.Config, "ts_project"), ruleName)
		moduleRule.SetAttr("srcs", tsSources)
		if len(nestjsConfig.Visibility.Labels) > 0 {
			moduleRule.SetAttr("visibility", nestjsConfig.Visibility.Labels)
		}

		if project.TsConfig != nil {
			tsConfigFile, err := filepath.Rel(args.Rel, project.TsConfigPath)

			// Tsconfig file is in the same directory as the project
			if err == nil && !strings.Contains(tsConfigFile, "..") {
				moduleRule.SetAttr("tsconfig", ":"+tsConfigFile)
			}

			if project.TsConfig.Extends != "" {
				if nestjsConfig.RootPkg == "" || nestjsConfig.RootPkg == "." {
					moduleRule.SetAttr("extends", "//:"+"root"+"_tsconfig")
				} else {
					moduleRule.SetAttr("extends", "//:"+nestjsConfig.RootPkg+"_tsconfig")
				}
			}
		}

		if nestjsConfig.Transpiler == swc {
			transpilerExpr := &bzl.CallExpr{
				X: &bzl.DotExpr{
					X:    &bzl.Ident{Name: "partial"},
					Name: "make",
				},
				List: []bzl.Expr{&bzl.Ident{Name: "swc"}, &bzl.AssignExpr{
					LHS: &bzl.LiteralExpr{Token: "swcrc"},
					Op:  "=",
					RHS: rule.ExprFromValue("//:.swcrc"),
				}},
			}

			moduleRule.SetAttr("transpiler", transpilerExpr)
		}

		if nestjsConfig.RootTsConfig != nil {
			moduleRule.SetAttr("declaration", nestjsConfig.RootTsConfig.CompilerOptions.Declaration)
			moduleRule.SetAttr("source_map", nestjsConfig.RootTsConfig.CompilerOptions.SourceMap)
			moduleRule.SetAttr("incremental", nestjsConfig.RootTsConfig.CompilerOptions.Incremental)
		}

		// folderImports, folderRule := lang.makeFolderRule(moduleRuleArgs{
		// 	pkgName:  project.Name,
		// 	cwd:      args.Rel,
		// 	ruleType: getKind(args.Config, "ts_project"),
		// 	srcs:     tsSources,
		// 	imports:  imports,
		// }, nestjsConfig)

		// generatedRules = append(generatedRules, folderRule)
		// generatedImports = append(generatedImports, folderImports)

		generatedRules = append(generatedRules, moduleRule)
		generatedImports = append(generatedImports, moduleImports)
	}

	return generatedRules, generatedImports
}

func (lang *NestJS) genTsProjectForTest(
	args language.GenerateArgs,
	nestjsConfig *NestjsConfig,
	project *Project,
	jestSources []string,
) ([]*rule.Rule, []interface{}) {
	var imports []imports

	generatedRules := make([]*rule.Rule, 0)
	generatedImports := make([]interface{}, 0)

	if nestjsConfig == nil || len(jestSources) <= 0 {
		return generatedRules, generatedImports
	}

	for _, baseName := range jestSources {
		filePath := path.Join(args.Dir, baseName)
		relativePart := ""
		if nestjsConfig.CollectWithTsConfig {
			relativePart = path.Dir(baseName)
		}
		imps, _ := readFileAndParse(filePath, relativePart)
		imports = append(imports, *imps)
	}

	if nestjsConfig.CollectWithTsConfig {
		// add as a folder
		// for _, existingRule := range lang.readExistingRules(args, false) {
		// 	// Look for existing rules with the same name, but different kind
		// 	if existingRule.Name() == name &&
		// 		existingRule.Kind() != getKind(args.Config, "ts_project") &&
		// 		existingRule.Kind() != getKind(args.Config, "js_library") {
		// 		name = name + "_ts"
		// 	}
		// }

		moduleImports := flattenImports(imports)

		// Use lists to make a rule
		moduleRule := rule.NewRule(getKind(args.Config, "ts_project"), project.Name+"_ts_test")
		moduleRule.SetAttr("srcs", jestSources)
		if len(nestjsConfig.Visibility.Labels) > 0 {
			moduleRule.SetAttr("visibility", nestjsConfig.Visibility.Labels)
		}

		if project.TsConfig != nil {
			tsConfigFile, err := filepath.Rel(args.Rel, project.TsConfigPath)

			// Tsconfig file is in the same directory as the project
			if err == nil && !strings.Contains(tsConfigFile, "..") {
				moduleRule.SetAttr("tsconfig", &bzl.DictExpr{})
			}

			if project.TsConfig.Extends != "" {
				if nestjsConfig.RootPkg == "" || nestjsConfig.RootPkg == "." {
					moduleRule.SetAttr("extends", "//:"+"root"+"_tsconfig")
				} else {
					moduleRule.SetAttr("extends", "//:"+nestjsConfig.RootPkg+"_tsconfig")
				}
			}
		}

		if nestjsConfig.Transpiler == swc {
			transpilerExpr := &bzl.CallExpr{
				X: &bzl.DotExpr{
					X:    &bzl.Ident{Name: "partial"},
					Name: "make",
				},
				List: []bzl.Expr{&bzl.Ident{Name: "swc"}, &bzl.AssignExpr{
					LHS: &bzl.LiteralExpr{Token: "swcrc"},
					Op:  "=",
					RHS: rule.ExprFromValue("//:.swcrc"),
				}},
			}

			moduleRule.SetAttr("transpiler", transpilerExpr)
		}

		if nestjsConfig.RootTsConfig != nil {
			moduleRule.SetAttr("declaration", nestjsConfig.RootTsConfig.CompilerOptions.Declaration)
			moduleRule.SetAttr("source_map", nestjsConfig.RootTsConfig.CompilerOptions.SourceMap)
			moduleRule.SetAttr("incremental", nestjsConfig.RootTsConfig.CompilerOptions.Incremental)
		}

		// folderImports, folderRule := lang.makeFolderRule(moduleRuleArgs{
		// 	pkgName:  project.Name,
		// 	cwd:      args.Rel,
		// 	ruleType: getKind(args.Config, "ts_project"),
		// 	srcs:     tsSources,
		// 	imports:  imports,
		// }, nestjsConfig)

		// generatedRules = append(generatedRules, folderRule)
		// generatedImports = append(generatedImports, folderImports)

		generatedRules = append(generatedRules, moduleRule)
		generatedImports = append(generatedImports, moduleImports)
	}

	return generatedRules, generatedImports
}
