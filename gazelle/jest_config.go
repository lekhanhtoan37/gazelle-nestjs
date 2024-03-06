package nestjs

import (
	"regexp"
	"strings"
)

type JestConfig struct {
	ModuleFileExtensions []string               `json:"moduleFileExtensions"`
	RootDir              string                 `json:"rootDir"`
	TestRegex            string                 `json:"testRegex"`
	Transform            map[string]interface{} `json:"transform"`
	Preset               string                 `json:"preset"`
	CollectCoverageFrom  []string               `json:"collectCoverageFrom"`
	CoverageDirectory    string                 `json:"coverageDirectory"`
	TestEnvironment      string                 `json:"testEnvironment"`
	Roots                []string               `json:"roots"`
	ModuleNameMapper     struct {
		AppAuth string `json:"^@app/auth(|/.*)$"`
		AppUser string `json:"^@app/user(|/.*)$"`
	} `json:"moduleNameMapper"`
	Verbose     bool `json:"verbose"`
	TestTimeout int  `json:"testTimeout"`
}

var e2eExtensions = []string{".*.e2e-spec.ts$", ".*.e2e-spec.js$"}

// func extensionPattern(extensions []string) *regexp.Regexp {
// 	escaped := make([]string, len(extensions))
// 	for i := range extensions {
// 		escaped[i] = fmt.Sprintf("(%s$)", regexp.QuoteMeta(extensions[i]))
// 	}
// 	return regexp.MustCompile(strings.Join(escaped, "|"))
// }

func (jestConfig *JestConfig) getTsTestPattern() *regexp.Regexp {
	return regexp.MustCompile(jestConfig.TestRegex)
}

func (jestConfig *JestConfig) getJsTestPattern() *regexp.Regexp {
	return regexp.MustCompile(jestConfig.TestRegex)
}

func (jestConfig *JestConfig) getE2eTestPattern() *regexp.Regexp {
	return regexp.MustCompile(strings.Join(e2eExtensions, "|"))
}
