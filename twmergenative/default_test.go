package twmergenative

import (
	"os"
	"testing"
)

func TestNewSimple(t *testing.T) {
	dat, err := os.ReadFile("./test.css")
	if err != nil {
		t.Fatal(err.Error())
	}
	_ = NewTwMerge(string(dat))
}

