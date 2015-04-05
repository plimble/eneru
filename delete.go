package eneru

type DeleteReq struct {
	client *Client
	index  string
	ty     string
	id     string
}

func NewDelete(client *Client, index string) *DeleteReq {
	return &DeleteReq{
		client: client,
		index:  index,
	}
}

func (req *DeleteReq) Type(ty string) *DeleteReq {
	req.ty = ty

	return req
}

func (req *DeleteReq) ID(id string) *DeleteReq {
	req.id = id

	return req
}

func (req *DeleteReq) Do() (*DeleteResp, error) {
	resp, err := req.client.Request(DELETE, buildPath(req.index, req.ty, req.id), nil, nil)
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
