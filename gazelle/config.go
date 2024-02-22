package nestjs

import (
	"fmt"
	"regexp"
)

type NestjsConfig struct {
	Enabled          bool
	RootTsConfigFile string
	PackageFile      string
	NpmDependencies  struct {
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}
	LookupTypes        bool
	ImportAliases      []struct{ From, To string }
	ImportAliasPattern *regexp.Regexp
	Visibility         Visibility
	CollectBarrels     bool
	CollectWebAssets   bool
	CollectAllAssets   bool
	CollectedAssets    map[string]bool
	CollectAll         bool
	CollectAllRoot     string
	CollectAllSources  map[string]bool
	Fix                bool
	JSRoot             string
	WebAssetSuffixes   map[string]bool
	Quiet              bool
	Verbose            bool
	DefaultNpmLabel    string
	JestConfig         string
	JestTestsPerShard  int
	JestSize           string
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
