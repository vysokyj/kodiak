package main

import (
	"flag"

	"github.com/vysokyj/kodiak/storage"
)

func main() {
	var t string
	var d string
	var l bool
	var e bool
	var r bool
	flag.StringVar(&t, "t", "movies", "Type: movies or series")
	flag.StringVar(&d, "d", "./", "Directory")
	flag.BoolVar(&l, "l", false, "List database")
	flag.BoolVar(&e, "e", false, "Expand database names")
	flag.BoolVar(&r, "r", false, "Reduce database names")
	flag.Parse()

	storage := storage.NewStorage(d)

	if e {
		storage.Expand()
	}

	if r {
		storage.Reduce()
	}

	if l {
		storage.List()
	}

}
