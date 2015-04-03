package eneru

type CountReq struct {
	client *Client
	index  string
	ty     string
}

func NewCount(client *Client) *CountReq {
	return &CountReq{
		client: client,
	}
}

func (req *CountReq) Index(index string) *CountReq {
	req.index = index

	return req
}

func (req *CountReq) Type(ty string) *CountReq {
	req.ty = ty

	return req
}

func (req *CountReq) getURL() string {
	return buildPathIndexTypeAction(req.index, req.ty, "_count")
}

func (req *CountReq) Do() (int, error) {
	resp, err := req.client.Request(GET, req.getURL(), nil, nil)
	if err != nil {
		return 0, err
	}

	ret := &CountResp{}
	err = decodeResp(resp, ret)
	return ret.Count, err
}

type CountResp struct {
	Count int `json:"count"`
}
