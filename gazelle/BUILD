load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "nestjs",
    srcs = glob(include=["*.go"], exclude=["*_test.go"]),
    importpath = "github.com/lekhanhtoan37/gazelle_nestjs/gazelle",
    visibility = ["//visibility:public"],
    deps = [
        "@bazel_gazelle//config:go_default_library",
        "@bazel_gazelle//label:go_default_library",
        "@bazel_gazelle//language:go_default_library",
        "@bazel_gazelle//repo:go_default_library",
        "@bazel_gazelle//resolve:go_default_library",
        "@bazel_gazelle//rule:go_default_library",
        "@com_github_bazelbuild_buildtools//labels:go_default_library",
    ],
)

go_test(
    name = "gazelle_test",
    srcs = glob(include=["*_test.go"]),
    embed = [":gazelle"],
)