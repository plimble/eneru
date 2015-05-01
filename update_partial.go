package eneru

import (
	"bytes"
)

type UpdatePartialReq struct {
	client *Client
	body   *bytes.Buffer
	index  string
	ty     string
	id     string
	Query  *Query
}

func NewUpdatePartial(client *Client, index, ty, id string) *UpdatePartialReq {
	return &UpdatePartialReq{
		client: client,
		index:  index,
		ty:     ty,
		id:     id,
	}
}

func (req *UpdatePartialReq) Body(body *bytes.Buffer) *UpdatePartialReq {
	req.body = body

	return req
}

func (req *UpdatePartialReq) Do() (*UpdatePartialResp, error) {
	resp, err := req.client.Request(POST, buildPath(req.index, req.ty, req.id, "_update"), req.Query, req.body)
	if err != nil {
		return nil, err
	}

	ret := &UpdatePartialResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type UpdatePartialResp struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
}
