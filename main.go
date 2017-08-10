package main

import (
	"flag"

	"github.com/vysokyj/kodiak/kodi"
)

func main() {
	var t string
	var d string
	flag.StringVar(&t, "t", "movies", "Type: movies or series")
	flag.StringVar(&d, "d", "./", "Directory")
	flag.Parse()

	if t == "movies" {
		stor := kodi.NewMovieStorage(d)
		stor.Scan()
	}
}
