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

	client, err := NewClient(t.server.URL, 512)
	if err != nil {
		panic(err)
	}

	t.client = client
}

func (t *CountSuite) TearDownSuite() {
	t.server.Close()
}

func (t *CountSuite) TestDo() {
	count, err := t.client.Count().Index("test").Type("user").Do()
	t.NoError(err)
	t.Equal(10, count)
}
