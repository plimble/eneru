package eneru

import (
	"encoding/json"
	"github.com/plimble/utils/strings2"
	"net/http"
	"strings"
)

func buildUrl(url, path string, query *Query) string {
	if query == nil {
		return string2.ConcatBase(url, path)
	}

	return string2.ConcatBase(url, path, query.String())
}

func addTailingSlash(url string) string {
	if !strings.HasSuffix(url, "/") {
		return string2.ConcatBase(url, "/")
	}

	return url
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}

	return newErrResp(resp.Body)
}

func decodeResp(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func encodeResp(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
}
