package eneru

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
)

type Client struct {
	url   string
	debug bool
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c *Client) Debug(debug bool) {
	c.debug = debug
}

func (c *Client) request(method, path string, query *url.Values, body io.Reader) (*http.Response, error) {
	r, _ := http.NewRequest(method, buildUrl(c.url, path, query), body)

	r.Header.Set("Content-Type", "application/json")

	if c.debug {
		c.dumpRequest(r)
	}

	resp, err := http.DefaultClient.Do(r)

	if c.debug {
		c.dumpResponse(resp)
	}

	return resp, err
}

// dumpRequest dumps the given HTTP request.
func (c *Client) dumpRequest(r *http.Request) {
	out, err := httputil.DumpRequestOut(r, true)
	if err == nil {
		log.Printf("%s\n\n", string(out))
	}
}

// dumpResponse dumps the given HTTP response.
func (c *Client) dumpResponse(resp *http.Response) {
	out, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Printf("%s\n\n", string(out))
	}
}
