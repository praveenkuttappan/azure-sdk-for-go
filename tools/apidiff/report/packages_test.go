// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package report

import "testing"

func TestPkgsReport_ToMarkdown_nil(t *testing.T) {
	pr := &PkgsReport{}
	s := pr.ToMarkdown(nil)
	if len(s) > 0 {
		t.Fatalf("failed nil param test. Expected empty string, Received: %v", s)
	}
}

func TestPkgsReport_ToMarkdown_withVersionOnly(t *testing.T) {
	pr := &PkgsReport{AddedPackages: []string{"foo"}}
	s := pr.ToMarkdown([]string{"v0.1.0"})
	expected := `# Release History

## v0.1.0 (Released)

### New Packages

- foo
`
	if s != expected {
		t.Fatalf("failed nil param test. Expected: %v, Received: %v", expected, s)
	}
}

func TestPkgsReport_ToMarkdown_withIncompleteArgs(t *testing.T) {
	pr := &PkgsReport{AddedPackages: []string{"foo"}}
	s := pr.ToMarkdown([]string{"v0.1.0", "arg1"})
	expected := `# Release History

## v0.1.0 (Released)

### New Packages

- foo
`
	if s != expected {
		t.Fatalf("failed nil param test. Expected: %v, Received: %v", expected, s)
	}
}

func TestPkgsReport_ToMarkdown_withCompleteArgsSet(t *testing.T) {
	pr := &PkgsReport{AddedPackages: []string{"foo"}}
	s := pr.ToMarkdown([]string{"v0.1.0", "arg1", "arg2"})
	expected := `# Release History

## v0.1.0 (Released)

Generated from https://github.com/Azure/azure-rest-api-specs/tree/arg1
Code generator @autorest/go@arg2

### New Packages

- foo
`
	if s != expected {
		t.Fatalf("failed nil param test. Expected: %v, Received: %v", expected, s)
	}
}
