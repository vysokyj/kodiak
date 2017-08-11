package storage

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestUnmarshallMarshall(t *testing.T) {
	xmlFile, err := os.Open("../movie.nfo")
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
	output, err := xml.MarshalIndent(movie, "", "  ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
	os.Stdout.WriteString("\n")
}
