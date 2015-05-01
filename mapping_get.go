package eneru

type GetMappingReq struct {
	client *Client
	index  string
	ty     string
}

func NewGetMapping(client *Client, index, ty string) *GetMappingReq {
	return &GetMappingReq{
		client: client,
		index:  index,
		ty:     ty,
	}
}

func (req *GetMappingReq) Do() (map[string]interface{}, error) {
	resp, err := req.client.Request(GET, buildPath(req.index, "_mapping", req.ty), nil, nil)
	if err != nil {
		return nil, err
	}

	var ret map[string]interface{}
	err = decodeResp(resp, &ret)
	return ret, err
}
