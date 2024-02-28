package twmergenative

import (
	"fmt"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type TwMerge interface {
}

type twMerge struct {
	Css string
}

func NewTwMerge(rawCss string) TwMerge {
	p := css.NewParser(parse.NewInputString(rawCss), false)

	for {
		gt, _, data := p.Next()
		if gt == css.ErrorGrammar {
			break
		}
		if gt == css.BeginAtRuleGrammar {
			fmt.Println("---")
			fmt.Println("Grammar type: " + gt.String() + ", Data: " + string(data))
			for _, val := range p.Values() {
				fmt.Println("Token: " + val.String())
			}
			for {
				gt, _, data := p.Next()
				fmt.Println("---")
				fmt.Println("Grammar type: " + gt.String() + ", Data: " + string(data))
				for _, val := range p.Values() {
					fmt.Println("Token: " + val.String())
				}
				if gt == css.EndAtRuleGrammar {
					break
				}
			}
		}
	}

	return nil
}

func (m twMerge) Merge(classes string) {

}
