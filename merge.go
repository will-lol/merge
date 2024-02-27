package merge

import ()

type Merger interface {
}

type merger struct {
}

type Attrs map[string]any

type MergeFunc func(existing Attrs, incoming Attrs) (comitted Attrs, remaining Attrs)

func New(attrsMap map[string]MergeFunc) Merger {
	return &merger{}
}
