package storage

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Bundle is directory with movie/series and its metadata files
type Bundle struct {
	dir        string
	parDir     string
	movie      *Movie
	containers []string
}

// NewBundle creates new bundle
func NewBundle(dir string) (*Bundle, error) {
	b := new(Bundle)
	b.parDir = path.Dir(dir)
	b.dir = path.Base(dir)
	b.containers = make([]string, 0, 1)
	err := b.Scan()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Bundle) renameContainers(newBase string) {
	for i := 0; i < len(b.containers); i++ {
		con := b.containers[i]
		ext := strings.ToLower(path.Ext(con))
		oldPath := path.Join(b.parDir, b.dir, con)
		newPath := path.Join(b.parDir, b.dir, newBase+ext)
		if oldPath == newPath {
			return
		}
		fmt.Printf("Rename '%s' -> '%s'\n", oldPath, newPath)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to rename %s to %s\n", oldPath, newPath)
		} else {
			b.containers[i] = path.Base(newPath)
		}
	}
}

func (b *Bundle) renameDirectory(newBase string) {
	oldPath := path.Join(b.parDir, b.dir)
	newPath := path.Join(b.parDir, newBase)
	if oldPath == newPath {
		return
	}
	fmt.Printf("Rename '%s' -> '%s'\n", oldPath, newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to rename %s to %s\n", oldPath, newPath)
	} else {
		b.dir = path.Base(newPath)
	}
}

// Scan check and load bundle content files
func (b *Bundle) Scan() error {
	files, _ := ioutil.ReadDir(b.dir)
	for _, f := range files {
		name := f.Name()
		ext := strings.ToLower(path.Ext(name))
		if ext == ".nfo" {
			b.movie = NewMovie(path.Join(b.dir, name))
		}
		if ext == ".avi" || ext == ".mkv" || ext == ".mp4" {
			b.containers = append(b.containers, name)
		}
	}
	if b.movie == nil {
		return errors.New("Missing movie.nfo")
	}
	return nil
}

// TitleAndYear returns movie name and year
func (b *Bundle) TitleAndYear() string {
	return fmt.Sprintf("%s (%d)", b.movie.Title, b.movie.Year)
}

// ID returns movie database ID
func (b *Bundle) ID() string {
	return fmt.Sprintf("%s", b.movie.ID)
}

func (b *Bundle) String() string {
	return b.TitleAndYear()
}

// Movie contains movie.nfo
func (b *Bundle) Movie() *Movie {
	return b.movie
}

// Expand names
func (b *Bundle) Expand() {
	b.renameContainers(b.movie.Title)
	b.renameDirectory(b.TitleAndYear())
}

// Reduce names
func (b *Bundle) Reduce() {
	b.renameContainers("movie")
	b.renameDirectory(b.ID())
}
