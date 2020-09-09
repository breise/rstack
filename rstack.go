package rstack

import (
	"fmt"
	"strings"
)

type RStack struct {
	v interface{}
	p *RStack
	sz int
}

func (s *RStack) Push(v interface{}) *RStack {
	sz := 0
	if s != nil {
		sz = s.sz
	}
	return &RStack{v, s, sz + 1}
}

func New() *RStack {
	return nil
}

func NewFromSlice(list []interface{}) *RStack {
	rv := New()
	for _, v := range list {
		rv = rv.Push(v)
	}
	return rv
}

func (s *RStack) Pop() (*RStack, interface{}, error) {
	if s == nil {
		return nil, nil, fmt.Errorf("%s", "Pop(): Cannot pop any empty RStack")
	}
	return s.p, s.v, nil
}

func (s *RStack) Length() int {
	if s == nil {
		return 0
	} else {
		return s.sz
	}
}

func (s *RStack) ToSlice() []interface{} {
	rv := make([]interface{}, s.Length())
	for p := s; p != nil; p = p.p {
		rv[p.sz-1] = p.v
	}
	return rv
}

func (s *RStack) ToStringSlice() []string {
	rv := make([]string, s.Length())
	for p := s; p != nil; p = p.p {
		rv[p.sz-1] = fmt.Sprintf("%v", p.v)
	}
	return rv
}

func (s *RStack) Join(sep string) string {
	return strings.Join(s.ToStringSlice(), sep)
}
