package eneru

type DeleteMappingReq struct {
	client *Client
	index  string
	ty     string
}

func NewDeleteMapping(client *Client, index, ty string) *DeleteMappingReq {
	return &DeleteMappingReq{
		client: client,
		index:  index,
		ty:     ty,
	}
}

func (req *DeleteMappingReq) Do() (*DeleteMappingResp, error) {
	resp, err := req.client.Request(DELETE, buildPath(req.index, "_mapping", req.ty), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := &DeleteMappingResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type DeleteMappingResp struct {
	Acknowledged bool `json:"acknowledged"`
}
