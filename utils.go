package eneru

import (
	"github.com/plimble/utils/strings2"
	"net/http"
	"net/url"
)

func buildUrl(url, path string, query *url.Values) string {
	if query != nil {
		return string2.ConcatBase(url, "/", path, "?", query.Encode())
	}

	return string2.ConcatBase(url, "/", path)
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}

	return newErrResp(resp.Body)
}
