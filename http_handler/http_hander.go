package main
import (
	"fmt"
	"log"
	"net/http"
)

type String string
type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(str))
}

func (stru *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	greetings := fmt.Sprintf("%v%v%v", stru.Greeting, stru.Punct, stru.Who)
	fmt.Fprint(w, greetings)
}

func main() {

	http.Handle("/string", String("I'm newbie"))
	http.Handle("/struct", &Struct{
		"hello",
		":",
		"gopher",
	})
	//handler := String("I'm newbie")
	
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}

	
	// struc_handler := Struct {
	// 	"hello",
	// 	":",
	// 	"gopher",
	// }

	// err = http.ListenAndServe("localhost:4000/struct", &struc_handler)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}