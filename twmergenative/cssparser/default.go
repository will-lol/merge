package cssparser

import (
	"bytes"
	"strconv"
	"fmt"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type cssParser struct {
	Parser *css.Parser
	State  state
	Css    CssClasses
}

type state struct {
	Class      string
	MediaQuery string
	Selector   string
}

type CssClasses map[string]MediaQueries

type MediaQueries map[string]Selectors

type Selectors map[string]Rulesets

type Rulesets map[string]string

func (c cssParser) consumeBeginAtRule() {
	var mediaQueryBuf bytes.Buffer
	for _, val := range c.Parser.Values() {
		mediaQueryBuf.Write(val.Data)
	}
	c.State.MediaQuery = mediaQueryBuf.String()
}

func (c cssParser) consumeEndAtRule() {
	c.State.MediaQuery = ""
}

func (c cssParser) consumeBeginRuleset() error {
	var selectorBytes bytes.Buffer
	var identBytes bytes.Buffer

	// Ensure identifier is a class
	vals := c.Parser.Values()
	if vals[0].TokenType != css.DelimToken {
		return nil
	} else if string(vals[0].Data) != "." {
		return nil
	}

	// Only the first identifier is the target class
	classWritten := false
	for _, val := range vals {
		if val.TokenType != css.IdentToken {
			selectorBytes.Write(val.Data)
		} else if !classWritten {
			identBytes.Write(val.Data)
			classWritten = true
		}
	}

	var err error
	c.State.Class, err = unescapeIdent(parse.NewInputBytes(identBytes.Bytes()))
	if err != nil {
		return err
	}
	c.State.Selector = selectorBytes.String()	
	return nil
}

func (c cssParser) consumeDeclaration(data []byte) {
	fmt.Println(string)
}

func (c cssParser) Parse() (CssClasses, error) {
	cssClasses := make(CssClasses)

	var mediaQuery string
loop:
	for {
		gt, _, data := c.Parser.Next()
		switch gt {
		case css.ErrorGrammar:
			break loop
		case css.BeginAtRuleGrammar: 

		case css.EndAtRuleGrammar:
			mediaQuery = ""

		case css.DeclarationGrammar:
			fmt.Println(string(data))

		case css.BeginRulesetGrammar:

			class, err := unescapeIdent(parse.NewInputBytes(identBytes.Bytes()))
			if err != nil {
				panic(err)
			}
			selector := selectorBytes.String()
			mediaQueries, ok := c[class]
			if !ok {
				mediaQueries = make(MediaQueries)
				c[class] = mediaQueries
			}
			selectors, ok := c[class][mediaQuery]
			if !ok {
				selectors = make(Selectors)
				c[class][mediaQuery] = selectors
			}
			rulesets, ok := c[class][mediaQuery][selector]
			if !ok {
				rulesets = make(Rulesets)
				c[class][mediaQuery][selector] = rulesets
			}
		}
	}
}

// https://www.w3.org/TR/css-syntax-3/#consume-escaped-code-point
func unescapeIdent(i *parse.Input) (string, error) {
	res := make([]rune, 0, i.Len())

	for {
		b := i.Peek(0)
		if b == 0 {
			break
		} else if isValidEscape(i) {
			i.Move(1)
			char, err := consumeEscape(i)
			if err != nil {
				return string(res), err
			}
			res = append(res, char)
		} else {
			res = append(res, rune(b))
		}
		i.Move(1)
	}

	return string(res), nil
}

func isValidEscape(i *parse.Input) bool {
	return i.Peek(0) == '\\' && i.Peek(1) != '\n'
}

func consumeEscape(in *parse.Input) (rune, error) {
	firstChar := in.Peek(0)
	if firstChar == 0 {
		return 0, in.Err()
	} else if isHexDigit(firstChar) {
		codePoint := make([]byte, 0, 6)
		codePoint = append(codePoint, firstChar)

		for i := 0; i < 5; i++ {
			char := in.Peek(1)
			if isHexDigit(char) {
				codePoint = append(codePoint, char)
			} else if isWhitespace(char) {
				in.Move(1)
				break
			} else {
				break
			}
			in.Move(1)
		}

		val, err := strconv.ParseInt(string(codePoint), 16, 32)
		if err != nil {
			return 0, err
		}

		return rune(val), nil
	} else {
		return rune(firstChar), nil
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isHexDigit(b byte) bool {
	return isDigit(b) || (b >= 'A' && b <= 'F') || (b >= 'a' && b <= 'f')
}

func isWhitespace(b byte) bool {
	return b == '\t' || b == ' ' || b == '\n'
}
