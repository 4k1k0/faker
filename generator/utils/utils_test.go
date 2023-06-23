package utils

import (
	"net/http"
	"strings"
	"testing"

	"github.com/jaswdr/faker/tool"
)

func TestHTTPClientImplCanDoGetRequests(t *testing.T) {
	client := HTTPClientImpl{}
	resp, err := client.Get("https://www.google.com")
	tool.Expect(t, err, nil)
	tool.Expect(t, resp.StatusCode, http.StatusOK)
}

func TestHTTPClientImplReturnsErrorWhenRequestFails(t *testing.T) {
	client := HTTPClientImpl{}
	_, err := client.Get("https://invalid")
	tool.NotExpect(t, err, nil)
}

func TestTempFileCreatorImplCanCreateTempFiles(t *testing.T) {
	creator := TempFileCreatorImpl{}
	f, err := creator.TempFile("prefix")
	tool.Expect(t, err, nil)
	tool.Expect(t, true, strings.Contains(f.Name(), "prefix"))
	tool.Expect(t, f.Close(), nil)
}
