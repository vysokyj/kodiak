package storage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Storage is directory with bundles
type Storage struct {
	dir     string
	bundles []*Bundle
}

// NewStorage creates new movies database base on some directory root
func NewStorage(dir string) *Storage {
	s := new(Storage)
	s.dir = dir
	s.bundles = make([]*Bundle, 0)
	s.Scan()
	return s
}

// Scan movies data
func (s *Storage) Scan() {
	s.bundles = s.bundles[:0]
	files, _ := ioutil.ReadDir(s.dir)
	for _, fi := range files {
		if fi.IsDir() {
			dir := path.Join(s.dir, fi.Name())
			bundle, err := NewBundle(dir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Skipping directory %s - %s\n", dir, err)
			} else {
				s.bundles = append(s.bundles, bundle)
			}

		}
	}
}

// List writes out storage content
func (s *Storage) List() {
	for _, b := range s.bundles {
		fmt.Println(b)
	}
}
