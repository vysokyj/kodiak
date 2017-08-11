package storage

import (
	"io/ioutil"
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
	files, _ := ioutil.ReadDir(s.dir)
	for _, fi := range files {
		if fi.IsDir() {
			dir := path.Join(s.dir, fi.Name())
			bundle := NewBundle(dir)
			s.bundles = append(s.bundles, bundle)
		}
	}
}
