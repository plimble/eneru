package eneru

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SearchSuite struct {
	suite.Suite
	server *httptest.Server
	client *Client
}

func TestSearchSuite(t *testing.T) {
	suite.Run(t, &SearchSuite{})
}

func (t *SearchSuite) SetupSuite() {
	t.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			{
			    "took": 20,
			    "timed_out": false,
			    "hits": {
			        "total": 2,
			        "max_score": 1,
			        "hits": [
			            {
			                "_index": "test",
			                "_type": "user",
			                "_id": "1",
			                "_score": 1,
			                "_source": {
			                    "username": "test",
			                    "age": 10,
			                    "score": 10.125,
			                    "scope": [
			                        "a",
			                        "b",
			                        "c"
			                    ],
			                    "latlon": [
			                        10,
			                        20
			                    ],
			                    "created_at": "2009-11-10 23:00:00 +0000 UTC"
			                }
			            },
			            {
			                "_index": "test",
			                "_type": "user",
			                "_id": "2",
			                "_score": 1,
			                "_source": {
			                    "username": "user2",
			                    "age": 20,
			                    "score": 30.333,
			                    "scope": [
			                        "a",
			                        "b",
			                        "d"
			                    ],
			                    "latlon": [
			                        30,
			                        30
			                    ],
			                    "created_at": "2009-11-11 23:00:00 +0000 UTC"
			                }
			            }
			        ]
			    }
			}
		`))
	}))

	t.client, _ = NewClient(t.server.URL, 512)
}

func (t *SearchSuite) TearDownSuite() {
	t.server.Close()
}

func (t *SearchSuite) TestBody() {
	req := t.client.Search()
	req.Body(bytes.NewBuffer(nil))
	t.NotNil(req.body)
}

func (t *SearchSuite) TestDo() {
	j := NewJson(func(j *Json) {
	})

	resp, err := t.client.Search().Index("test").Type("user").Body(j).Do()
	t.NoError(err)
	t.Equal(int(resp.Took), 20)
	t.Equal(resp.TimedOut, false)
	t.Equal(resp.Hits.Total, 2)
	t.Equal(int(resp.Hits.MaxScore), 1)
	t.Len(resp.Hits.Hits, 2)
	t.Equal(resp.Hits.Hits[0].ID, "1")
	t.Equal(resp.Hits.Hits[1].ID, "2")
}
