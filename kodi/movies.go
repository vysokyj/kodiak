package kodi

// Movies is movie database
type Movies struct {
	root string
}

// NewMovies creates new movies database base on some directory root
func NewMovies(root string) *Movies {
	m := new(Movies)
	m.root = root
	return m
}

// Scan movies data
func (m *Movies) Scan() {

}
