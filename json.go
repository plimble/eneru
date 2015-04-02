package eneru

import (
	"bytes"
	"strconv"
)

const (
	comma  = 0x2c
	colon  = 0x3a
	lbrace = 0x7b
	rbrace = 0x7d
	quote  = 0x22
)

type ObjectFunc func(j *Json)

type Json struct {
	buf  *bytes.Buffer
	more bool
}

func NewJson() *Json {
	j := &Json{
		buf:  &bytes.Buffer{},
		more: false,
	}

	j.buf.WriteByte(lbrace)

	return j
}

func (j *Json) S(key, val string) {
	if j.more {
		j.buf.WriteByte(comma)
	}
	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteByte(quote)
	j.buf.WriteString(val)
	j.buf.WriteByte(quote)
	j.more = true
}

func (j *Json) I(key string, val int) {
	if j.more {
		j.buf.WriteByte(comma)
	}
	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteString(strconv.Itoa(val))
	j.more = true
}

func (j *Json) O(key string, fn ObjectFunc) {
	if j.more {
		j.buf.WriteByte(comma)
	}
	j.more = false
	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteByte(lbrace)

	fn(j)

	j.buf.WriteByte(rbrace)
}

func (j *Json) Bytes() []byte {
	j.buf.WriteByte(rbrace)
	return j.buf.Bytes()
}

func (j *Json) String() string {
	j.buf.WriteByte(rbrace)
	return j.buf.String()
}

func (j *Json) Buffer() *bytes.Buffer {
	j.buf.WriteByte(rbrace)
	return j.buf
}

// package eneru

// import (
//   "github.com/plimble/utils/strings2"
//   "strconv"
// )

// const (
//   comma  = 0x2c
//   colon  = 0x3a
//   lbrace = 0x7b
//   rbrace = 0x7d
//   quote  = 0x22
// )

// type ObjectFunc func(j *Json)

// type Json struct {
//   buf  []byte
//   more bool
// }

// func NewJson() *Json {
//   j := &Json{
//     buf:  []byte{},
//     more: false,
//   }

//   j.buf = append(j.buf, lbrace)

//   return j
// }

// func (j *Json) S(key, val string) {
//   j.commaEnd()
//   j.quote(key)
//   j.buf = append(j.buf, colon)
//   j.quote(val)
//   j.more = true
// }

// func (j *Json) I(key string, val int) {
//   j.commaEnd()
//   j.quote(key)
//   j.buf = append(j.buf, colon)
//   j.buf = append(j.buf, []byte(strconv.Itoa(val))...)
//   j.more = true
// }

// func (j *Json) O(key string, fn ObjectFunc) {
//   j.commaEnd()
//   j.more = false
//   j.quote(key)
//   j.buf = append(j.buf, colon, lbrace)

//   fn(j)

//   j.buf = append(j.buf, rbrace)
// }

// func (j *Json) commaEnd() {
//   if j.more {
//     j.buf = append(j.buf, comma)
//   }
// }

// func (j *Json) quote(s string) {
//   j.buf = append(j.buf, quote)
//   j.buf = append(j.buf, string2.StringToBytes(s)...)
//   j.buf = append(j.buf, quote)
// }

// func (j *Json) Bytes() []byte {
//   j.buf = append(j.buf, rbrace)
//   return j.buf
// }

// func (j *Json) String() string {
//   j.buf = append(j.buf, rbrace)
//   return string2.BytesToString(j.buf)
// }
