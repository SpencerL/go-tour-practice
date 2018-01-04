package main

import (
	"io"
	//"os"
	"strings"
	"fmt"
)
type rot13Reader struct {
	r io.Reader
}



func (rot_reader rot13Reader) Read(b []byte)  {
	//b := make([]byte, 8)

	for {
		n, e := rot_reader.r.Read(b)
		if e != nil {
			fmt.Printf("(%v, %v) %v\n", n, e, b)
		}
		if e == io.EOF {
			break;
		}
	}
	
}

func main() {
	s := strings.NewReader("abcd")
	rot_reader := rot13Reader{s}
	
	
	
}