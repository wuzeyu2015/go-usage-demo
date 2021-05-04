//go:generate statik -src=./public -include=*.jpg,*.txt,*.html,*.css,*.js

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "testStatik/statik"
	"github.com/rakyll/statik/fs"
)

// Before buildling, run go generate.
// Then, run the main program and visit http://localhost:8080/public/hello.txt
func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	// Access individual files by their paths.
	r, err := statikFS.Open("/crush.json")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	http.ListenAndServe(":8080", nil)
}
