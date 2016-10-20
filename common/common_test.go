package common

import (
	"net/http"
	"testing"
)

func TestCommonHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	if req.Header.Get("Content-Type") != "" {
		t.Fail()
	}
	CommonHeaders(*req, "token", "application/x-www-form-urlencoded")
	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Fail()
	}
}
