package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JettZheng/httpservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/",hh)

	http.HandleFunc( )
	http.ListenAndServe("localhost:9090", sm)
}
