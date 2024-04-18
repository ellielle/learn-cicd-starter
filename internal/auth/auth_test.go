package auth

import (
	"net/http"
	"strings"
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("Request bad")
	}

	tests := map[string]struct {
		header string
		want   string
	}{
		"splits normal": {header: "Authorization ApiKey 8913ijoij", want: "8913ijoij"},
		"no api key":    {header: "Authorization ApiKey ", want: ""},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			req.Header.Del("Authorization")
			split := splitHeader(tc.header)
			req.Header.Add(split[0], split[1]+" "+split[2])
			got, err := GetAPIKey(req.Header)
			if err != nil {
				t.Fatal(err)
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

// splitHeader is a helper function to add headers to the test request
func splitHeader(header string) []string {
	split := strings.Split(header, " ")
	return split
}
