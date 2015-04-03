package eneru

import (
	"fmt"
)

type DeleteReq struct {
	client *Client
	index  string
	ty     string
}

func NewDelete(client *Client) *DeleteReq {
	return &DeleteReq{
		client: client,
	}
}

func (req *DeleteReq) Index(index string) *DeleteReq {
	req.index = index

	return req
}

func (req *DeleteReq) Type(ty string) *DeleteReq {
	req.ty = ty

	return req
}

func (req *DeleteReq) Do() (*DeleteResp, error) {
	resp, err := req.client.Request(DELETE, buildPathIndexType(req.index, req.ty), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := &DeleteResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type DeleteResp struct {
	Acknowledged bool `json:"acknowledged"`
}
