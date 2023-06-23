package tool

import "net/http"

type ErrorRaiserHTTPClient struct {
	err error
}

func (client ErrorRaiserHTTPClient) Get(url string) (*http.Response, error) {
	return nil, client.err
}

// HTTPClient does HTTP requests to remote servers
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

// HTTPClientImpl is the default implementation of HTTPClient
type HTTPClientImpl struct{}

// Get do a GET request and returns a *http.Response
func (HTTPClientImpl) Get(url string) (resp *http.Response, err error) {
	for i := 0; i < maxRetries; i++ {
		resp, err = http.Get(url)
		if err == nil {
			return resp, err
		}
	}

	return resp, err
}
