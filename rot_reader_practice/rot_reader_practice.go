package main

import (
	//"fmt"
	"io"
	"os"
	"strings"
)
type rot13Reader struct {
	r io.Reader
}



func (rot_reader rot13Reader) Read(b []byte) (n int, err error){
	n, err = rot_reader.r.Read(b)
	for i := 0; i < len(b); i++ {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] = b[i] + 13
			if (b[i] > 90) {
				b[i] = b[i] - 91 + 65
			}
		}
		if (b[i] >= 97 && b[i] <= 122) {
			b[i] = b[i] + 13
			if (b[i] > 122) {
				b[i] = b[i] - 123 + 97
			}
		}		
	}
	return n, err
}


func main() {

	//a := []int{1, 2, 3}
	//test_modify_arr(a)
	//fmt.Println(a)
	// raw_reader := strings.NewReader("Lbh penpxrq gur pbqr!")
	raw_reader := strings.NewReader("HELLO")
	rot_reader := rot13Reader{raw_reader}

	/*
	b := make([]byte, 5)
	for {
		_, err := rot_reader.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("%v ", b)
	}
	*/
	
	io.Copy(os.Stdout, &rot_reader)
}