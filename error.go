package eneru

import (
	"encoding/json"
	"io"
)

type ErrorResp struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

func newErrResp(data io.Reader) error {
	errResp := &ErrorResp{}

	if err := json.NewDecoder(data).Decode(errResp); err != nil {
		return err
	}

	return errResp
}

func (err *ErrorResp) Error() string {
	return err.Err
}
