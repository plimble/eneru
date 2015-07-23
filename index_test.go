package eneru

import (
	"bytes"
	"github.com/plimble/tsplitter"
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
			Index:   "test",
			Type:    "book",
			ID:      "1",
			Version: 1,
			Created: true,
		})
	}))

	t.client, _ = NewClient(t.server.URL, 512)
}

func (t *IndexSuite) TearDownSuite() {
	t.server.Close()
}

func (t *IndexSuite) TestBody() {
	req := t.client.Index("test")
	req.Body(bytes.NewBuffer(nil))
	t.NotNil(req.body)
}

func (t *IndexSuite) TestBodyTsplitter() {
	t.client.tsplitterEnable(tsplitter.NewFileDict("./dictionary.txt"))
	req := t.client.Index("test")
	req.Body(generateSampleData())
	t.NotNil(req.body)
	t.NoError(checkSampleData(req.body))
}

func (t *IndexSuite) TestDo() {
	j := NewJson(func(j *Json) {
		j.S("title", "Elasticsearch: The Definitive Guide")
		j.S("isbn-10", "1449358543")
		j.AS("tags", "search", "computer")
	})

	resp, err := t.client.Index("test").Type("book").ID("1").Body(j).Do()
	t.NoError(err)
	t.Equal(resp.Index, "test")
	t.Equal(resp.Type, "book")
	t.Equal(resp.ID, "1")
	t.Equal(resp.Version, 1)
	t.True(resp.Created, true)
}
