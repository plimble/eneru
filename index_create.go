package eneru

import (
	"bytes"
)

type CreateIndexReq struct {
	Query *Query
	body  *bytes.Buffer
	path  string
}

func NewCreateIndexReq(index string) *CreateIndexReq {
	return &CreateIndexReq{
		Query: NewQuery(),
		path:  index,
	}
}

func (c *CreateIndexReq) Body(body *bytes.Buffer) *CreateIndexReq {
	c.body = body

	return c
}

func (c *CreateIndexReq) do(client Client) (*CreateIndexResp, error) {
	resp, err := client.Request(PUT, c.path, c.Query, c.body)
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
