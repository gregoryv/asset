package asset

import "testing"

func TestSrcWriter(t *testing.T) {
	a := NewSrcWriter()
	a.Create("generated_assets.go", "package asset")
	a.WriteConst("x", "go.mod")
	a.Close()
}
