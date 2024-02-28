package merge

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestClassMergeFuncSimple(t *testing.T) {
	m := New(map[string][]MergeFunc{
		"class": {ClassMergeFunc},
	})

	existingClasses := []string{"px-3", "mx-2", "bg-red-500"}
	incomingClasses := []string{"font-bold"}

	res := m.Merge(map[string]any{"class": strings.Join(existingClasses, " ")}, map[string]any{"class": strings.Join(incomingClasses, " ")})
	fmt.Println(res)

	for _, val := range []string{"px-3", "mx-2", "bg-red-500", "font-bold"} {
		classesStr := fmt.Sprint(res["class"])
		arr := strings.Split(classesStr, " ")

		if !slices.Contains(arr, val) {
			t.Fatalf("Expected %v to include class=%q", res, val)
		}
	}
}

func TestClassMergeFuncCollision(t *testing.T) {
	m := New(map[string][]MergeFunc{
		"class": {ClassMergeFunc},
	})

	existingClasses := []string{"px-4", "mx-2", "bg-red-500"}
	incomingClasses := []string{"font-bold", "px-4"}

	res := m.Merge(map[string]any{"class": strings.Join(existingClasses, " ")}, map[string]any{"class": strings.Join(incomingClasses, " ")})
	fmt.Println(res)

	classesStr := fmt.Sprint(res["class"])
	arr := strings.Split(classesStr, " ")
	expectedArr := []string{"px-4", "mx-2", "bg-red-500", "font-bold"}

	for _, val := range expectedArr {
		if !slices.Contains(arr, val) {
			t.Fatalf("Expected %v to include class=%q", res, val)
		}
	}

	if len(arr) != len(expectedArr) {
		t.Fatal("Too many classes")
	}
}
