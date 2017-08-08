package kodi

import "encoding/xml"

type Movie struct {
	XMLName       xml.Name `xml:"movie"`
	Title         string   `xml:"title"`
	OriginalTitle string   `xml:"originaltitle"`
	Rating        float32  `xml:"rating"`
	Votes         int      `xml:"votes"`
	Top250        int      `xml:"top250"`
	Year          int      `xml:"year"`
	Plot          string   `xml:"plot"`
	Outline       string   `xml:"outline"`
	Tagline       string   `xml:"tagline"`
	Runtime       int      `xml:"runtime"`
	MPAA          string   `xml:"mpaa"`
	PlayCount     int      `xml:"playcount"`
	ID            string   `xml:"id"`
	TmdbID        int64    `xml:"tmdbid"`
	Set           string   `xml:"set"`
	SortTitle     string   `xml:"sorttitle"`
	Trailer       string   `xml:"trailer"`
	Watched       bool     `xml:"watched"`
	Credits       string   `xml:"credits"`
	Director      string   `xml:"director"`
	Studio        []string `xml:"studio"`
	Genre         []string `xml:"genre"`
	Country       []string `xml:"country"`
}

type Thumb struct {
	XMLName xml.Name `xml:"thumb"`
}

type
