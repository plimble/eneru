package eneru

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UpdateSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestUpdateSuite(t *testing.T) {
	suite.Run(t, &UpdateSuite{})
}

func (t *UpdateSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodeResp(w, &UpdateResp{
			Created: false,
			Version: 2,
		})
	}))

	t.client, _ = NewClient(t.server.URL)
}

func (t *UpdateSuite) TearDownSuite() {
	t.server.Close()
}

func (t *UpdateSuite) TestBody() {
	req := t.client.Update("test", "book")
	req.Body(bytes.NewBuffer(nil))
	t.NotNil(req.body)
}

func (t *UpdateSuite) TestDo() {
	j := NewJson(func(j *Json) {
		j.AS("tags", "search", "computer")
	})

	resp, err := t.client.Update("test", "book").Id("1").Body(j).Do()
	t.NoError(err)
	t.False(resp.Created)
	t.NotEqual(resp.Version, 1)
}
