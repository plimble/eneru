package eneru

import ()

type DeleteIndexReq struct {
	client *Client
	index  string
}

func NewDeleteIndex(client *Client, index string) *DeleteIndexReq {
	return &DeleteIndexReq{
		client: client,
		index:  index,
	}
}

func (req *DeleteIndexReq) Do() (*DeleteIndexResp, error) {
	resp, err := req.client.Request(DELETE, req.index, nil, nil)
	if err != nil {
		return nil, err
	}

	ret := &DeleteIndexResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type DeleteIndexResp struct {
	Acknowledged bool `json:"acknowledged"`
}
