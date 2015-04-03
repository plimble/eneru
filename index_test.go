package eneru

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type IndexSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestIndexSuite(t *testing.T) {
	suite.Run(t, &IndexSuite{})
}

func (t *IndexSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodeResp(w, &IndexResp{
			Created: true,
		})
	}))

	t.client, _ = NewClient(t.server.URL)
}

func (t *IndexSuite) TearDownSuite() {
	t.server.Close()
}

func (t *IndexSuite) TestBody() {
	req := t.client.Index("test")
	req.Body(bytes.NewBuffer(nil))
	t.NotNil(req.body)
}

func (t *IndexSuite) TestDo() {
	j := NewJson(func(j *Json) {
		j.S("title", "Elasticsearch: The Definitive Guide")
		j.S("isbn-10", "1449358543")
		j.AS("tags", "search", "computer")
	})

	resp, err := t.client.Index("test").Type("book").Body(j).Do()
	t.NoError(err)
	t.True(resp.Created, true)
}
