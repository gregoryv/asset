package asset_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/asset"
	"github.com/gregoryv/workdir"
)

func Test_gen(t *testing.T) {
	wd, _ := workdir.TempDir()
	defer wd.RemoveAll()
	wd.WriteFile("index.html", htmlData)

	sw := asset.NewSrcWriter()
	sw.Files = []string{wd.Join("index.html")}
	sw.Package = "x"
	sw.Strip = string(wd) + "/"
	out := "internal/x/out.go"
	w, _ := os.Create(out)
	err := sw.WriteTo(w)
	w.Close()
	assert := asserter.New(t)
	assert(err == nil).Error(err)
	res, err := exec.Command("gofmt", "-d", out).Output()
	assert(len(res) == 0).Errorf("%q", string(res))
}

var htmlData = []byte(`
<!DOCTYPE html>

<html>
  <head>
    <meta charset="utf-8">
  </head>
  <body>

  </body>
</html>
`)
