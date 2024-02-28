// Package merge defines the main Merger interface and its implementation. Import and use this package to instantiate a Merger. This package is used primarily to manage the merging of HTML attribute sets.
package merge

// The Merger type merges two attribute sets. These may be any kind of attributes but this package was made for merging HTML attributes
type Merger interface {
	Merge(existing map[string]any, incoming map[string]any) map[string]any
}

type merger struct {
	AttrsFuncMap map[string][]MergeFunc
}

// A MergeFunc is a function that defines how two conflicting attributes should be merged.
// MergeFuncs should prioritise the incoming values and override existing values.
// If a MergeFunc does not know how to merge part of an attribute, it may return 'remaining'.
// Merged attributes should be returned in 'committed' in a type that is able to be converted to a string by fmt.Sprint(). 'remaining' attributes may be handled by downstream MergeFuncs.
type MergeFunc func(existing any, incoming any) (remaining any, committed any)

// New instantiates a Merger for you. You must pass an 'attrsFuncMap' that defines the MergeFuncs, though this may be empty.
// A string key is used for the name of the attribute. A slice of MergeFuncs is used to define the merge behaviour for this attribute.
// When Merger.Merge is called, the MergeFuncs are iterated through, which the 'remaining' and 'committed' values of each being passed into the following MergeFunc. The MergeFunc with index 0 is called first.
// If a MergeFunc leaves some part of an attribute un-merged in the 'remaining' variable, it will be handled by MergeFuncs defined with a higher index in the array.
// The DefaultMergeFunc is added to the end of all slices of MergeFuncs that simply discards all remaining attributes.
func New(attrsFuncMap map[string][]MergeFunc) Merger {
	for _, funcs := range attrsFuncMap {
		funcs = append(funcs, DefaultMergeFunc)
	}

	return &merger{
		AttrsFuncMap: attrsFuncMap,
	}
}

// The DefaultMergeFunc simply discards existing attributes. It is applied to all attributes without a defined merge behaviour.
func DefaultMergeFunc(existing any, incoming any) (remaining any, committed any) {
	return nil, incoming
}

func isFullyMerged(remaining any) bool {
	return remaining == nil || remaining == ""
}

// The main Merge function merges two attribute sets based on the config in its Merger object. The 'incoming' attributes are always prioritised over the 'existing' attributes.
func (m merger) Merge(existing map[string]any, incoming map[string]any) map[string]any {
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
