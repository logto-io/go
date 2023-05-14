package client

import "net/http"

type CustomTransport struct {
	appId     string
	appSecret string
	http.RoundTripper
}

func (ct *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if ct.appSecret != "" {
		req.SetBasicAuth(ct.appId, ct.appSecret)
	}
	return ct.RoundTripper.RoundTrip(req)
}
