package eneru

import (
	"encoding/json"
	"github.com/plimble/utils/errors2"
	"io"
)

var (
	ErrUnableConnect = errors2.NewInternal("unable to connect elastic search")
)

type ErrorResp struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

func newErrResp(data io.Reader) error {
	errResp := &ErrorResp{}

	if err := json.NewDecoder(data).Decode(errResp); err != nil {
		return errors2.NewInternal(err.Error())
	}

	return errors2.NewError(errResp.Status, "", errResp.Err, errors2.Internal)
}

func (err *ErrorResp) Error() string {
	return err.Err
}
