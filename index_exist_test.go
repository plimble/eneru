package eneru

import (
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ExistIndexSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestExistIndexSuite(t *testing.T) {
	suite.Run(t, &ExistIndexSuite{})
}

func (t *ExistIndexSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))

	client, err := NewClient(t.server.URL)
	if err != nil {
		panic(err)
	}

	t.client = client
}

func (t *ExistIndexSuite) TearDownSuite() {
	t.server.Close()
}

func (t *ExistIndexSuite) TestDo() {
	exist, err := t.client.ExistIndex("test").Do()
	t.NoError(err)
	t.True(exist)
}
