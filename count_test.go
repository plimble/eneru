package eneru

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CountSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestCountSuite(t *testing.T) {
	suite.Run(t, &CountSuite{})
}

func (t *CountSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodeResp(w, &CountResp{
			Count: 10,
		})
	}))

	client, err := NewClient(t.server.URL)
	if err != nil {
		panic(err)
	}

	t.client = client
}

func (t *CountSuite) TearDownSuite() {
	t.server.Close()
}

func (t *CountSuite) TestGetURL() {
	req := t.client.Count().Index("test")
	t.Equal("test/_count", req.getURL())

	req = t.client.Count().Index("test").Type("user")
	t.Equal("test/user/_count", req.getURL())

	req = t.client.Count()
	t.Equal("_count", req.getURL())
}

func (t *CountSuite) TestDo() {
	count, err := t.client.Count().Index("test").Type("user").Do()
	t.NoError(err)
	t.Equal(10, count)
}
