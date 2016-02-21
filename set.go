// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package set

import (
	"fmt"
	"math/rand"
)

// An Element can be an element of a Set.
//
// A type satisfying Element must implement just one method, Equal.
// Equal returns true if the argument is to be considered equal to the
// receiver in some sense.  For a valid implementation of Equal, the
// functions Reflexive, Symmetric, and Transitive must return true for
// all possible element values.
//
// Implementations will typically type assert the argument to a value
// of the receiver type and then perform some value comparison.
type Element interface {
	Equal(Element) bool
}

// Refexive validates the reflexive property of the Equal method for the given
// element value.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Reflexive(e Element) bool {
	return e.Equal(e)
}

// Symmetric validates the symmetric property of the Equal method for the given
// element value.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Symmetric(a, b Element) bool {
	return a.Equal(b) == b.Equal(a)
}

// Transitive validates the transitive property of the Equal method for the
// given element values.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Transitive(a, b, c Element) bool {
	if a.Equal(b) && b.Equal(c) {
		return a.Equal(c)
	}
	return true
}

// Set is a type implementing the mathematical concept of a set.
//
// This type also implements the Element interface, so Sets can be elements
// of Sets.
//
// Set implementation is an Element slice.  Sets are conceptually unordered.
// While slices are ordered, this order in a Set must be considered irrelevant.
// Also, while slices allow duplicate values, a Set must be maintained so that
// Equal returns false for any pair of elements.
type Set []Element

// HasElement returns true if set s contains element e.
//
// Mathematically, this is the fundamental binary relation defining sets.
func (s Set) HasElement(e Element) bool {
	for _, ex := range s {
		if e.Equal(ex) {
			return true
		}
	}
	return false
}

// Equal satisfies the Element interface.
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

// AddElement adds a single element to a set.
func (p *Set) AddElement(e Element) {
	if !p.HasElement(e) {
		// always allocate new backing array to avoid strange case
		// of self-referential sets.
		s := *p
		*p = append(s[:len(s):len(s)], e)
	}
}

// RemoveElement removes a single element from a set.
//
// If the element is not in the set, the method has no effect.
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

// String satisfies fmt.Stringer, providing a printable representation of a set.
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

// PowerSet returns the power set of a set.
//
// The power set of s is the set of all possible subsets of s.
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
