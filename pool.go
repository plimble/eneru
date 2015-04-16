package eneru

import (
	"github.com/plimble/utils/pool"
)

var bufPool *pool.BufferPool = pool.NewBufferPool(10000)

func SetPoolSize(n int) {
	bufPool = pool.NewBufferPool(n)
}
