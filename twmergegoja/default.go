// Package twmergegoja manages the Goja runtime and the running of the tailwind-merge JavaScript bundle. The bundle is fetched and built using updateTwMerge.sh. It is embedded in this Go package using [embed].
package twmergegoja

import (
	_ "embed"
	"errors"

	"github.com/dop251/goja"
)

type TwMerge interface {
	Merge(existing string, incoming string) (*string, error)
	Close()
}

type twMerge struct {
	Runtime *goja.Runtime
	Func *goja.Callable
}

//go:embed lib/bundle.js
var twMergeJs string

// NewTwMerge is a long running function (upwards of 10ms). I suggest you run it in a seperate goroutine.
func NewTwMerge() (TwMerge, error) {
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

	res, err := f(goja.Undefined(), m.Runtime.ToValue(existing), m.Runtime.ToValue(incoming))
	if err != nil {
		return nil, err
	}

	val := res.String()
	return &val, nil
}

func (m twMerge) Close() {
	m.Runtime.Interrupt("halt")
}
