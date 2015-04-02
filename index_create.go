package eneru

import (
	"bytes"
)

type CreateIndexReq struct {
	client *Client
	Query  *Query
	body   *bytes.Buffer
	path   string
}

func NewCreateIndex(client *Client, index string) *CreateIndexReq {
	return &CreateIndexReq{
		client: client,
		Query:  NewQuery(),
		path:   index,
	}
}

func (c *CreateIndexReq) Body(body *bytes.Buffer) *CreateIndexReq {
	c.body = body

	return c
}

func (c *CreateIndexReq) Do() (*CreateIndexResp, error) {
	resp, err := c.client.Request(PUT, c.path, c.Query, c.body)
	if err != nil {
		return nil, err
	}

	ret := &CreateIndexResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type CreateIndexResp struct {
	Acknowledged bool `json:"acknowledged"`
}
