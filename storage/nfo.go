package storage

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Movie is root xml element
type Movie struct {
	XMLName       xml.Name    `xml:"movie"`
	Title         string      `xml:"title"`
	OriginalTitle string      `xml:"originaltitle"`
	Rating        float32     `xml:"rating"`
	Votes         int         `xml:"votes"`
	Top250        int         `xml:"top250"`
	Year          int         `xml:"year"`
	Plot          string      `xml:"plot"`
	Outline       string      `xml:"outline"`
	Tagline       string      `xml:"tagline"`
	Runtime       int         `xml:"runtime"`
	MPAA          string      `xml:"mpaa"`
	PlayCount     int         `xml:"playcount"`
	ID            string      `xml:"id"`
	TmdbID        int64       `xml:"tmdbid"`
	Set           string      `xml:"set"`
	SortTitle     string      `xml:"sorttitle"`
	Trailer       string      `xml:"trailer"`
	Watched       bool        `xml:"watched"`
	Credits       []string    `xml:"credits"`
	Directors     []string    `xml:"director"`
	Studios       []string    `xml:"studio"`
	Genres        []string    `xml:"genre"`
	Countries     []string    `xml:"country"`
	Thumbnails    []Thumbnail `xml:"thumb"`
	Fanart        Fanart      `xml:"fanart"`
	Actors        []*Actor    `xml:"actor"`
	FileInfo      FileInfo    `xml:"fileinfo"`
}

// Actor is nested element in movie
type Actor struct {
	XMLName   xml.Name  `xml:"actor"`
	Name      string    `xml:"name"`
	Role      string    `xml:"role"`
	Thumbnail Thumbnail `xml:"thumb"`
}

// Fanart is nested element in movies
type Fanart struct {
	XMLName    xml.Name    `xml:"fanart"`
	Thumbnails []Thumbnail `xml:"thumb"`
}

// Thumbnail is nested element in movie and others
type Thumbnail struct {
	XMLName xml.Name `xml:"thumb"`
	Preview string   `xml:"preview,attr,omitempty"`
	Path    string   `xml:",chardata"`
}

// FileInfo holds file info
type FileInfo struct {
	XMLName       xml.Name      `xml:"fileinfo"`
	StreamDetails StreamDetails `xml:"streamdetails"`
}

// StreamDetails holds stream details
type StreamDetails struct {
	XMLName xml.Name `xml:"streamdetails"`
	Video   Video    `xml:"video"`
	Audio   Audio    `xml:"audio"`
}

// Video holds vido data
type Video struct {
	XMLName xml.Name `xml:"video"`
	//Aspect            float32  `xml:"aspect"`
	Aspect            string `xml:"aspect"`
	Codec             string `xml:"codec"`
	DurationInSeconds int64  `xml:"durationinseconds"`
	Height            int    `xml:"height"`
	Width             int    `xml:"width"`
	Scantype          string `xml:"scantype"`
}

// Audio holds audio data
type Audio struct {
	XMLName  xml.Name `xml:"audio"`
	Channels int      `xml:"channels"`
	Codec    string   `xml:"codec"`
}

// NewMovie scand movie.nfo
func NewMovie(path string) *Movie {
	xmlFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	b, _ := ioutil.ReadAll(xmlFile)

	movie := &Movie{}
	err = xml.Unmarshal(b, &movie)
	if err != nil {
		panic(err)
	}
	return movie
}
