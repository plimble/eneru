package eneru

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
)

//go:generate mockery --name Client

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
)

type Client struct {
	url        string
	debug      bool
	pretty     bool
	httpClient *http.Client
}

func NewClient(url string) (*Client, error) {
	c := &Client{
		url:        addTailingSlash(url),
		httpClient: http.DefaultClient,
	}

	if err := c.ping(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) CreateIndex(index string) *CreateIndexReq {
	return NewCreateIndex(c, index)
}

func (c *Client) ExistIndex(index string) *ExistIndexReq {
	return NewExistIndex(c, index)
}

func (c *Client) Count() *CountReq {
	return NewCount(c)
}

func (c *Client) Debug(debug bool) {
	c.debug = debug
}

func (c *Client) Pretty(pretty bool) {
	c.pretty = pretty
}

func (c *Client) ping() error {
	resp, err := c.Request(HEAD, "/", nil, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return ErrUnableConnect
	}

	return nil
}

func (c *Client) Request(method, path string, query *Query, body *bytes.Buffer) (*http.Response, error) {
	c.doPretty(query, body)

	r := c.buildRequest(method, path, query, body)
	if c.debug {
		c.dumpRequest(r)
	}

	resp, err := c.httpClient.Do(r)
	if c.debug {
		c.dumpResponse(resp)
	}

	if err := c.checkResponse(resp); err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) buildRequest(method, path string, query *Query, body *bytes.Buffer) *http.Request {
	var r *http.Request
	if body == nil {
		r, _ = http.NewRequest(method, buildUrl(c.url, path, query), nil)
	} else {
		r, _ = http.NewRequest(method, buildUrl(c.url, path, query), body)
	}

	r.Header.Set("Content-Type", "application/json")

	return r
}

func (c *Client) doPretty(query *Query, body *bytes.Buffer) {
	if c.pretty {
		if query == nil {
			query = NewQuery()
		}
		query.Add("pretty", "true")

		if body != nil {
			data := make([]byte, body.Len())
			copy(data, body.Bytes())
			body.Truncate(0)
			json.Indent(body, data, "", "    ")
		}
	}
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

func (c *Client) checkResponse(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}

	return newErrResp(resp.Body)
}
