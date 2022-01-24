/*
Copyright © 2022 Chris Mellard chris.mellard@icloud.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package registry

import (
	"testing"
)

func TestGetRegistryURL(t *testing.T) {
	for _, c := range []struct {
		in, want string
		wantErr  bool
	}{
		{"myregistry.azurecr.io", "myregistry.azurecr.io", false},
		{"myregistry.azurecr.io/with/repo", "myregistry.azurecr.io", false},
		{"INVALID -- @*)%(@*)(@#*%", "", true},
	} {
		t.Run(c.in, func(t *testing.T) {
			got, err := getRegistryURL(c.in)
			if err != nil && !c.wantErr {
				t.Fatalf("unexpected error: %v", err)
			}
			if err == nil && c.wantErr {
				t.Fatal("wanted error, got nil")
			}
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}
