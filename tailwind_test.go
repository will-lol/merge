package merge

import (
	"fmt"
	"testing"
)

func TestTailwindMergeFuncSimple(t *testing.T) {
	merger, err := NewTailwindMerge()
	if err != nil {
		t.Fatal(err)
	}

	m := New(map[string][]MergeFunc{
		"class": {merger.TailwindMergeFunc},
	})

	attrs := m.Merge(Attrs{"class": "px-2 py-1 bg-red hover:bg-dark-red"}, Attrs{"class": "p-3 bg-[#B91C1C]"})
	fmt.Println(attrs)

	const expected = "hover:bg-dark-red p-3 bg-[#B91C1C]"
	if attrs["class"] != expected {
		t.Fatalf("Expected %q but got %q", expected, attrs["class"])
	}
}
