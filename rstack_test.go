package rstack

import (
	"fmt"
	"strings"
	"testing"
)

var cases = []struct {
	name                string
	initialStackSlice   []interface{}
	action              func(s *RStack) (*RStack, interface{}, error)
	expectedStackString string
	expectedValue       string
}{
	{
		name:              "Push one on empty",
		initialStackSlice: []interface{}{},
		action: func(s *RStack) (*RStack, interface{}, error) {
			return s.Push("one"), "", nil
		},
		expectedStackString: "one",
	},
	{
		name:              "Push one on one",
		initialStackSlice: []interface{}{"one"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			return s.Push("two"), "", nil
		},
		expectedStackString: "one, two",
	},
	{
		name:              "Push two on one",
		initialStackSlice: []interface{}{"one", "two"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			return s.Push("three"), "", nil
		},
		expectedStackString: "one, two, three",
	},
	{
		name:              "Push two on three",
		initialStackSlice: []interface{}{"one", "two", "three"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			return s.Push("four").Push("five"), "", nil
		},
		expectedStackString: "one, two, three, four, five",
	},
	{
		name:              "Pop two from five",
		initialStackSlice: []interface{}{"one", "two", "three", "four", "five"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			sFour, five, err := s.Pop()
			if err != nil {
				return nil, "", err
			}
			sThree, four, err := sFour.Pop()
			if err != nil {
				return nil, "", err
			}
			return sThree, strings.Join([]string{four.(string), five.(string)}, ", "), nil
		},
		expectedStackString: "one, two, three",
		expectedValue:       "four, five",
	},
	{
		name:              "Pop two from three",
		initialStackSlice: []interface{}{"one", "two", "three"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			sTwo, three, err := s.Pop()
			if err != nil {
				return nil, "", err
			}
			sOne, two, err := sTwo.Pop()
			if err != nil {
				return nil, "", err
			}
			return sOne, strings.Join([]string{two.(string), three.(string)}, ", "), nil
		},
		expectedStackString: "one",
		expectedValue:       "two, three",
	},
	{
		name:              "Pop one from two",
		initialStackSlice: []interface{}{"one", "two"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			sOne, two, err := s.Pop()
			if err != nil {
				return nil, "", err
			}
			return sOne, two.(string), nil
		},
		expectedStackString: "one",
		expectedValue:       "two",
	},
	{
		name:              "Pop one from one",
		initialStackSlice: []interface{}{"one"},
		action: func(s *RStack) (*RStack, interface{}, error) {
			sNil, one, err := s.Pop()
			if err != nil {
				return nil, "", err
			}
			return sNil, one.(string), nil
		},
		expectedStackString: "",
		expectedValue:       "one",
	},
}

func TestRStack(t *testing.T) {
	for i, tc := range cases {
		desc := fmt.Sprintf("Test Case %d: %s", i, tc.name)
		t.Run(desc, func(t *testing.T) {
			initialStack := NewFromSlice(tc.initialStackSlice)
			got, v, err := tc.action(initialStack)
			if err != nil {
				t.Errorf("%s: Cannot action(). Error: %s", desc, err)
			} else {
				gotString := got.Join(", ")
				if gotString != tc.expectedStackString {
					t.Errorf("%s:\n\texp: %s\n\tgot: %s", desc, tc.expectedStackString, gotString)
				}
			}
			if tc.expectedValue != "" {
				if tc.expectedValue != v {
					t.Errorf("%s:\n\texp Value: %s\n\tgot Value: %s", desc, tc.expectedValue, v)
				}
			}
		})
	}
}
