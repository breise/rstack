package rstack

import (
	"fmt"
	"strings"
)

type RStack struct {
	v interface{}
	p *RStack
}

func (s *RStack) Push(v interface{}) *RStack {
	return &RStack{v, s}
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
		return nil, "", fmt.Errorf("%s", "Pop(): Canmot pop any empty RStack")
	}
	return s.p, s.v, nil
}

func (s *RStack) ToSlice() []interface{} {
	if s == nil {
		return []interface{}{}
	}
	return append(s.p.ToSlice(), s.v)
}

func (s *RStack) ToStringSlice() []string {
	if s == nil {
		return []string{}
	}
	v := fmt.Sprintf("%v", s.v)
	return append(s.p.ToStringSlice(), v)
}

func (s *RStack) Join(sep string) string {
	return strings.Join(s.ToStringSlice(), sep)
}
