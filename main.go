package main

import (
	"flag"

	"github.com/vysokyj/kodiak/storage"
)

func main() {
	var t string
	var d string
	flag.StringVar(&t, "t", "movies", "Type: movies or series")
	flag.StringVar(&d, "d", "./", "Directory")
	flag.Parse()

	storage := storage.NewStorage(d)
	storage.List()
}
