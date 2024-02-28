package merge

import (
	"fmt"
	"testing"
)

func TestMergeSimple(t *testing.T) {
	m := New(map[string][]MergeFunc{})
	res := m.Merge(map[string]any{"class": "px-5"}, map[string]any{"id": "button"})
	fmt.Println(res)
	if res["id"] != "button" && res["class"] != "px-5" {
		t.Fatalf("Needed id=button and class=px-5 but got %v", res)
	}
}

func TestMergeCollision(t *testing.T) {
	m := New(map[string][]MergeFunc{})
	res := m.Merge(map[string]any{"class": "px-5"}, map[string]any{"class": "mt-5"})
	fmt.Println(res)
	if res["class"] != "mt-5" {
		t.Fatalf("Needed class=mt-5 but got %v", res)
	}
}
