package twmergegoja

import (
	"fmt"
	"testing"
	"time"
)

func TestMergeSimple(t *testing.T) {
	m, err := NewTwMerge()
	if err != nil {
		t.Fatal(err)
	}
	res, err := m.Merge("p-3 bg-[#B91C1C]", "px-2 py-1 bg-red hover:bg-dark-red")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*res)
	const expected = "hover:bg-dark-red p-3 bg-[#B91C1C]"
	if *res != expected {
		t.Fatalf("Expected %q but got %q", expected, *res)
	}
}

func TestMergePerf(t *testing.T) {
	m, err := NewTwMerge()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		start := time.Now()
		_, err := m.Merge("p-3 bg-[#B91C1C]", "px-2 py-1 bg-red hover:bg-dark-red")
		if err != nil {
			t.Fatal(err)
		}
		elapsed := time.Since(start)
		fmt.Printf("Merge took %s\n", elapsed)
	}
}
