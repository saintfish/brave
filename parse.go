package brave

import (
	"fmt"
	"github.com/saintfish/trie.go"
	"reflect"
	"strings"
	"unicode/utf8"
)

type parserState struct {
	input  string
	offset int
}

func (ps parserState) done() bool {
	return ps.offset >= len(ps.input)
}

func (ps parserState) advanceBy(offset int) parserState {
	newState := ps
	newState.offset += offset
	return newState
}

func (ps parserState) advanceToEnd() parserState {
	newState := ps
	newState.offset = len(ps.input)
	return newState
}

type tokenizer interface {
	getToken(state parserState) (match bool, newState parserState, value interface{})
}

type tokenizerParser interface {
	getTokenizer(state parserState) (newState parserState, tokenizer tokenizer, value interface{})
}

type spaceSkipper struct {
	optional bool
}

func (ss *spaceSkipper) getTokenizer(state parserState) (newState parserState, tokenizer tokenizer, value interface{}) {
	if state.done() {
		return state, nil, nil
	}
	for i, c := range state.input[state.offset:] {
		if !strings.ContainsRune(" \t", c) {
			if i == 0 {
				return state, nil, nil
			}
			return state.advanceBy(i), ss, nil
		}
	}
	return state.advanceToEnd(), ss, nil
}

func (ss *spaceSkipper) getToken(state parserState) (bool, parserState, interface{}) {
	if state.done() {
		return false, state, nil
	}
	for i, c := range state.input[state.offset:] {
		if !strings.ContainsRune(" \t\u2029", c) {
			match := i != 0 || ss.optional
			return match, state.advanceBy(i), nil
		}
	}
	return true, state.advanceToEnd(), nil
}

type dict struct {
	trie        *trie.Trie
	patternTrie *trie.Trie
}

func (d *dict) add(key string, value interface{}, pattern bool) {
	if len(key) == 0 {
		return
	}
	if d.trie == nil {
		d.trie = trie.NewTrie()
	}
	if d.patternTrie == nil {
		d.patternTrie = trie.NewTrie()
	}
	d.trie.Add([]byte(key), value)
	if pattern {
		d.patternTrie.Add([]byte(key), value)
	}
}

func (d *dict) getTokenizer(state parserState) (newState parserState, tokenizer tokenizer, value interface{}) {
	if state.done() {
		return state, nil, nil
	}
	m, found := d.patternTrie.MatchLongestPrefixString(state.input[state.offset:])
	if !found || m.PrefixLength == 0 {
		return state, nil, nil
	}
	return state.advanceBy(m.PrefixLength), d, m.Value
}

func (d *dict) getToken(state parserState) (bool, parserState, interface{}) {
	if state.done() {
		return false, state, nil
	}
	m, found := d.trie.MatchLongestPrefixString(state.input[state.offset:])
	if !found {
		return false, state, nil
	}
	return m.PrefixLength > 0, state.advanceBy(m.PrefixLength), m.Value
}

type charChecker struct {
	char string
}

func (cc *charChecker) getToken(state parserState) (bool, parserState, interface{}) {
	if state.done() {
		return false, state, nil
	}
	if !strings.HasPrefix(state.input[state.offset:], cc.char) {
		return false, state, nil
	}
	return true, state.advanceBy(len(cc.char)), nil
}

type charCheckerParser struct {
}

func (*charCheckerParser) getTokenizer(state parserState) (newState parserState, tokenizer tokenizer, value interface{}) {
	if state.done() {
		return state, nil, nil
	}
	_, size := utf8.DecodeRuneInString(state.input[state.offset:])
	if size == 0 {
		return state, nil, nil
	}
	char := state.input[state.offset : state.offset+size]
	return state.advanceBy(size), &charChecker{char}, nil
}

type tokenizerValue struct {
	tokenizer tokenizer
	value     interface{}
}

func parsePattern(pattern string, tokenizerParser []tokenizerParser) []tokenizerValue {
	var ps = parserState{pattern, 0}
	var tvList []tokenizerValue
	for !ps.done() {
		maxPs := ps
		maxTv := tokenizerValue{}
		for _, tp := range tokenizerParser {
			candidatePs, tokenizer, value := tp.getTokenizer(ps)
			if tokenizer != nil && candidatePs.offset > maxPs.offset {
				maxPs = candidatePs
				maxTv = tokenizerValue{tokenizer, value}
			}
		}
		if maxPs.offset == ps.offset {
			return nil
		}
		tvList = append(tvList, maxTv)
		ps = maxPs
	}
	return tvList
}

type parser struct {
	tokenizerParser []tokenizerParser
	rules           []*rule
}

type rule struct {
	tvList      []tokenizerValue
	templateObj interface{}
}

func newParser(tokenizerParser []tokenizerParser) *parser {
	return &parser{tokenizerParser: tokenizerParser}
}

func (p *parser) addRule(pattern string, templateObj interface{}) error {
	tvList := parsePattern(pattern, p.tokenizerParser)
	if tvList == nil {
		return fmt.Errorf("Unable to parse pattern")
	}
	testMap := newMap(false)
	for _, tv := range tvList {
		if tv.value != nil {
			if testMap.has(tv.value) {
				return fmt.Errorf("Dup key %#v in pattern %s", tv.value, pattern)
			}
			testMap.set(tv.value, tv.value)
		}
	}
	testObj, err := mapData(templateObj, testMap)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(templateObj, testObj) {
		return fmt.Errorf("Object mapping failed")
	}
	p.rules = append(p.rules, &rule{tvList, templateObj})
	return nil
}

func (p *parser) parse(input string) (match bool, offset int, value interface{}) {
	state := parserState{input, 0}
	maxOffset := 0
	var maxValue interface{}
	for _, r := range p.rules {
		m, s, v := r.parse(state)
		if m && s.offset > maxOffset {
			maxOffset = s.offset
			maxValue = v
		}
	}
	if maxOffset == 0 {
		return false, 0, nil
	}
	return true, maxOffset, maxValue
}

func (p *rule) parse(state parserState) (match bool, newState parserState, value interface{}) {
	m := newMap(false)
	newState = state
	for _, tv := range p.tvList {
		match, tokenState, value := tv.tokenizer.getToken(newState)
		if !match {
			return false, state, nil
		}
		newState = tokenState
		if tv.value != nil {
			m.set(tv.value, value)
		}
	}
	obj, err := mapData(p.templateObj, m)
	if err != nil {
		return false, state, nil
	}
	return true, newState, obj
}
