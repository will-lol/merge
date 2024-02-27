package merge

type Merger interface {
	Merge(existing Attrs, incoming Attrs) Attrs
}

type merger struct {
	AttrsFuncMap map[string][]MergeFunc
}

type Attrs map[string]any

type MergeFunc func(existing any, incoming any) (committed any, remaining any)

func New(attrsFuncMap map[string][]MergeFunc) Merger {
	for _, funcs := range attrsFuncMap {
		funcs = append(funcs, DefaultMergeFunc)
	}

	return &merger{
		AttrsFuncMap: attrsFuncMap,
	}
}

func DefaultMergeFunc(existing any, incoming any) (remaining any, committed any) {
	return nil, incoming
}

func isFullyMerged(remaining any) bool {
	return remaining == nil || remaining == ""
}

func (m merger) Merge(existing Attrs, incoming Attrs) Attrs {
	for key, existingValue := range existing {
		incomingValue, ok := incoming[key]

		if ok {
			mergeFuncs, ok := m.AttrsFuncMap[key]

			var remaining, committed any
			remaining = existingValue
			committed = incomingValue

			if ok {
				for _, mergeFunc := range mergeFuncs {
					if !isFullyMerged(remaining) {
						remaining, committed = mergeFunc(remaining, committed)
					} else {
						break
					}
				}
				incoming[key] = committed
			} else {
				remaining, committed = DefaultMergeFunc(remaining, committed)
				incoming[key] = committed
			}
		} else {
			incoming[key] = existingValue
		}
	}

	return incoming
}
