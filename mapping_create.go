package eneru

import (
	"bytes"
)

type CreateMappingReq struct {
	client *Client
	body   *bytes.Buffer
	index  string
	ty     string
}

func NewCreateMapping(client *Client, index, ty string) *CreateMappingReq {
	return &CreateMappingReq{
		client: client,
		index:  index,
		ty:     ty,
	}
}

func (req *CreateMappingReq) Body(body *bytes.Buffer) *CreateMappingReq {
	req.body = body

	return req
}

func (req *CreateMappingReq) Do() (*CreateMappingResp, error) {
	resp, err := req.client.Request(PUT, buildPath(req.index, "_mapping", req.ty), nil, req.body)
	if err != nil {
		return nil, err
	}

	ret := &CreateMappingResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type CreateMappingResp struct {
	Acknowledged bool `json:"acknowledged"`
}
