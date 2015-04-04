package eneru

import (
	"bytes"
)

type UpdateReq struct {
	client *Client
	body   *bytes.Buffer
	index  string
	ty     string
	id     string
	Query  *Query
}

func NewUpdate(client *Client, index, ty string) *UpdateReq {
	return &UpdateReq{
		client: client,
		index:  index,
		ty:     ty,
	}
}

func (req *UpdateReq) Body(body *bytes.Buffer) *UpdateReq {
	req.body = body

	return req
}

func (req *UpdateReq) ID(id string) *UpdateReq {
	req.id = id

	return req
}

func (req *UpdateReq) Do() (*UpdateResp, error) {
	resp, err := req.client.Request(PUT, buildPath(req.index, req.ty), req.Query, req.body)
	if err != nil {
		return nil, err
	}

	ret := &UpdateResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type UpdateResp struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Created bool   `json:"created"`
}
