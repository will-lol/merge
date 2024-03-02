package twmergenative

import (

)

type TwMerge interface {
}

type twMerge struct {
}

func NewTwMerge(rawCss string) TwMerge {

	return nil
}


func (m twMerge) Merge(classes string) {

}

type cssDeclaration struct {
	Property string
	Value string
}

type cssProperty string

type shorthandExpander func(property cssProperty) []cssProperty 

type isCssType func(value string) bool

func isTime(value string) bool {

}

type cssValue struct {
	TypeFunc isCssType
	Representing []cssProperty
}

func ordered(values []cssValue) shorthandExpander {

}
