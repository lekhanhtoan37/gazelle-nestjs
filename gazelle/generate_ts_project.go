package nestjs

import (
	"path"

	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/rule"
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

	name := PkgName(args.Rel) + "_ts"
	if nestjsConfig.CollectWithTsConfig {
		// add as a folder
		for _, existingRule := range lang.readExistingRules(args, false) {
			// Look for existing rules with the same name, but different kind
			if existingRule.Name() == name &&
				existingRule.Kind() != getKind(args.Config, "ts_project") &&
				existingRule.Kind() != getKind(args.Config, "js_library") {
				name = name + "_ts"
			}
		}

		folderImports, folderRule := lang.makeFolderRule(moduleRuleArgs{
			pkgName:  name,
			cwd:      args.Rel,
			ruleType: getKind(args.Config, "ts_project"),
			srcs:     tsSources,
			imports:  imports,
		}, nestjsConfig)
		generatedRules = append(generatedRules, folderRule)
		generatedImports = append(generatedImports, folderImports)
	}

	return generatedRules, generatedImports
}
