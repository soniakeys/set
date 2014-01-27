// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

// Sets as lists of interfaces.
//
// Interface Element has a single method Equal.  This allows generality
// in element types and general defintions of equality. Sets are Element
// slices.  You lose O(1) access in comparison to maps for example, but
// you gain this greater generality.
//
// Usually your application will allow something more efficent.  There are
// packages out there for map-based sets and for bitsets, for example.
// But just to demonstrate the possibility, here is an example of sets as
// lists of interfaces.
package set

import (
	"fmt"
	"math/rand"
)

type Element interface {
	Equal(Element) bool
}

type Set []Element

func (p *Set) AddElement(e Element) {
	if !p.HasElement(e) {
		*p = append(*p, e)
	}
}

func (p *Set) RemoveElement(e Element) {
	for i, ex := range *p {
		if e.Equal(ex) {
			s := *p
			last := len(s) - 1
			s[i], s[last] = s[last], s[i]
			*p = s[:last]
			return
		}
	}
}

func (s Set) HasElement(e Element) bool {
	for _, ex := range s {
		if e.Equal(ex) {
			return true
		}
	}
	return false
}

// Equal satifies the Element interface.
//
// An object of type Set can be an element of a set.
func (s Set) Equal(e Element) bool {
	t, ok := e.(Set)
	if !ok {
		return false
	}
	if len(s) != len(t) {
		return false
	}
	for _, se := range s {
		if !t.HasElement(se) {
			return false
		}
	}
	return true
}

func (s Set) String() string {
	r := "{"
	for i, j := range rand.Perm(len(s)) {
		if i > 0 {
			r += " "
		}
		r += fmt.Sprint(s[j])
	}
	return r + "}"
}

func (s Set) PowerSet() Set {
	r := Set{Set{}}
	for _, es := range s {
		var u Set
		for _, er := range r {
			u = append(u, append(er.(Set), es))
		}
		r = append(r, u...)
	}
	return r
}
