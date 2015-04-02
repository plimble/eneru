package eneru

import ()

type ExistIndexReq struct {
	client *Client
	path   string
}

func NewExistIndex(client *Client, index string) *ExistIndexReq {
	return &ExistIndexReq{
		client: client,
		path:   index,
	}
}

func (req *ExistIndexReq) Do() (bool, error) {
	resp, err := req.client.Request(HEAD, req.path, nil, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}
