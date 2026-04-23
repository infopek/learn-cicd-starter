package auth

import (
	"testing"
	"net/http"
	"reflect"
)

func TestGetAPIKey(t *testing.T) {
	validHeader := http.Header{}
	validHeader.Add("Authorization", "ApiKey v89hp98hj23pj")

	invalidHeader1 := http.Header{}
	invalidHeader1.Add("Authentication", "ApiKey randomkey212")

	tests := map[string]struct {
		input   http.Header
		isError bool
		want    string
	}{
		"successful apikey": { input: validHeader, isError: false, want: "v89hp98hj23pj"},
		"invalid apikey": { input: invalidHeader1, isError: true, want: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != nil {
				if !tc.isError {
					t.Fatalf("expected: %v, got an unexpected error: %v", tc.want, err)
				}
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
