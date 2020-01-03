package asset

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/gregoryv/nexus"
)

func NewSrcWriter() *SrcWriter {
	return &SrcWriter{}
}

type SrcWriter struct {
	io.WriteCloser
}

func (cw *SrcWriter) Create(filename, body string) error {
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	cw.WriteCloser = w
	fmt.Fprint(w, body)
	fmt.Fprintln(w)
	fmt.Fprintln(w)
	return nil
}

func (w *SrcWriter) WriteConst(name, filename string) error {
	return w.gen("const "+name, filename)
}

func (w *SrcWriter) gen(name, filename string) error {
	enc := base64.NewEncoder(base64.StdEncoding, w)
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()
	p, perr := nexus.NewPrinter(w)
	p.Printf("%s = %s(\n", name, "string")
	line := make([]byte, 30)
	for {
		n, _ := in.Read(line)
		// Assume we can read the entire file
		p.Print("\t\"")
		enc.Write(line[:n])
		if n < len(line) {
			enc.Close()
			p.Print("\")\n")
			break
		}
		p.Print("\" +\n")

	}
	p.Print("\n")
	return *perr
}
