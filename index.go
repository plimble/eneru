package eneru

import (
	"bytes"
)

type IndexReq struct {
	client *Client
	body   *bytes.Buffer
	index  string
	ty     string
}

func NewIndex(client *Client, index string) *IndexReq {
	return &IndexReq{
		client: client,
		index:  index,
	}
}

func (req *IndexReq) Body(body *bytes.Buffer) *IndexReq {
	req.body = body

	return req
}

func (req *IndexReq) Type(ty string) *IndexReq {
	req.ty = ty

	return req
}

func (req *IndexReq) Do() (*IndexResp, error) {
	resp, err := req.client.Request(PUT, buildPath(req.index, req.ty), nil, req.body)
	if err != nil {
		return nil, err
	}

	ret := &IndexResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type IndexResp struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Created bool   `json:"created"`
}
