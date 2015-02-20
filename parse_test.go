package brave

import (
	"reflect"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
	numDict := &dict{}
	for i := 0; i <= 9; i++ {
		// Even numbers can be used in pattern
		numDict.add(strconv.Itoa(i), i, i%2 == 0)
	}
	opDict := &dict{}
	opDict.add("+", '+', true)
	opDict.add("-", '-', true)
	opDict.add("*", '*', true)
	opDict.add("/", '/', true)
	tokenizerList := []tokenizerParser{
		numDict,
		opDict,
		&spaceSkipper{true},
	}
	badPatterns := []string{
		"",
		"2+3",
		"a",
		"2+2",
	}
	dummyTemplate := 1
	for i, pattern := range badPatterns {
		p := newParser(tokenizerList)
		err := p.addRule(pattern, dummyTemplate)
		if err == nil {
			t.Errorf("Pattern %d should return error", i)
		}
	}
	pattern := "0 + 2 - 4 * 6 / 8"
	templateObj := []interface{}{0, 2, 4, 6, 8, '+', '-', '*', '/'}
	p := newParser(tokenizerList)
	err := p.addRule(pattern, templateObj)
	if err != nil {
		t.Error(err)
	}
	err = p.addRule("0*2+4", []interface{}{0, 2, 4, '*'})
	if err != nil {
		t.Error(err)
	}
	input := "1* 2+3+ 4 - 5"
	m, o, v := p.parse(input)
	if !m || o != len(input) {
		t.Error(m, o)
	}
	if !reflect.DeepEqual(v, []interface{}{1, 2, 3, 4, 5, '*', '+', '+', '-'}) {
		t.Error(v)
	}
}
