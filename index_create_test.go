package eneru

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CreateIndexSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestCreateIndexSuite(t *testing.T) {
	suite.Run(t, &CreateIndexSuite{})
}

func (t *CreateIndexSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodeResp(w, &CreateIndexResp{
			Acknowledged: true,
		})
	}))

	t.client, _ = NewClient(t.server.URL)
}

func (t *CreateIndexSuite) TearDownSuite() {
	t.server.Close()
}

func (t *CreateIndexSuite) TestBody() {
	req := t.client.CreateIndex("test")
	req.Body(bytes.NewBuffer(nil))
	t.NotNil(req.body)
}

func (t *CreateIndexSuite) TestDo() {
	j := NewJson(func(j *Json) {
		j.O("mappings", func(j *Json) {
			j.O("book", func(j *Json) {
				j.O("properties", func(j *Json) {
					j.O("name", func(j *Json) {
						j.S("type", "string")
					})
					j.O("email", func(j *Json) {
						j.S("type", "string")
						j.S("index", "not_analyzed")
					})
				})
			})
		})
	})

	resp, err := t.client.CreateIndex("test").Body(j).Do()
	t.NoError(err)
	t.True(resp.Acknowledged, true)
}
