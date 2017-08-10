package kodi

import (
	"io/ioutil"
	"path"
)

// MovieBundle holds one movie
type MovieBundle struct {
	dir        string
	movieFiles []string
}

// NewMovieBundle creates new bundle
func NewMovieBundle(dir string) *MovieBundle {
	m := new(MovieBundle)
	m.dir = dir
	m.movieFiles = make([]string, 0)
	return m
}

func (m *MovieBundle) Scan() {
	//files, _ := ioutil.ReadDir(m.dir)

}

// MovieStorage is movie database
type MovieStorage struct {
	dir     string
	bundles []*MovieBundle
}

// NewMovieStorage creates new movies database base on some directory root
func NewMovieStorage(dir string) *MovieStorage {
	m := new(MovieStorage)
	m.dir = dir
	m.bundles = make([]*MovieBundle, 0)
	return m
}

// Scan movies data
func (m *MovieStorage) Scan() {
	files, _ := ioutil.ReadDir(m.dir)
	for _, fi := range files {
		if fi.IsDir() {
			//fmt.Println(fi.Name())
			dir := path.Join(m.dir, fi.Name())
			bundle := NewMovieBundle(dir)
			bundle.Scan()
			m.bundles = append(m.bundles, bundle)
		}
	}
}
