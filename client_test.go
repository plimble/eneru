package eneru

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/plimble/tsplitter"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ClientSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}

func (t *ClientSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	t.client, _ = NewClient(t.server.URL, 512)
	t.client.tsplitterEnable(tsplitter.NewFileDict("./dictionary.txt"))
}

func getError(typeErr string) error {
	return errors.New(fmt.Sprintf("Failed %s not match", typeErr))
}

func generateSampleData() *bytes.Buffer {
	data := map[string]interface{}{
		"Name":   "พลังงานไฟฟ้า",
		"Detail": "พลังงานเกิดจากแสงอาทิตย์",
		"ISBN":   12321342,
		"Tags": []string{
			"ชีวจิตสุขภาพ",
			"การเมืองที่ทำงาน",
			"งานบ้านออฟฟิตคอนโด",
		},
		"Codes": []int{
			45,
			2124,
		},
	}

	bj, _ := json.Marshal(data)
	return bytes.NewBuffer(bj)
}

func checkSampleData(bj *bytes.Buffer) error {
	sampleArrayInt := []int{
		45,
		2124,
	}

	var data map[string]interface{}
	err := json.Unmarshal(bj.Bytes(), &data)
	if err != nil {
		return getError("map json")
	}

	if "พลังงานไฟฟ้า" != data["Name"] {
		return getError("string")
	}
	if "พลังงาน เกิด จาก แสงอาทิตย์" != data["Detail"] {
		return getError("text")
	}

	dataInt := int(data["ISBN"].(float64))
	if 12321342 != dataInt {
		return getError("int")
	}

	setInt, _ := data["Codes"].([]int)
	for i, n := range setInt {
		if sampleArrayInt[i] != n {
			return getError("array int")
		}
	}

	resultString := []string{
		"ชีวจิต สุขภาพ",
		"การเมือง ที่ทำงาน",
		"งานบ้าน ออฟฟิต คอนโด",
	}
	setString, _ := data["Tags"].([]interface{})
	for i, n := range setString {
		if resultString[i] != n {
			return getError("array string")
		}
	}

	return nil
}

func (t *ClientSuite) TestSplitString() {
	b, err := t.client.splitString(generateSampleData())
	t.NoError(err)
	t.NoError(checkSampleData(b))
}
