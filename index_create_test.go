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
	ci     *CreateIndexReq
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

	client, err := NewClient(t.server.URL)
	if err != nil {
		panic(err)
	}

	t.client = client
}

func (t *CreateIndexSuite) TearDownSuite() {
	t.server.Close()
}

func (t *CreateIndexSuite) SetupTest() {
	t.ci = t.client.CreateIndex("test")
}

func (t *CreateIndexSuite) TestBody() {
	t.ci.Body(bytes.NewBuffer(nil))
	t.NotNil(t.ci.body)
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

	resp, err := t.ci.Body(j).Do()
	t.NoError(err)
	t.True(resp.Acknowledged, true)
}
