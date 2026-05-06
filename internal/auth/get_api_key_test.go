package auth

import (
	"strings"
	"testing"
	"net/http"
	"fmt"
)

func TestGetAPIKey(t *testing.T){
	type test struct {
		key 	 string
		value    string
		expected string
		experr   string
	}

	tests := []test{
		{
			key:      "",
			value:    "",
			expected: "",
			experr:   "no authorization header",
		},
		{
			key:      "Authorization",
			value:    "",
			expected: "",
			experr:   "no authorization header",
		},
		{
			key:      "Authorization",
			value:    "Bearer some-api-key",
			expected: "",
			experr:   "malformed authorization header",
		},
		{
			key:      "Authorization",
			value:    "ApiKey some-api-key",
			expected: "some-api-key",
			experr:   "no authorization header",
		},
	}

	for i, testcase := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T){
			header := http.Header{}
			header.Add(testcase.key, testcase.value)
			
			result, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), testcase.experr) {
					return
				}
				t.Errorf("Unexpected error in TestGetAPIKey: %v\n", err)
				return
			}
			if result != testcase.expected {
				t.Errorf("Unexpected result in TestGetAPIKey: %v\n", result)
				return
			}
		})
	}
}