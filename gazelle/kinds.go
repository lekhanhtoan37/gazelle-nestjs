// Copyright 2019 The Bazel Authors. All rights reserved.
// Modifications copyright (C) 2021 BenchSci Analytics Inc.
// Modifications copyright (C) 2018 Ecosia GmbH

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nestjs

import (
	"github.com/bazelbuild/bazel-gazelle/rule"
)

// Kinds returns a map of maps rule names (kinds) and information on how to
// match and merge attributes that may be found in rules of those kinds. All
// kinds of rules generated for this language may be found here.
func (*NestJS) Kinds() map[string]rule.KindInfo {
	return map[string]rule.KindInfo{
		"js_library": {
			MatchAny: false,
			NonEmptyAttrs: map[string]bool{
				"srcs": true,
			},
			MergeableAttrs: map[string]bool{
				"srcs": true,
				"tags": true,
			},
			ResolveAttrs: map[string]bool{
				"deps": true,
				"data": true,
			},
		},
		"ts_project": {
			MatchAny: false,
			NonEmptyAttrs: map[string]bool{
				"srcs":       true,
				"tsconfig":   true,
				"extends":    true,
				"transpiler": true,
				"deps":       true,
			},
			MergeableAttrs: map[string]bool{
				"srcs": true,
				"tags": true,
			},
			ResolveAttrs: map[string]bool{
				"deps": true,
			},
		},
		"jest_test": {
			MatchAny: false,
			NonEmptyAttrs: map[string]bool{
				"data":         true,
				"config":       true,
				"node_modules": true,
			},
			MergeableAttrs: map[string]bool{
				"data": true,
			},
			ResolveAttrs: map[string]bool{
				"data": true,
			},
		},
		"npm_package": {
			MatchAny: false,
			NonEmptyAttrs: map[string]bool{
				"srcs": true,
			},
			ResolveAttrs: map[string]bool{
				"srcs": true,
			},
		},
		"ts_config": {
			MatchAny: false,
			NonEmptyAttrs: map[string]bool{
				"src": true,
			},
			MergeableAttrs: map[string]bool{
				"src": true,
			},
			ResolveAttrs: map[string]bool{
				"src":  true,
				"data": true,
			},
		},
		"exports_files": {
			MatchAny:       false,
			NonEmptyAttrs:  map[string]bool{},
			ResolveAttrs:   map[string]bool{},
			MergeableAttrs: map[string]bool{},
		},
	}
}
