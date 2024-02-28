package merge

import (
	"fmt"
	"slices"
	"strings"
)

// The ClassMergeFunc is designed to be used with CSS classes. All of the classes in the incoming string that are not duplicates of those in the existing string will be merged together.
func ClassMergeFunc(existing any, incoming any) (remaining any, committed any) {
	existingString := fmt.Sprint(existing)
	incomingString := fmt.Sprint(incoming)
	existingClasses := strings.Split(existingString, " ")
	incomingClasses := strings.Split(incomingString, " ")

	for _, existingClass := range existingClasses {
		if !slices.Contains(incomingClasses, existingClass) {
			incomingClasses = append(incomingClasses, existingClass)
		}
	}

	return nil, strings.Join(incomingClasses, " ")
}
