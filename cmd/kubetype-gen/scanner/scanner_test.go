// Copyright Istio Authors
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

package scanner

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/gengo/v2/codetags"
)

func TestTagValues(t *testing.T) {
	tags := codetags.Extract("+", []string{
		"+kubetype-gen:groupVersion=extensions.istio.io/v1alpha1",
		"+kubetype-gen:kubeType=TrafficExtension",
		"+kubetype-gen",
	})

	got := tagValues(tags, groupVersionTagName)
	want := []string{"extensions.istio.io/v1alpha1"}
	if len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("tagValues(groupVersion) = %#v, want %#v", got, want)
	}

	got = tagValues(tags, kubeTypeTagName)
	want = []string{"TrafficExtension"}
	if len(got) != len(want) || got[0] != want[0] {
		t.Fatalf("tagValues(kubeType) = %#v, want %#v", got, want)
	}
}

func TestGetGroupVersion(t *testing.T) {
	tags := codetags.Extract("+", []string{
		"+kubetype-gen:groupVersion=networking.istio.io/v1alpha3",
	})

	gv, err := getGroupVersion(tags, nil)
	if err != nil {
		t.Fatalf("getGroupVersion returned error: %v", err)
	}
	if gv == nil {
		t.Fatal("getGroupVersion returned nil")
	}
	want := schema.GroupVersion{Group: "networking.istio.io", Version: "v1alpha3"}
	if *gv != want {
		t.Fatalf("getGroupVersion = %#v, want %#v", *gv, want)
	}
}
