package eneru

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type JsonSuite struct {
	suite.Suite
}

func TestJsonSuite(t *testing.T) {
	suite.Run(t, &JsonSuite{})
}

func (t *JsonSuite) TestJson() {

	j := NewJson(func(j *Json) {
		j.O("mappings", func(j *Json) {
			j.O("book", func(j *Json) {
				j.O("properties", func(j *Json) {
					j.O("name", func(j *Json) {
						j.S("type", "string")
						j.I("max", 10000)
					})
					j.O("email", func(j *Json) {
						j.S("type", "string")
						j.S("index", "not_analyzed")
					})
				})
			})
		})
	})

	expJson := "{\"mappings\":{\"book\":{\"properties\":{\"name\":{\"type\":\"string\",\"max\":10000},\"email\":{\"type\":\"string\",\"index\":\"not_analyzed\"}}}}}"

	t.Equal(expJson, j.String())
}
