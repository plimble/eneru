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
						j.I("int", 10000)
						j.AI("ai", 10, 20, 30, 40)
						j.AS("as", "10", "20", "30", "40")
						j.AF("af", 2, 10.12, 20.321, 30.553, 40.22222222)
					})
					j.O("email", func(j *Json) {
						j.F("float", 10.123, 3)
						j.S("type", "string")
						j.S("index", "not_analyzed")
					})
				})
			})
		})
	})

	expJson := "{\"mappings\":{\"book\":{\"properties\":{\"name\":{\"type\":\"string\",\"int\":10000,\"ai\":[10,20,30,40],\"as\":[\"10\",\"20\",\"30\",\"40\"],\"af\":[10.12,20.32,30.55,40.22]},\"email\":{\"float\":10.123,\"type\":\"string\",\"index\":\"not_analyzed\"}}}}}"

	t.Equal(expJson, j.String())
}
