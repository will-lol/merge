package twmergenative

import (
	"os"
	"testing"
	"fmt"

	"github.com/tdewolff/parse/v2"
)

func TestNewSimple(t *testing.T) {
	dat, err := os.ReadFile("./test.css")
	if err != nil {
		t.Fatal(err.Error())
	}
	_ = NewTwMerge(string(dat))
}

func TestUnescape(t *testing.T) {
	res, err := unescapeIdent(parse.NewInputString("prose-pre\\:-translat\\00005Ce-x-1\\/2"))
	fmt.Println(unescapeIdent(parse.NewInputString("prose-img\\:max-h-\\[70vh\\]")))
	if err != nil {
		t.Fatal(err.Error())
	}
	const expected = "prose-pre:-translat\\e-x-1/2"
	if res != expected {
		t.Fatalf("Expected %q but got %q", expected, res)
	}
}

