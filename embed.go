package asset

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gregoryv/nexus"
)

type Embeded map[string]string

func (em Embeded) Open(filename string) (io.ReadCloser, error) {
	body, found := em[filename]
	if !found {
		return nil, fmt.Errorf("file not found: %s", filename)
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(body))
	return ioutil.NopCloser(dec), nil
}

func NewSrcWriter() *SrcWriter {
	return &SrcWriter{
		Files:   make([]string, 0),
		MapName: "Asset",
	}
}

type SrcWriter struct {
	Files   []string
	Package string
	Strip   string
	MapName string
}

func (sw *SrcWriter) WriteTo(w io.Writer) error {
	p, perr := nexus.NewPrinter(w)
	p.Printf("package %s\n\n", sw.Package)
	p.Println("import \"strings\"")
	p.Println()
	p.Println("func init() {")

	enc := base64.NewEncoder(base64.StdEncoding, w)
	for _, filename := range sw.Files {
		in, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer in.Close()

		p.Printf(
			"\t%s[%q] = strings.Join([]string{\n",
			sw.MapName, filename[len(sw.Strip):],
		)
		line := make([]byte, 35)
		for {
			n, _ := in.Read(line)
			// Assume we can read the entire file
			p.Print("\t\t\"")
			enc.Write(line[:n])
			p.Print("\",\n")
			if n < len(line) {
				enc.Close()
				p.Print("\t")
				p.Println(`}, "")`)
				break
			}
		}
		p.Print("\n")
	}
	p.Println("}")
	return *perr
}
