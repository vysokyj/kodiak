package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

// Bundle is directory with movie/series and its metadata files
type Bundle struct {
	dir   string
	movie *Movie
}

// NewBundle creates new bundle
func NewBundle(dir string) (*Bundle, error) {
	b := new(Bundle)
	b.dir = dir
	err := b.Scan()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Scan check and load bundle content files
func (b *Bundle) Scan() error {
	files, _ := ioutil.ReadDir(b.dir)
	for _, f := range files {
		if f.Name() == "movie.nfo" {
			p := path.Join(b.dir, "movie.nfo")
			b.movie = NewMovie(p)
		}
	}
	if b.movie == nil {
		return errors.New("Missing movie.nfo")
	}
	return nil
}

func (b *Bundle) String() string {
	return fmt.Sprintf("%s (%d)", b.movie.Title, b.movie.Year)
	//return b.dir

}
