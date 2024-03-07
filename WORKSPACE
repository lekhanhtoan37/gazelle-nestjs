workspace(name = 'gazelle_nestjs')

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")


http_archive(
    name = "io_bazel_rules_go",
    sha256 = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

load("//tools:deps.bzl", "gazelle_deps")

# gazelle:repository_macro deps.bzl%gazelle_deps
gazelle_deps()

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

gazelle_dependencies()

http_archive(
    name = "aspect_rules_js",
    sha256 = "630a71aba66c4023a5b16ab3efafaeed8b1a2865ccd168a34611eb73876b3fc4",
    strip_prefix = "rules_js-1.37.1",
    url = "https://github.com/aspect-build/rules_js/releases/download/v1.37.1/rules_js-v1.37.1.tar.gz",
)

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")

rules_js_dependencies()

load("@rules_nodejs//nodejs:repositories.bzl", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "nodejs",
    node_version = "18.17.0",
)

http_archive(
    name = "aspect_bazel_lib",
    sha256 = "979667bb7276ee8fcf2c114c9be9932b9a3052a64a647e0dcaacfb9c0016f0a3",
    strip_prefix = "bazel-lib-2.4.1",
    url = "https://github.com/aspect-build/bazel-lib/releases/download/v2.4.1/bazel-lib-v2.4.1.tar.gz",
)
