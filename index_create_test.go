package eneru

import (
	"fmt"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	client := NewClient("http://172.17.8.101:9200")
	client.Debug(true)
	jack := NewCreateIndex(client, "jack")
	jack.Pretty(true)

	j := NewJson()
	j.O("mappings", func(j *Json) {
		j.O("book", func(j *Json) {
			j.O("properties", func(j *Json) {
				j.O("name", func(j *Json) {
					j.S("type", "string")
				})
				j.O("email", func(j *Json) {
					j.S("type", "string")
					j.S("index", "not_analyzed")
				})
			})
		})
	})

	jack.Body(j.Buffer())
	resp, err := jack.Do()
	fmt.Println(resp, err)
}
