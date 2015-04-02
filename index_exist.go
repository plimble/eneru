package eneru

// import ()

// type ExistIndexReq struct {
// 	client *Client
// 	index  string
// }

// func NewExistIndex(client *Client, index string) *ExistIndexReq {
// 	return &ExistIndexReq{
// 		client: client,
// 		index:  index,
// 	}
// }

// func (c *ExistIndexReq) Do(client *Client) (bool, error) {
// 	resp, err := client.request(HEAD, c.index, nil, nil)
// 	if err != nil {
// 		return false, err
// 	}

// 	if resp.StatusCode != 200 {
// 		return false, nil
// 	}

// 	return true, nil
// }
