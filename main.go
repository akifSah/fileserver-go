package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "Port to serve on.")
	dir := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))

	log.Fatal(http.ListenAndServe(*port, fs))

}
