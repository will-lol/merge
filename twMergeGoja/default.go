package twMergeGoja

import (
	_ "embed"
	"errors"

	"github.com/dop251/goja"
	"github.com/will-lol/merge"
)

type twMerge struct {
	Runtime *goja.Runtime
	Func *goja.Callable
}

//go:embed lib/bundle.js
var twMergeJs string

func NewTwMerge() (merge.TwMerge, error) {
	vm := goja.New()

	_, err := vm.RunScript("twMerge.js", twMergeJs)
	if err != nil {
		return nil, err
	}

	obj, err := vm.RunString("m.twMerge")
	if err != nil {
		return nil, errors.New("twMerge not found")
	}
	f, ok := goja.AssertFunction(obj)
	if !ok {
		return nil, errors.New("twMerge not a function")
	}

	merge := &twMerge{
		Runtime: vm,
		Func: &f,
	}

	// tailwind-merge takes a long time in the first call, but is much faster in subsequent calls. We get that first call over with.
	merge.Merge("", "")

	return merge, nil
}

func (m twMerge) Merge(existing string, incoming string) (*string, error) {
	f := *m.Func

	res, err := f(goja.Undefined(), m.Runtime.ToValue(incoming), m.Runtime.ToValue(existing))
	if err != nil {
		return nil, err
	}

	val := res.String()
	return &val, nil
}

func (m twMerge) Close() {
	m.Runtime.Interrupt("halt")
}
