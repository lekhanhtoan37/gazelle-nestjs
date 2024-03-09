load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "gazelle",
    srcs = [
        "colors.go",
        "config.go",
        "fix.go",
        "fix_package_json.go",
        "generate.go",
        "generate_jest.go",
        "generate_ts_project.go",
        "import.go",
        "jest_config.go",
        "kinds.go",
        "lang.go",
        "nest_cli_config.go",
        "parse.go",
        "resolve.go",
        "ts_config.go",
        "type.go",
    ],
    importpath = "github.com/lekhanhtoan37/gazelle-nestjs/gazelle",
    visibility = ["//visibility:public"],
    deps = [
        "@bazel_gazelle//config:go_default_library",
        "@bazel_gazelle//label:go_default_library",
        "@bazel_gazelle//language:go_default_library",
        "@bazel_gazelle//repo:go_default_library",
        "@bazel_gazelle//resolve:go_default_library",
        "@bazel_gazelle//rule:go_default_library",
        "@com_github_bazelbuild_buildtools//build",
        "@com_github_bazelbuild_buildtools//labels",
    ],
)

go_test(
    name = "gazelle_test",
    srcs = [
        "generate_jest_test.go",
        "generate_ts_project_test.go",
        "parse_test.go",
    ],
    embed = [":gazelle"],
)
