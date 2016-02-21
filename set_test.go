// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package set_test

import (
	"math"
	"testing"
	"testing/quick"

	"github.com/soniakeys/set"
)

// show that float64 is not reflexive
type naiveFEle float64

func (x naiveFEle) Equal(e set.Element) bool {
	return x == e.(naiveFEle)
}

func TestN(t *testing.T) {
	n := naiveFEle(math.NaN())
	if set.Reflexive(n) {
		t.Fatal("reflexive NaN")
	}
}

// define a float64 element that is reflexive
type fEle float64

func (x fEle) Equal(e set.Element) bool {
	y, ok := e.(fEle)
	if !ok {
		return false
	}
	if math.IsNaN(float64(x)) && math.IsNaN(float64(y)) {
		return true
	}
	if x != y {
		return false
	}
	return math.Signbit(float64(x)) == math.Signbit(float64(y))
}

func TestQ(t *testing.T) {
	cf := &quick.Config{MaxCount: 1000}
	fr := func(x fEle) bool { return set.Reflexive(x) }
	fs := func(x, y fEle) bool { return set.Symmetric(x, y) }
	ft := func(x, y, z fEle) bool { return set.Transitive(x, y, z) }
	if err := quick.Check(fr, cf); err != nil {
		t.Fatal("reflexive fail:", err)
	}
	if err := quick.Check(fs, cf); err != nil {
		t.Fatal("symmetric fail:", err)
	}
	if err := quick.Check(ft, cf); err != nil {
		t.Fatal("transitive fail:", err)
	}
}

func TestC(t *testing.T) {
	// little tests for code coverage
	a := fEle(3)
	b := fEle(3)
	c := fEle(3)
	if !set.Transitive(a, b, c) {
		t.Fatal("transitive fail")
	}

	s := set.Set{a}
	if s.Equal(a) {
		t.Fatal("set equals element")
	}
	if s.Equal(set.Set{fEle(4)}) {
		t.Fatal("not equal")
	}
}
