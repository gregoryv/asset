package x

import (
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/asset"
)

// will be populated by the SrcWriter
var Asset = make(asset.Embeded)

func Test_generated(t *testing.T) {
	got, err := Asset.Open("index.html")
	assert := asserter.New(t)
	assert(err == nil).Error(err)
	// same as htmlData in asset package
	assert().Contains(got, `<!DOCTYPE html>`)
}
