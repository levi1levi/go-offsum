package offsum

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestReadPage(t *testing.T) {
	filename := "page.zlib"
	zlibfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader, err := zlib.NewReader(zlibfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	newfilename := strings.TrimSuffix(filename, ".zlib")

	writer, err := os.Create(newfilename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer writer.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Decompressed to ", newfilename)

}

func TestZlib(t *testing.T) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()

	fmt.Println(string(b.Bytes()))

	r, _ := zlib.NewReader(&b)
	io.Copy(os.Stdout, r)
	r.Close()

}
