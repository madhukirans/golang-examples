package main

import (
	//"io"
	//"bufio"
	//"compress/gzip"
	//"bytes"
	//"os"
	//	"fmt"
	//"os"
	"fmt"
	//	"bytes"
	//	"io"
	"compress/gzip"
	//"os"

	"bytes"
	"io"
)

func main() {

	str := `hello asdfas dfasd asdfasdf asdfas dfgs sdfasdfasdfasfasfasdf hjggggggggggggggggggg yutuytuytuyt tuytuyt78687tugjhbjb`

	var buf bytes.Buffer
	w := io.Writer(&buf)
	gw := gzip.NewWriter(w)
	gw.Write([]byte(str))
	gw.Close()
	fmt.Println(string(buf.Bytes()), len(buf.Bytes()))

	//var rbuf bytes.Buffer
	r := io.Reader(&buf)
	gr, _ := gzip.NewReader(r)
	b := make([]byte, len(str))
	n, err := gr.Read(b)
	fmt.Println(string(b), n, err)

	str1 := `helloಒತ್ತಕ್ಷರ`
	fmt.Println(len([]rune(str1)))

	x := 'ಒ'
	fmt.Println(x)

	//fw1, _ := os.Create("/tmp/aaa.gz")
	//gw1 := gzip.NewWriter(fw1)
	//gw1.Write([]byte(str))
	//gw1.Close()
	//fw1.Close()
	//
	//r, _ := os.Open("/tmp/aaa.gz")
	//gr1, _:= gzip.NewReader(r)
	//b := make ([]byte, 2)
	//count, err := gr1.Read(b)
	//fmt.Println(string(b), count, err)

}
