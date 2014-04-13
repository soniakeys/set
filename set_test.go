// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package set_test

import (
	"testing"

	"github.com/soniakeys/set"
)

func Test(t *testing.T) {
	// little tests for code coverage
	var s set.Set
	if s.Equal(nil) {
		t.Fatal("nil") // nil set != nil interface
	}
	s.AddElement(intEle(3))
	var s2 set.Set
	s2.AddElement(intEle(4))
	if s.Equal(s2) {
		t.Fatal("neq") // {3) != {4}
	}
	if !set.Reflexive(s) {
		t.Fatal("reflexive")
	}
	if !set.Symetric(s, s2) {
		t.Fatal("symetric")
	}
	s3 := set.Set{intEle(3)}
	if !set.Transitive(s, s2, s3) {
		t.Fatal("3=4?")
	}
	s2.RemoveElement(intEle(4))
	s2.AddElement(intEle(3))
	if !set.Transitive(s, s2, s3) {
		t.Fatal("3!=3?")
	}
}
