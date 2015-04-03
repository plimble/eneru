package eneru

import (
	"bytes"
	"encoding/json"
)

type SearchReq struct {
	client *Client
	index  string
	ty     string
	query  *Query
	body   *bytes.Buffer
}

func NewSearch(client *Client) *SearchReq {
	return &SearchReq{
		client: client,
		query:  NewQuery(),
	}
}

func (req *SearchReq) Index(index string) *SearchReq {
	req.index = index

	return req
}

func (req *SearchReq) Type(ty string) *SearchReq {
	req.ty = ty

	return req
}

func (req *SearchReq) Source(s string) *SearchReq {
	req.query.Add("_source", s)

	return req
}

func (req *SearchReq) Analyzer(a string) *SearchReq {
	req.query.Add("analyzer", a)

	return req
}

func (req *SearchReq) Body(body *bytes.Buffer) *SearchReq {
	req.body = body

	return req
}

func (req *SearchReq) Do() (*SearchResp, error) {
	resp, err := req.client.Request(POST, buildPath(req.index, req.ty, "_search"), req.query, req.body)
	if err != nil {
		return nil, err
	}

	ret := &SearchResp{}
	err = decodeResp(resp, ret)
	return ret, err
}

type SearchResp struct {
	TookInMillis int64         `json:"took"`            // search time in milliseconds
	ScrollId     string        `json:"_scroll_id"`      // only used with Scroll and Scan operations
	Hits         *SearchHits   `json:"hits"`            // the actual search hits
	Suggest      SearchSuggest `json:"suggest"`         // results from suggesters
	Facets       SearchFacets  `json:"facets"`          // results from facets
	Aggregations Aggregations  `json:"aggregations"`    // results from aggregations
	TimedOut     bool          `json:"timed_out"`       // true if the search timed out
	Error        string        `json:"error,omitempty"` // used in MultiSearch only
}

type SearchHits struct {
	TotalHits int64        `json:"total"`     // total number of hits found
	MaxScore  *float64     `json:"max_score"` // maximum score of all hits
	Hits      []*SearchHit `json:"hits"`      // the actual hits returned
}

type SearchSuggest map[string][]*SearchSuggestion

type SearchSuggestion struct {
	Text    string                    `json:"text"`
	Offset  int                       `json:"offset"`
	Length  int                       `json:"length"`
	Options []*SearchSuggestionOption `json:"options"`
}

type SearchSuggestionOption struct {
	Text    string      `json:"text"`
	Score   float32     `json:"score"`
	Freq    int         `json:"freq"`
	Payload interface{} `json:"payload"`
}

type SearchFacets map[string]*SearchFacet

type SearchFacet struct {
	Type    string              `json:"_type"`
	Missing int                 `json:"missing"`
	Total   int                 `json:"total"`
	Other   int                 `json:"other"`
	Terms   []*SearchFacetTerm  `json:"terms"`
	Ranges  []*SearchFacetRange `json:"ranges"`
	Entries []*SearchFacetEntry `json:"entries"`
}

type SearchFacetTerm struct {
	Term  string `json:"term"`
	Count int    `json:"count"`
}

type SearchFacetRange struct {
	From       *float64 `json:"from"`
	FromStr    *string  `json:"from_str"`
	To         *float64 `json:"to"`
	ToStr      *string  `json:"to_str"`
	Count      int      `json:"count"`
	Min        *float64 `json:"min"`
	Max        *float64 `json:"max"`
	TotalCount int      `json:"total_count"`
	Total      *float64 `json:"total"`
	Mean       *float64 `json:"mean"`
}

type SearchFacetEntry struct {
	// Key for this facet, e.g. in histograms
	Key interface{} `json:"key"`
	// Date histograms contain the number of milliseconds as date:
	// If e.Time = 1293840000000, then: Time.at(1293840000000/1000) => 2011-01-01
	Time int64 `json:"time"`
	// Number of hits for this facet
	Count int `json:"count"`
	// Min is either a string like "Infinity" or a float64.
	// This is returned with some DateHistogram facets.
	Min interface{} `json:"min,omitempty"`
	// Max is either a string like "-Infinity" or a float64
	// This is returned with some DateHistogram facets.
	Max interface{} `json:"max,omitempty"`
	// Total is the sum of all entries on the recorded Time
	// This is returned with some DateHistogram facets.
	Total float64 `json:"total,omitempty"`
	// TotalCount is the number of entries for Total
	// This is returned with some DateHistogram facets.
	TotalCount int `json:"total_count,omitempty"`
	// Mean is the mean value
	// This is returned with some DateHistogram facets.
	Mean float64 `json:"mean,omitempty"`
}

type Aggregations map[string]*json.RawMessage

type SearchHit struct {
	Score       *float64               `json:"_score"`       // computed score
	Index       string                 `json:"_index"`       // index name
	Id          string                 `json:"_id"`          // external or internal
	Type        string                 `json:"_type"`        // type
	Version     *int64                 `json:"_version"`     // version number, when Version is set to true in SearchService
	Sort        []interface{}          `json:"sort"`         // sort information
	Highlight   SearchHitHighlight     `json:"highlight"`    // highlighter information
	Source      *json.RawMessage       `json:"_source"`      // stored document source
	Fields      map[string]interface{} `json:"fields"`       // returned fields
	Explanation *SearchExplanation     `json:"_explanation"` // explains how the score was computed

}

type SearchHitHighlight map[string][]string

type SearchExplanation struct {
	Value       float64              `json:"value"`             // e.g. 1.0
	Description string               `json:"description"`       // e.g. "boost" or "ConstantScore(*:*), product of:"
	Details     []*SearchExplanation `json:"details,omitempty"` // recursive details
}
