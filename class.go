package merge

import (
	"fmt"
	"slices"
	"strings"
)

func ClassMergeFunc(existing any, incoming any) (committed any, remaining any) {
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
