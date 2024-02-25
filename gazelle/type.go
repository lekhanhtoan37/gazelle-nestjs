package nestjs

import (
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

type NestJS struct {
}

// Embeds implements language.Language.
func (*NestJS) Embeds(r *rule.Rule, from label.Label) []label.Label {
	panic("unimplemented")
}

// Imports implements language.Language.
func (*NestJS) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
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
