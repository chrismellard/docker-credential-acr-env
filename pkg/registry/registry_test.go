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

import "testing"

func TestGetRegistryURL(t *testing.T) {
	cases := []struct {
		testDescription string
		input           string
		expectedOutput  string
		expectedError   bool
	}{
		{
			testDescription: "plain registry without scheme",
			input:           "myregistry.azurecr.io",
			expectedOutput:  "https://myregistry.azurecr.io",
			expectedError:   false,
		},
		{
			testDescription: "plain registry with scheme",
			input:           "https://myregistry.azurecr.io",
			expectedOutput:  "https://myregistry.azurecr.io",
			expectedError:   false,
		},
		{
			testDescription: "registry with container",
			input:           "myregistry.azurecr.io/mycontainer",
			expectedOutput:  "https://myregistry.azurecr.io",
			expectedError:   false,
		},
		{
			testDescription: "registry with container and query",
			input:           "myregistry.azurecr.io/mycontainer?foo=bar",
			expectedOutput:  "https://myregistry.azurecr.io",
			expectedError:   false,
		},
	}

	for i, c := range cases {
		t.Logf("Test #%d: %s", i, c.testDescription)
		result, err := getRegistryURL(c.input)

		if err != nil && !c.expectedError {
			t.Errorf("received an unexpected error: %v", err)
			continue
		}

		if err == nil && c.expectedError {
			t.Error("expected an error but didn't receive one")
			continue
		}

		if result.String() != c.expectedOutput {
			t.Errorf("got: %s, want: %s", result.String(), c.expectedOutput)
		}
	}
}
