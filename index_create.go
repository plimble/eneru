package eneru

import (
	"encoding/json"
	"io"
	"net/url"
)

type CreateIndexReq struct {
	client *Client
	query  *url.Values
	body   io.Reader
	index  string
}

func NewCreateIndex(client *Client, index string) *CreateIndexReq {
	return &CreateIndexReq{
		client: client,
		query:  &url.Values{},
		index:  index,
	}
}

func (c *CreateIndexReq) Body(body io.Reader) *CreateIndexReq {
	c.body = body

	return c
}

func (c *CreateIndexReq) Pretty(pretty bool) *CreateIndexReq {
	if pretty {
		c.query.Set("pretty", "true")
	}

	return c
}

func (c *CreateIndexReq) Do() (*CreateIndexResp, error) {
	resp, err := c.client.request(PUT, c.index, c.query, c.body)
	if err != nil {
		return nil, err
	}

	if err := checkResponse(resp); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ret := &CreateIndexResp{}
	if err := json.NewDecoder(resp.Body).Decode(ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type CreateIndexResp struct {
	Acknowledged bool `json:"acknowledged"`
}
