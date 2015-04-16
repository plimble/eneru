package eneru

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

const (
	comma    = 0x2c
	colon    = 0x3a
	lbrace   = 0x7b
	rbrace   = 0x7d
	quote    = 0x22
	lbracket = 0x5b
	rbracket = 0x5d
	minus    = 0x2d
)

type ObjectFunc func(j *Json)

type Json struct {
	buf  *bytes.Buffer
	more bool
}

func (j *Json) getNumberByte(a int) byte {
	switch a {
	case 0:
		return 0x30
	case 1:
		return 0x31
	case 2:
		return 0x32
	case 3:
		return 0x33
	case 4:
		return 0x34
	case 5:
		return 0x35
	case 6:
		return 0x36
	case 7:
		return 0x37
	case 8:
		return 0x38
	case 9:
		return 0x39
	}

	return 0x00
}

func (j *Json) getInt(val int) {
	if val < 0 {
		j.buf.WriteByte(minus)
		val *= -1
	}

	digit := 0
	a := val
	for {
		a = a / 10
		if a == 0 {
			break
		}
		digit++
	}

	for i := digit; i >= 0; i-- {
		pow := int(math.Pow10(i))
		j.buf.WriteByte(j.getNumberByte(val / pow))
		val = val % pow
	}
}

// func (j *Json) GetFloat(val float64) {
// 	fmt.Println("val", val)
// 	l := int(val)
// 	fmt.Println("left", l)
// 	r := val - 0.2

// 	fmt.Println("right", r)
// }

func NewJson(fn ObjectFunc) *bytes.Buffer {
	j := &Json{
		buf:  bufPool.Get(),
		more: false,
	}

	j.buf.WriteByte(lbrace)
	fn(j)
	j.buf.WriteByte(rbrace)

	return j.buf
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

func (j *Json) AI(key string, vals ...int) {
	if j.more {
		j.buf.WriteByte(comma)
	}

	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteByte(lbracket)
	for i := 0; i < len(vals); i++ {
		j.getInt(vals[i])
		if i != len(vals)-1 {
			j.buf.WriteByte(comma)
		}
	}
	j.buf.WriteByte(rbracket)
	j.more = true
}

func (j *Json) B(key string, val bool) {
	if j.more {
		j.buf.WriteByte(comma)
	}

	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	if val {
		j.buf.WriteString("true")
	} else {
		j.buf.WriteString("false")
	}
	j.more = true

}

func (j *Json) AS(key string, vals ...string) {
	if j.more {
		j.buf.WriteByte(comma)
	}

	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteByte(lbracket)
	for i := 0; i < len(vals); i++ {
		j.buf.WriteByte(quote)
		j.buf.WriteString(vals[i])
		j.buf.WriteByte(quote)
		if i != len(vals)-1 {
			j.buf.WriteByte(comma)
		}
	}
	j.buf.WriteByte(rbracket)
	j.more = true
}

func (j *Json) AF(key string, prec int, vals ...float64) {
	if j.more {
		j.buf.WriteByte(comma)
	}

	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteByte(lbracket)
	for i := 0; i < len(vals); i++ {
		j.buf.WriteString(strconv.FormatFloat(vals[i], 'f', prec, 32))
		if i != len(vals)-1 {
			j.buf.WriteByte(comma)
		}
	}
	j.buf.WriteByte(rbracket)
	j.more = true
}

func (j *Json) AO(key string, fns ...ObjectFunc) {
	if j.more {
		j.buf.WriteByte(comma)
	}
	j.more = false
	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteByte(lbracket)

	for i := 0; i < len(fns); i++ {
		if j.more {
			j.buf.WriteByte(comma)
		}
		j.more = false
		j.buf.WriteByte(lbrace)
		fns[i](j)
		j.buf.WriteByte(rbrace)
	}

	j.buf.WriteByte(rbracket)
}

func (j *Json) I(key string, val int) {
	if j.more {
		j.buf.WriteByte(comma)
	}

	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.getInt(val)
	j.more = true
}

func (j *Json) F(key string, val float64, prec int) {
	if j.more {
		j.buf.WriteByte(comma)
	}

	j.buf.WriteByte(quote)
	j.buf.WriteString(key)
	j.buf.WriteByte(quote)
	j.buf.WriteByte(colon)
	j.buf.WriteString(strconv.FormatFloat(val, 'f', prec, 32))
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

// import (
//     "github.com/plimble/utils/strings2"
//     "strconv"
// )

// const (
//     comma  = 0x2c
//     colon  = 0x3a
//     lbrace = 0x7b
//     rbrace = 0x7d
//     quote  = 0x22
// )

// type ObjectFunc func(j *Json)

// type Json struct {
//     cache [][]byte
//     buf   []byte
//     more  bool
// }

// func NewJson() *Json {
//     j := &Json{
//         buf:  []byte{},
//         more: false,
//     }

//     j.buf = append(j.buf, lbrace)

//     return j
// }

// func (j *Json) S(key, val string) {

//     j.commaEnd()
//     j.quote(key)
//     j.buf = append(j.buf, colon)
//     j.quote(val)
//     j.more = true
// }

// func (j *Json) I(key string, val int) {
//     j.commaEnd()
//     j.quote(key)
//     j.buf = append(j.buf, colon)
//     j.buf = append(j.buf, []byte(strconv.Itoa(val))...)
//     j.more = true
// }

// func (j *Json) O(key string, fn ObjectFunc) {
//     j.commaEnd()
//     j.more = false
//     j.quote(key)
//     j.buf = append(j.buf, colon, lbrace)

//     fn(j)

//     j.buf = append(j.buf, rbrace)
// }

// func (j *Json) commaEnd() {
//     if j.more {
//         j.buf = append(j.buf, comma)
//     }
// }

// func (j *Json) quote(s string) {
//     j.buf = append(j.buf, quote)
//     j.buf = append(j.buf, string2.StringToBytes(s)...)
//     j.buf = append(j.buf, quote)
// }

// func (j *Json) Bytes() []byte {
//     j.buf = append(j.buf, rbrace)
//     return j.buf
// }

// func (j *Json) String() string {
//     j.buf = append(j.buf, rbrace)
//     return string2.BytesToString(j.buf)
// }
