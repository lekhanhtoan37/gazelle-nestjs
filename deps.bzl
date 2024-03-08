load("@bazel_gazelle//:deps.bzl", "go_repository")

# # Copyright 2017 The Bazel Authors. All rights reserved.
# #
# # Licensed under the Apache License, Version 2.0 (the "License");
# # you may not use this file except in compliance with the License.
# # You may obtain a copy of the License at
# #
# #    http://www.apache.org/licenses/LICENSE-2.0
# #
# # Unless required by applicable law or agreed to in writing, software
# # distributed under the License is distributed on an "AS IS" BASIS,
# # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# # See the License for the specific language governing permissions and
# # limitations under the License.

# load(
#     "@bazel_tools//tools/build_defs/repo:http.bzl",
#     "http_archive",
# )
# load(
#     "//internal:go_repository.bzl",
#     _go_repository = "go_repository",
# )
# load(
#     "//internal:go_repository_cache.bzl",
#     "go_repository_cache",
# )
# load(
#     "//internal:go_repository_tools.bzl",
#     "go_repository_tools",
# )
# load(
#     "//internal:go_repository_config.bzl",
#     "go_repository_config",
# )
# load(
#     "//internal:is_bazel_module.bzl",
#     "is_bazel_module",
# )

# # Re-export go_repository . Users should get it from this file.
# go_repository = _go_repository

# def gazelle_dependencies(
#         go_sdk = "",
#         go_repository_default_config = "@//:WORKSPACE",
#         go_env = {}):
#     _maybe(
#         http_archive,
#         name = "bazel_skylib",
#         urls = [
#             "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
#             "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
#         ],
#         sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
#     )

#     if go_sdk:
#         go_repository_cache(
#             name = "bazel_gazelle_go_repository_cache",
#             go_sdk_name = go_sdk,
#             go_env = go_env,
#         )
#     else:
#         go_sdk_info = {}
#         for name, r in native.existing_rules().items():
#             # match internal rule names but don't reference them directly.
#             # New rules may be added in the future, and they might be
#             # renamed (_go_download_sdk => go_download_sdk).
#             if name != "go_sdk" and ("go_" not in r["kind"] or "_sdk" not in r["kind"]):
#                 continue
#             if r.get("goos", "") and r.get("goarch", ""):
#                 platform = r["goos"] + "_" + r["goarch"]
#             else:
#                 platform = "host"
#             go_sdk_info[name] = platform
#         go_repository_cache(
#             name = "bazel_gazelle_go_repository_cache",
#             go_sdk_info = go_sdk_info,
#             go_env = go_env,
#         )

#     go_repository_tools(
#         name = "bazel_gazelle_go_repository_tools",
#         go_cache = "@bazel_gazelle_go_repository_cache//:go.env",
#     )

#     go_repository_config(
#         name = "bazel_gazelle_go_repository_config",
#         config = go_repository_default_config,
#     )

#     is_bazel_module(
#         name = "bazel_gazelle_is_bazel_module",
#         is_bazel_module = False,
#     )
#     _maybe(
#         go_repository,
#         name = "co_honnef_go_tools",
#         importpath = "honnef.co/go/tools",
#         sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
#         version = "v0.0.0-20190523083050-ea95bdfd59fc",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_bazelbuild_buildtools",
#         importpath = "github.com/bazelbuild/buildtools",
#         sum = "h1:2Gc2Q6hVR1SJ8bBI9Ybzoggp8u/ED2WkM4MfvEIn9+c=",
#         version = "v0.0.0-20231115204819-d4c9dccdfbb1",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_bazelbuild_rules_go",
#         importpath = "github.com/bazelbuild/rules_go",
#         sum = "h1:uJStI9o5obVWSwquy9WxKNWfZxf2sKA2rpEsX6x5RVM=",
#         version = "v0.44.0",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_bmatcuk_doublestar_v4",
#         importpath = "github.com/bmatcuk/doublestar/v4",
#         sum = "h1:FH9SifrbvJhnlQpztAx++wlkk70QBf0iBWDwNy7PA4I=",
#         version = "v4.6.1",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_burntsushi_toml",
#         importpath = "github.com/BurntSushi/toml",
#         sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
#         version = "v0.3.1",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_census_instrumentation_opencensus_proto",
#         importpath = "github.com/census-instrumentation/opencensus-proto",
#         sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
#         version = "v0.2.1",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_chzyer_logex",
#         importpath = "github.com/chzyer/logex",
#         sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
#         version = "v1.1.10",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_chzyer_readline",
#         importpath = "github.com/chzyer/readline",
#         sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
#         version = "v0.0.0-20180603132655-2972be24d48e",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_chzyer_test",
#         importpath = "github.com/chzyer/test",
#         sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
#         version = "v0.0.0-20180213035817-a1ea475d72b1",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_client9_misspell",
#         importpath = "github.com/client9/misspell",
#         sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
#         version = "v0.3.4",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_envoyproxy_go_control_plane",
#         importpath = "github.com/envoyproxy/go-control-plane",
#         sum = "h1:4cmBvAEBNJaGARUEs3/suWRyfyBfhf7I60WBZq+bv2w=",
#         version = "v0.9.1-0.20191026205805-5f8ba28d4473",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_envoyproxy_protoc_gen_validate",
#         importpath = "github.com/envoyproxy/protoc-gen-validate",
#         sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
#         version = "v0.1.0",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_fsnotify_fsnotify",
#         importpath = "github.com/fsnotify/fsnotify",
#         sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
#         version = "v1.7.0",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_golang_glog",
#         importpath = "github.com/golang/glog",
#         sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
#         version = "v0.0.0-20160126235308-23def4e6c14b",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_golang_mock",
#         importpath = "github.com/golang/mock",
#         sum = "h1:ErTB+efbowRARo13NNdxyJji2egdxLGQhRaY+DUumQc=",
#         version = "v1.6.0",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_golang_protobuf",
#         importpath = "github.com/golang/protobuf",
#         sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
#         version = "v1.5.2",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_google_go_cmp",
#         importpath = "github.com/google/go-cmp",
#         sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
#         version = "v0.6.0",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_pmezard_go_difflib",
#         importpath = "github.com/pmezard/go-difflib",
#         sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
#         version = "v1.0.0",
#     )
#     _maybe(
#         go_repository,
#         name = "com_github_prometheus_client_model",
#         importpath = "github.com/prometheus/client_model",
#         sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
#         version = "v0.0.0-20190812154241-14fe0d1b01d4",
#     )
#     _maybe(
#         go_repository,
#         name = "com_google_cloud_go",
#         importpath = "cloud.google.com/go",
#         sum = "h1:e0WKqKTd5BnrG8aKH3J3h+QvEIQtSUcf2n5UZ5ZgLtQ=",
#         version = "v0.26.0",
#     )
#     _maybe(
#         go_repository,
#         name = "net_starlark_go",
#         importpath = "go.starlark.net",
#         sum = "h1:xwwDQW5We85NaTk2APgoN9202w/l0DVGp+GZMfsrh7s=",
#         version = "v0.0.0-20210223155950-e043a3d3c984",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_google_appengine",
#         importpath = "google.golang.org/appengine",
#         sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
#         version = "v1.4.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_google_genproto",
#         importpath = "google.golang.org/genproto",
#         sum = "h1:+kGHl1aib/qcwaRi1CbqBZ1rk19r85MNUf8HaBghugY=",
#         version = "v0.0.0-20200526211855-cb27e3aa2013",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_google_grpc",
#         importpath = "google.golang.org/grpc",
#         sum = "h1:fPVVDxY9w++VjTZsYvXWqEf9Rqar/e+9zYfxKK+W+YU=",
#         version = "v1.50.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_google_protobuf",
#         importpath = "google.golang.org/protobuf",
#         sum = "h1:w43yiav+6bVFTBQFZX0r7ipe9JQ1QsbMgHwbBziscLw=",
#         version = "v1.28.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_crypto",
#         importpath = "golang.org/x/crypto",
#         sum = "h1:VklqNMn3ovrHsnt90PveolxSbWFaJdECFbxSq0Mqo2M=",
#         version = "v0.0.0-20190308221718-c2843e01d9a2",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_exp",
#         importpath = "golang.org/x/exp",
#         sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
#         version = "v0.0.0-20190121172915-509febef88a4",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_lint",
#         importpath = "golang.org/x/lint",
#         sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
#         version = "v0.0.0-20190313153728-d0100b6bd8b3",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_mod",
#         importpath = "golang.org/x/mod",
#         sum = "h1:dGoOF9QVLYng8IHTm7BAyWqCqSheQ5pYWGhzW00YJr0=",
#         version = "v0.14.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_net",
#         importpath = "golang.org/x/net",
#         sum = "h1:7eBu7KsSvFDtSXUIDbh3aqlK4DPsZ1rByC8PFfBThos=",
#         version = "v0.16.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_oauth2",
#         importpath = "golang.org/x/oauth2",
#         sum = "h1:vEDujvNQGv4jgYKudGeI/+DAX4Jffq6hpD55MmoEvKs=",
#         version = "v0.0.0-20180821212333-d2e6202438be",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_sync",
#         importpath = "golang.org/x/sync",
#         sum = "h1:60k92dhOjHxJkrqnwsfl8KuaHbn/5dl0lUPUklKo3qE=",
#         version = "v0.5.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_sys",
#         importpath = "golang.org/x/sys",
#         sum = "h1:h48lPFYpsTvQJZF4EKyI4aLHaev3CxivZmv7yZig9pc=",
#         version = "v0.15.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_text",
#         importpath = "golang.org/x/text",
#         sum = "h1:ablQoSUd0tRdKxZewP80B+BaqeKJuVhuRxj/dkrun3k=",
#         version = "v0.13.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_tools",
#         importpath = "golang.org/x/tools",
#         sum = "h1:zdAyfUGbYmuVokhzVmghFl2ZJh5QhcfebBgmVPFYA+8=",
#         version = "v0.15.0",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_tools_go_vcs",
#         importpath = "golang.org/x/tools/go/vcs",
#         sum = "h1:cOIJqWBl99H1dH5LWizPa+0ImeeJq3t3cJjaeOWUAL4=",
#         version = "v0.1.0-deprecated",
#     )
#     _maybe(
#         go_repository,
#         name = "org_golang_x_xerrors",
#         importpath = "golang.org/x/xerrors",
#         sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
#         version = "v0.0.0-20200804184101-5ec99f83aff1",
#     )

# def _maybe(repo_rule, name, **kwargs):
#     if name not in native.existing_rules():
#         repo_rule(name = name, **kwargs)

def go_dependencies():
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
        version = "v0.0.0-20190523083050-ea95bdfd59fc",
    )
    go_repository(
        name = "com_github_bazelbuild_bazel_gazelle",
        importpath = "github.com/bazelbuild/bazel-gazelle",
        sum = "h1:Bvg+zEHWYwWrhJT4WxyvcU3y1DEJpT/XlPYEfIn9lUI=",
        version = "v0.35.0",
    )
    go_repository(
        name = "com_github_bazelbuild_buildtools",
        importpath = "github.com/bazelbuild/buildtools",
        sum = "h1:wycIfuqZJzkloRT+fcazTM3NjvAMyAi1qC2QXmEZP4s=",
        version = "v0.0.0-20240207142252-03bf520394af",
    )
    go_repository(
        name = "com_github_bazelbuild_rules_go",
        importpath = "github.com/bazelbuild/rules_go",
        sum = "h1:uJStI9o5obVWSwquy9WxKNWfZxf2sKA2rpEsX6x5RVM=",
        version = "v0.44.0",
    )
    go_repository(
        name = "com_github_bmatcuk_doublestar_v4",
        importpath = "github.com/bmatcuk/doublestar/v4",
        sum = "h1:FH9SifrbvJhnlQpztAx++wlkk70QBf0iBWDwNy7PA4I=",
        version = "v4.6.1",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        importpath = "github.com/chzyer/logex",
        sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
        version = "v1.1.10",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        importpath = "github.com/chzyer/readline",
        sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
        version = "v0.0.0-20180603132655-2972be24d48e",
    )
    go_repository(
        name = "com_github_chzyer_test",
        importpath = "github.com/chzyer/test",
        sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
        version = "v0.0.0-20180213035817-a1ea475d72b1",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:4cmBvAEBNJaGARUEs3/suWRyfyBfhf7I60WBZq+bv2w=",
        version = "v0.9.1-0.20191026205805-5f8ba28d4473",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
        version = "v0.0.0-20160126235308-23def4e6c14b",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:JjCZWpVbqXDqFVmTfYWEVTMIYrL/NPdPSCHPJ0T/raM=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
        version = "v0.0.0-20190812154241-14fe0d1b01d4",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:e0WKqKTd5BnrG8aKH3J3h+QvEIQtSUcf2n5UZ5ZgLtQ=",
        version = "v0.26.0",
    )
    go_repository(
        name = "net_starlark_go",
        importpath = "go.starlark.net",
        sum = "h1:xwwDQW5We85NaTk2APgoN9202w/l0DVGp+GZMfsrh7s=",
        version = "v0.0.0-20210223155950-e043a3d3c984",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
        version = "v1.4.0",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:+kGHl1aib/qcwaRi1CbqBZ1rk19r85MNUf8HaBghugY=",
        version = "v0.0.0-20200526211855-cb27e3aa2013",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:rRYRFMVgRv6E0D70Skyfsr28tDXIuuPZyWGMPdMcnXg=",
        version = "v1.27.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:Ejskq+SyPohKW+1uil0JJMtmHCgJPJ/qWTxr8qp+R4c=",
        version = "v1.25.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:VklqNMn3ovrHsnt90PveolxSbWFaJdECFbxSq0Mqo2M=",
        version = "v0.0.0-20190308221718-c2843e01d9a2",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
        version = "v0.0.0-20190121172915-509febef88a4",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
        version = "v0.0.0-20190313153728-d0100b6bd8b3",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:dGoOF9QVLYng8IHTm7BAyWqCqSheQ5pYWGhzW00YJr0=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:oWX7TPOiFAMXLq8o0ikBYfCJVlRHBcsciT5bXOrH628=",
        version = "v0.0.0-20190311183353-d8887717615a",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:vEDujvNQGv4jgYKudGeI/+DAX4Jffq6hpD55MmoEvKs=",
        version = "v0.0.0-20180821212333-d2e6202438be",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:60k92dhOjHxJkrqnwsfl8KuaHbn/5dl0lUPUklKo3qE=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:h48lPFYpsTvQJZF4EKyI4aLHaev3CxivZmv7yZig9pc=",
        version = "v0.15.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:Iey4qkscZuv0VvIt8E0neZjtPVQFSc870HQ448QgEmQ=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_tools_go_vcs",
        importpath = "golang.org/x/tools/go/vcs",
        sum = "h1:cOIJqWBl99H1dH5LWizPa+0ImeeJq3t3cJjaeOWUAL4=",
        version = "v0.1.0-deprecated",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
