package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	headers := make(http.Header)
	// headers.Add("Authorization", "ApiKey your-api-key-here")

	got, _ := GetAPIKey(headers)
	want, _ := "", ErrNoAuthHeaderIncluded
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
