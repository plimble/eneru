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

func decodeResp(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}

func encodeResp(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
}

func buildPathIndexType(index, ty string) string {
	switch {
	case index != "" && ty == "":
		return string2.Concat(index)
	case index != "" && ty != "":
		return string2.Concat(index, "/", ty)
	}

	return string2.Concat(index, "/", ty)
}

func buildPathIndexTypeAction(index, ty, action string) string {
	switch {
	case index == "" && ty == "" && action != "":
		return action
	case index != "" && ty == "" && action != "":
		return string2.Concat(index, "/", action)
	}

	return string2.Concat(index, "/", ty, "/", action)
}
