package storage

import (
	"fmt"
	"testing"
)

func TestFilenames(t *testing.T) {
	input := "Příšerný film: Zůčtování?"
	expected := "Příšerný film - Zůčtování"
	output := SecureFileName(input)
	if output != expected {
		fmt.Println(output)
		t.Fail()
	}
}
