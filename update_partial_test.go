package eneru

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UpdatePartialSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestUpdatePartialSuite(t *testing.T) {
	suite.Run(t, &UpdatePartialSuite{})
}

func (t *UpdatePartialSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encodeResp(w, &UpdatePartialResp{
			Index:   "test",
			Type:    "book",
			ID:      "1",
			Version: 2,
		})
	}))

	t.client, _ = NewClient(t.server.URL, 512)
}

func (t *UpdatePartialSuite) TearDownSuite() {
	t.server.Close()
}

func (t *UpdatePartialSuite) TestBody() {
	req := t.client.UpdatePartial("test", "book")
	req.Body(bytes.NewBuffer(nil))
	t.NotNil(req.body)
}

func (t *UpdatePartialSuite) TestDo() {
	j := NewJson(func(j *Json) {
		j.AS("tags", "search", "computer")
	})

	resp, err := t.client.UpdatePartial("test", "book").ID("1").Body(j).Do()
	t.NoError(err)
	t.Equal(resp.Index, "test")
	t.Equal(resp.Type, "book")
	t.Equal(resp.ID, "1")
	t.NotEqual(resp.Version, 1)
}
