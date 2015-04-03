package eneru

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilsSuite struct {
	suite.Suite
}

func TestUtilsSuite(t *testing.T) {
	suite.Run(t, &UtilsSuite{})
}

func (t *UtilsSuite) TestBuildPath() {
	path := buildPath("index", "type", "id", "_action")
	t.Equal("index/type/id/_action/", path)

	path = buildPath("index", "", "id", "_action")
	t.Equal("index/id/_action/", path)
}
