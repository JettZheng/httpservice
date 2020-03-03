package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw,"OOps",http.StatusBadRequest)
			//rw.WriteHeader(http.StatusBadRequest)
			//rw.Write([]byte("Ooops"))
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.ListenAndServe("localhost:9090", nil)
}
