package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}

// func Test(t *testing.T) {
// 	testCases := []struct {
// 		name	string
// 		input	http.Header
// 		want	string
// 	}{
// 		{name: "Correct format", input: http.Header{"Authorization": []string{"ApiKey abc123"}}, want: "abc123"},
// 		{name: "Incorrect format", input: http.Header{"Authorization": []string{"Bearer abc123"}}, want: ""},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.name, func(t *testing.T) {
// 			got, _ := GetAPIKey(tC.input)
// 			if !reflect.DeepEqual(tC.want, got) {
// 				t.Fatalf("expected %s, but got %s", tC.want, got)
// 			}
// 		})
// 	}
//}
