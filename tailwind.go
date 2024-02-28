package merge

import (
	"fmt"
	"github.com/will-lol/merge/twMergeGoja"
)

type tailwindMerge struct {
	Merger twMergeGoja.TwMerge
}

type TailwindMerge interface {
	TailwindMergeFunc(existing any, incoming any) (remaining any, committed any)
}

// NewTailwindMerge creates an instance of twMergeGoja, the provider of the tailwind-merge library in JavaScript. It runs in [github.com/dop251/goja], making it far slower than the other included Merge functions. It is recommended that you run the NewTailwindMerge function in a seperate goroutine as it can have a runtime of up to 14ms. Subsequent calls of TailwindMergeFunc are much faster, and do not need to run in a goroutine.
func NewTailwindMerge() (TailwindMerge, error) {
	merger, err := twMergeGoja.NewTwMerge()
	if err != nil {
		return nil, err
	}
	return &tailwindMerge{
		Merger: merger,
	}, nil
}

// TailwindCSS is a CSS framework. Its classes can be intelligently merged by this MergeFunc.
// This MergeFunc uses https://github.com/dcastil/tailwind-merge and as such follows its rules.
// It will never leave classes unmerged returned as 'remaining'.
func (t tailwindMerge) TailwindMergeFunc(existing any, incoming any) (remaining any, committed any) {
	existingString := fmt.Sprint(existing)
	incomingString := fmt.Sprint(incoming)

	res, err := t.Merger.Merge(existingString, incomingString)
	if err != nil {
		panic(err)
	}

	return nil, *res
}
