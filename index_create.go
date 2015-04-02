package eneru

import (
	"bytes"
)

type CreateIndexReq struct {
	client *Client
	body   *bytes.Buffer
	path   string
}

func NewCreateIndex(client *Client, index string) *CreateIndexReq {
	return &CreateIndexReq{
		client: client,
		path:   index,
	}
}

func (req *CreateIndexReq) Body(body *bytes.Buffer) *CreateIndexReq {
	req.body = body

	return req
}

func (req *CreateIndexReq) Do() (*CreateIndexResp, error) {
	resp, err := req.client.Request(PUT, req.path, nil, req.body)
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
