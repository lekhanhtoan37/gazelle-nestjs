# Design of Gazelle

<https://github.com/bazelbuild/bazel-gazelle/blob/master/Design.rst>

## See ./examples for nestjs monorepo

## To using plugin gazelle-nestjs for generating BUILD file. Make sure to load `rules_go`, `bazel_gazelle`, `gazelle_nestjs`

```starlark
http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-fHbWI2so/2laoozzX5XeMXqUcv0fsUrHl8m/aE8Js3w=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

http_archive(
  name = "com_github_lekhanhtoan37_gazelle_nestjs",
  sha256 = "7eef232ea8845e2a92d2bc534dae46a304a51deff07f770645b4871de95ae052",
  strip_prefix = "gazelle-nestjs-0.0.1",
  urls = [
    "https://github.com/lekhanhtoan37/gazelle-nestjs/archive/refs/tags/v0.0.1.tar.gz",
  ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################


go_rules_dependencies()

go_register_toolchains(version = "1.19")

gazelle_dependencies(go_repository_default_config = "@//:WORKSPACE.bazel")
```

## Add directive in root BUILD file

```starlark
load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")
load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "node_modules")

gazelle(
    name = "gazelle",
    args = ["-build_file_name=BUILD.bazel"],
    gazelle = ":gazelle_nestjs",
    visibility = ["//visibility:public"],
)

gazelle_binary(
    name = "gazelle_nestjs",
    languages = DEFAULT_LANGUAGES + [
        "@com_github_lekhanhtoan37_gazelle_nestjs//gazelle",
    ],
    visibility = ["//visibility:public"],
)

# gazelle:js_package_file package.json :node_modules
# gazelle:js_root
```

## Author
lekhanhtoan37