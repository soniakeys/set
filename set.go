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

// An Element can be an element of a Set.
//
// A type satisfying Element must implement just one method, Equal.
// Equal returns true if the argument is to be considered equal to the
// receiver in some sense.  For a valid implementation of Equal, the
// functions Reflexive, Symetric, and Transitive must return true for
// all possible receivers and arguments.
//
// Implementations will typically type assert the argument to a value
// of the receiver type and then perform some value comparison.
type Element interface {
	Equal(Element) bool
}

// Set is a type implementing the mathematical concept of a set.
// This type also implements the Element interface, so Sets can be elements
// of Sets.
//
// Set implementation is an Element slice.  Sets are conceptually unordered.
// While slices are ordered, this order in a Set must be considered irrelevant.
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
		*p = append(*p, e)
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

// Refexive validates the reflexive property of the Equal method for an element.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Reflexive(e Element) bool {
	return e.Equal(e)
}

// Symetric validates the symetric property of the Equal method for an element.
//
// The function should return true for all possible elements.
// If it returns false, it means the element's implementation of the Equal
// method is invalid.
func Symetric(a, b Element) bool {
	return a.Equal(b) == b.Equal(a)
}

// Transitive validates the transitive property of the Equal method for an element.
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
