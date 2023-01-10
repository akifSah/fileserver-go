package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()

	port := flag.String("p", ":8080", "Port to serve on.")
	dir := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	fs := http.FileServer(http.Dir(filepath.FromSlash(filepath.Join(pwd, *dir))))

	fmt.Printf("Serving '%v' at %v", filepath.Join(pwd, *dir), *port)
	log.Fatal(http.ListenAndServe(*port, fs))

}
