package eneru

import (
	"bytes"
	"encoding/json"
	"github.com/plimble/tsplitter"
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

func buildPath(a ...string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := 0
	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			n += len(a[i]) + 1
		}
	}

	b := make([]byte, n)
	bp := 0
	for _, s := range a {
		if s != "" {
			bp += copy(b[bp:], s)
			bp += copy(b[bp:], "/")
		}
	}

	return string(b)
}

func splitString(dict *tsplitter.FileDict, data *bytes.Buffer) (*bytes.Buffer, error) {
	var mapJson map[string]interface{}
	err := json.Unmarshal(data.Bytes(), &mapJson)
	if err != nil {
		return nil, err
	}

	for index, value := range mapJson {
		switch v := value.(type) {
		case string:
			mapJson[index] = strings.Join(tsplitter.Split(dict, value.(string)).All(), " ")
		case []interface{}:
			for i, n := range value.([]interface{}) {
				str, ok := n.(string)
				if !ok {
					break
				}
				v[i] = strings.Join(tsplitter.Split(dict, str).All(), " ")
			}
		}
	}
	splitData, err := json.Marshal(mapJson)
	return bytes.NewBuffer(splitData), err
}
