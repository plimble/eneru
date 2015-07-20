package eneru

import (
	"bytes"
	"encoding/json"
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

func (t *ClientSuite) TestSplitString() {
	sampleArrayInt := []int{
		45,
		2124,
	}

	sampleJson := map[string]interface{}{
		"Name":   "พลังงานไฟฟ้า",
		"Detail": "พลังงานเกิดจากแสงอาทิตย์",
		"ISBN":   12321342,
		"Tags": []string{
			"ชีวจิตสุขภาพ",
			"การเมืองที่ทำงาน",
			"งานบ้านออฟฟิตคอนโด",
		},
		"Codes": sampleArrayInt,
	}

	bj, err := json.Marshal(sampleJson)
	t.NoError(err)

	b, err := t.client.splitString(bytes.NewBuffer(bj))
	t.NoError(err)

	var data map[string]interface{}
	err = json.Unmarshal(b.Bytes(), &data)
	t.NoError(err)

	t.Equal("พลังงานไฟฟ้า", data["Name"])
	t.Equal("พลังงาน เกิด จาก แสงอาทิตย์", data["Detail"])

	dataInt := int(data["ISBN"].(float64))
	t.Equal(12321342, dataInt)

	setInt, _ := data["Codes"].([]int)
	for i, n := range setInt {
		t.Equal(sampleArrayInt[i], n)
	}

	resultString := []string{
		"ชีวจิต สุขภาพ",
		"การเมือง ที่ทำงาน",
		"งานบ้าน ออฟฟิต คอนโด",
	}
	setString, _ := data["Tags"].([]interface{})
	for i, n := range setString {
		t.Equal(resultString[i], n)
	}
}
