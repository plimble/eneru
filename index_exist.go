package eneru

import ()

type ExistIndexReq struct {
	client *Client
	path   string
	ty     string
}

func NewExistIndex(client *Client, index string) *ExistIndexReq {
	return &ExistIndexReq{
		client: client,
		path:   index,
	}
}

func (req *ExistIndexReq) Type(ty string) *ExistIndexReq {
	req.ty = ty
	return req
}

func (req *ExistIndexReq) Do() (bool, error) {
	resp, err := req.client.Request(HEAD, buildPath(req.path, req.ty), nil, nil)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, nil
	}

	return true, nil
}
