package hextree

import (
	"strings"

	http "github.com/bogdanfinn/fhttp"

	tlsclient "github.com/bogdanfinn/tls-client"
)

// Client - holds our PX payload struct, proxy, device, and HTTP client
type Client struct {
	HTTPClient tlsclient.HttpClient
}

// MakeClient - makes a *Client struct
func MakeClient() (*Client, error) {
	opts := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(30),
		tlsclient.WithInsecureSkipVerify(),
		tlsclient.WithClientProfile(tlsclient.Chrome_112),
	}
	h, err := tlsclient.NewHttpClient(nil, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		HTTPClient: h,
	}, nil
}

// FormatHeaders turns a string of headers seperated by `|` into a http.Header map
func (c *Client) FormatHeaders(h string) http.Header {
	headers := http.Header{}
	for _, header := range strings.Split(h, "|") {
		parts := strings.Split(header, ": ")
		headers.Set(parts[0], parts[1])
	}
	return headers
}
