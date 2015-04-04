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
			Index:   "test",
			Type:    "book",
			ID:      "1",
			Version: 2,
			Created: false,
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

	resp, err := t.client.Update("test", "book").ID("1").Body(j).Do()
	t.NoError(err)
	t.Equal(resp.Index, "test")
	t.Equal(resp.Type, "book")
	t.Equal(resp.ID, "1")
	t.NotEqual(resp.Version, 1)
	t.False(resp.Created)
}
