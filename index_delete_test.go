package eneru

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DeleteIndexSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestDeleteIndexSuite(t *testing.T) {
	suite.Run(t, &DeleteIndexSuite{})
}

func (t *DeleteIndexSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodeResp(w, &DeleteIndexResp{
			Acknowledged: true,
		})
	}))

	client, err := NewClient(t.server.URL)
	if err != nil {
		panic(err)
	}

	t.client = client
}

func (t *DeleteIndexSuite) TearDownSuite() {
	t.server.Close()
}

func (t *DeleteIndexSuite) TestDo() {
	resp, err := t.client.DeleteIndex("test", "").Do()

	t.NoError(err)
	t.True(resp.Acknowledged)
}
