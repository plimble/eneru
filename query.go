package eneru

import (
	"bytes"
	"net/url"
)

type Query struct {
	buf      *bytes.Buffer
	firstKey bool
}

func NewQuery() *Query {
	q := &Query{
		buf:      &bytes.Buffer{},
		firstKey: true,
	}

	q.buf.WriteByte(0x3f)

	return q
}

func (q *Query) Add(key, val string) {
	if !q.firstKey {
		q.buf.WriteByte(0x26)
	}

	q.buf.WriteString(key)
	q.buf.WriteByte(0x3d)
	q.buf.WriteString(url.QueryEscape(val))
	q.firstKey = true
}

func (q *Query) String() string {
	return q.buf.String()
}
