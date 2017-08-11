package storage

// Bundle is directory with movie/series and its metadata files
type Bundle struct {
	dir        string
	movieFiles []string
}

// NewBundle creates new bundle
func NewBundle(dir string) *Bundle {
	b := new(Bundle)
	b.dir = dir
	b.movieFiles = make([]string, 0)
	b.Scan()
	return b
}

// Scan check and load bundle content files
func (b *Bundle) Scan() {
	//files, _ := ioutil.ReadDir(m.dir)

}
