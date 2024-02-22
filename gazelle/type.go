package nestjs

import (
	"flag"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/language"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

type NestJS struct {
}

// CheckFlags implements language.Language.
func (*NestJS) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	panic("unimplemented")
}

// Configure implements language.Language.
func (*NestJS) Configure(c *config.Config, rel string, f *rule.File) {
	panic("unimplemented")
}

// Embeds implements language.Language.
func (*NestJS) Embeds(r *rule.Rule, from label.Label) []label.Label {
	panic("unimplemented")
}

// Fix implements language.Language.
func (*NestJS) Fix(c *config.Config, f *rule.File) {
	panic("unimplemented")
}

// GenerateRules implements language.Language.
func (*NestJS) GenerateRules(args language.GenerateArgs) language.GenerateResult {
	panic("unimplemented")
}

// Imports implements language.Language.
func (*NestJS) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	panic("unimplemented")
}

// KnownDirectives implements language.Language.
func (*NestJS) KnownDirectives() []string {
	panic("unimplemented")
}

// Loads implements language.Language.
func (*NestJS) Loads() []rule.LoadInfo {
	panic("unimplemented")
}

// RegisterFlags implements language.Language.
func (*NestJS) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
	panic("unimplemented")
}

// Resolve implements language.Language.
func (*NestJS) Resolve(
	c *config.Config,
	ix *resolve.RuleIndex,
	rc *repo.RemoteCache,
	r *rule.Rule,
	imports interface{},
	from label.Label,
) {
	panic("unimplemented")
}
