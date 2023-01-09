package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", ":8080", "Port to serve on.")
	dir := flag.String("d", "./tmp", "the directory of static file to host")
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))

	fmt.Printf("Serving '%v' at %v", *dir, *port)
	log.Fatal(http.ListenAndServe(*port, fs))

}
