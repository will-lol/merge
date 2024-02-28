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

func NewTailwindMerge() (TailwindMerge, error) {
	merger, err := twMergeGoja.NewTwMerge()
	if err != nil {
		return nil, err
	}
	return &tailwindMerge{
		Merger: merger,
	}, nil
}

func (t tailwindMerge) TailwindMergeFunc(existing any, incoming any) (remaining any, committed any) {
	existingString := fmt.Sprint(existing)
	incomingString := fmt.Sprint(incoming)

	res, err := t.Merger.Merge(existingString, incomingString)
	if err != nil {
		panic(err)
	}

	return nil, *res
}
