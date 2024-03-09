load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")

# gazelle:prefix github.com/lekhanhtoan37/gazelle-nestjs
gazelle(
    name = "gazelle",
    gazelle = ":gazelle_bin",
    prefix = "github.com/lekhanhtoan37/gazelle-nestjs",
    visibility = ["//visibility:public"],
)

gazelle_binary(
    name = "gazelle_bin",
    languages = DEFAULT_LANGUAGES,
    visibility = ["//visibility:public"],
)

gazelle(
    name = "test_gazelle",
    gazelle = ":test_gazelle_bin",
    prefix = "github.com/lekhanhtoan37/gazelle-nestjs",
    visibility = ["//visibility:public"],
)

gazelle_binary(
    name = "test_gazelle_bin",
    languages = DEFAULT_LANGUAGES + [
        "//gazelle:gazelle",
    ],
    visibility = ["//visibility:public"],
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%gazelle_deps",
        "-prune",
    ],
    command = "update-repos",
)
