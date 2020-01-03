package asset

import (
	"testing"

	"github.com/gregoryv/asserter"
)

func TestSrcWriter(t *testing.T) {
	a := NewSrcWriter()
	a.Create("generated_assets.go", "package asset")
	defer a.Close()
	a.WriteConst("x", "go.mod")

	err := a.WriteConst("bad", "no_such_file")
	assert := asserter.New(t)
	assert(err != nil).Error(err)
}
