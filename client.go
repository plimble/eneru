package eneru

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/plimble/utils/pool"
	"net/http"
	"net/http/httputil"
	"time"
)

//go:generate mockery --name Client

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
)

var bufPool *pool.BufferPool

type Client struct {
	url        string
	debug      bool
	pretty     bool
	httpClient *http.Client
}

func NewClient(url string, poolSize int) (*Client, error) {
	c := &Client{
		url:        addTailingSlash(url),
		httpClient: http.DefaultClient,
	}

	var err error
	for i := 0; i < 5; i++ {
		log.Info("Try to connect elasticsearch...")
		if err = c.ping(); err != nil {
			log.Warnf("Try #%d: %s", i, err.Error())
		} else {
			log.Info("Elasticsearch Connected")
			break
		}

		time.Sleep(time.Second * 2)
	}

	if poolSize == 0 {
		poolSize = 512
	}
	bufPool = pool.NewBufferPool(poolSize)

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

func (c *Client) Delete(index string) *DeleteReq {
	return NewDelete(c, index)
}

func (c *Client) Index(index string) *IndexReq {
	return NewIndex(c, index)
}

func (c *Client) Update(index, ty string) *UpdateReq {
	return NewUpdate(c, index, ty)
}

func (c *Client) Search() *SearchReq {
	return NewSearch(c)
}

func (c *Client) UpdatePartial(index, ty, id string) *UpdatePartialReq {
	return NewUpdatePartial(c, index, ty, id)
}

func (c *Client) CreateMapping(index, ty string) *CreateMappingReq {
	return NewCreateMapping(c, index, ty)
}

func (c *Client) GetMapping(index, ty string) *GetMappingReq {
	return NewGetMapping(c, index, ty)
}

func (c *Client) DeleteMapping(index, ty string) *DeleteMappingReq {
	return NewDeleteMapping(c, index, ty)
}

func (c *Client) Debug(debug bool) {
	c.debug = debug
}

func (c *Client) Pretty(pretty bool) {
	c.pretty = pretty
}

func (c *Client) ping() error {
	resp, err := c.Request(GET, "", nil, nil)
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
	if err != nil {
		return nil, err
	}

	if c.debug {
		c.dumpResponse(resp)
	}

	if body != nil {
		bufPool.Put(body)
	}

	if err := c.checkResponse(resp); err != nil {
		return resp, err
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
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return newErrResp(resp.Body)
	}

	return nil
}
