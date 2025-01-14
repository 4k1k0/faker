package faker

import (
	"net/http"
	"os"
	"strings"
	"testing"
)

type ErrorRaiserHTTPClient struct {
	err error
}

func (client ErrorRaiserHTTPClient) Get(url string) (*http.Response, error) {
	return nil, client.err
}

type ErrorRaiserTempFileCreator struct {
	err error
}

func (creator ErrorRaiserTempFileCreator) TempFile(prefix string) (*os.File, error) {
	return nil, creator.err
}

func TestHTTPClientImplCanDoGetRequests(t *testing.T) {
	client := HTTPClientImpl{}
	resp, err := client.Get("https://www.google.com")
	Expect(t, err, nil)
	Expect(t, resp.StatusCode, http.StatusOK)
}

func TestHTTPClientImplReturnsErrorWhenRequestFails(t *testing.T) {
	client := HTTPClientImpl{}
	_, err := client.Get("https://invalid")
	NotExpect(t, err, nil)
}

func TestTempFileCreatorImplCanCreateTempFiles(t *testing.T) {
	creator := TempFileCreatorImpl{}
	f, err := creator.TempFile("prefix")
	Expect(t, err, nil)
	Expect(t, true, strings.Contains(f.Name(), "prefix"))
	Expect(t, f.Close(), nil)
}
