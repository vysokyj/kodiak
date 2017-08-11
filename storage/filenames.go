package storage

import "strings"

func SecureFileName(s string) string {
	s = strings.Replace(s, "<", "", -1)
	s = strings.Replace(s, ">", "", -1)
	s = strings.Replace(s, ":", " -", -1)
	s = strings.Replace(s, "/", "", -1)
	s = strings.Replace(s, "\\", "", -1)
	s = strings.Replace(s, "|", "", -1)
	s = strings.Replace(s, "?", "", -1)
	s = strings.Replace(s, "*", "", -1)
	return s
}
