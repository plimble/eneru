package eneru

import (
	"bytes"
	"encoding/json"
)

type IndexReq struct {
	client *Client
	body   *bytes.Buffer
	index  string
	ty     string
	id     string
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

func (req *IndexReq) BodyJson(v interface{}) *IndexReq {
	req.body = bytes.NewBuffer(nil)
	json.NewEncoder(req.body).Encode(v)

	return req
}

func (req *IndexReq) Type(ty string) *IndexReq {
	req.ty = ty

	return req
}

func (req *IndexReq) ID(id string) *IndexReq {
	req.id = id

	return req
}

func (req *IndexReq) Do() (*IndexResp, error) {
	var err error
	ret := &IndexResp{}

	if req.client.tsplitter {
		req.body, err = splitString(req.client.dict, req.body)
		if err != nil {
			return ret, err
		}
	}

	resp, err := req.client.Request(PUT, buildPath(req.index, req.ty, req.id), nil, req.body)
	if err != nil {
		return nil, err
	}

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
