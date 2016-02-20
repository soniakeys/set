// Copyright 2016 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package set

import (
	"fmt"
	"math/rand"
)

// SetM implements a superset of methods found in a number of popular set
// packages:
//   github.com/deckarep/golang-set
//   github.com/dropbox/godropbox/container/set
//   gopkg.in/fatih/set.v0
//   github.com/Workiva/go-datastructures/set
//   github.com/shopsmart/set
//
// The SetM type is defined identically to the Set type of this package.
type SetM []Element

// NewSetM returns a new set with the given elements.
//
// While a SetM is easy enough to construct as a slice, this constructor
// adds each element individually, preventing duplicate elements in the
// result.
func NewSetM(es ...Element) SetM {
	var n SetM
	for _, e := range es {
		n.Add(e)
	}
	return n
}

// Add adds a single element to a set.
//
// Returns true if e was added.  Returns false if e was already present.
//
// See SetM.AddV for a variadic version.
func (p *SetM) Add(e Element) bool {
	if p.HasElement(e) {
		return false
	}
	// always allocate new backing array.
	s1 := *p
	s2 := make(SetM, len(s1)+1)
	copy(s2, s1)
	s2[len(s1)] = e
	*p = s2
	return true
}

// AddV adds multiple elements to a set.
//
// Returns true if any element was added.  Returns false if all argument
// elements were already present.
//
// See SetM.Add for a non-variadic single element version.
func (r *SetM) AddV(es ...Element) bool {
	i := 0
	for ; ; i++ {
		if i == len(es) {
			return false
		}
		if !r.HasElement(es[i]) {
			break
		}
	}
	// copy to allocate new backing array.
	s := r.Copy()
	s = append(s, es[i])
	for {
		i++
		if i == len(es) {
			*r = s
			return true
		}
		if !s.HasElement(es[i]) {
			s = append(s, es[i])
		}
	}
}

// Cardinality returns the number of elements in the set.
//
// With the slice type, this is simply len(s).
func (s SetM) Cardinality() int { return len(s) }

// OrderedPair is an ordered pair of elements.
type OrderedPair struct {
	a, b Element
}

// Equal satisfies the Element interface, allowing OrderedPair values to be
// elements of SetMs.
func (p OrderedPair) Equal(q Element) bool {
	r, ok := q.(OrderedPair)
	return ok && p == r
}

// CartesianProduct returns a new set containing the cartesian product of s
// and t.
//
// Elements of the result will have the dynamic type OrderedPair.
func (s SetM) CartesianProduct(t SetM) SetM {
	p := make(SetM, len(s)*len(t))
	i := 0
	for _, es := range s {
		for _, et := range t {
			p[i] = OrderedPair{es, et}
			i++
		}
	}
	return p
}

// Clear removes all elements from the set, leaving the emtpy set.
func (p *SetM) Clear() {
	// truncating the slice is tempting but not safe.
	*p = nil
}

// Copy returns a copy, or a clone of a set.
//
// The returned Set is based on a newly allocated slice.  Elements though
// are the shared.  That is, this is a shallow copy.
func (s SetM) Copy() SetM { return append(SetM{}, s...) }

// Contains tests whether the given elements are all in the set.
//
// See SetM.HasElement for the non-variadic single element version.
func (s SetM) Contains(es ...Element) bool {
	for _, e := range es {
		if !s.HasElement(e) {
			return false
		}
	}
	return true
}

// Difference returns a new set containing elements of s not in t.
//
// The algorithm here starts with an empty set and adds elements of s not
// in t.
//
// See SetM.Difference2 for a different algorithm.
//
// See SetM.DifferenceV for a variadic version.
func (s SetM) Difference(t SetM) (d SetM) {
	for _, e := range s {
		if !t.HasElement(e) {
			d = append(d, e)
		}
	}
	return
}

// DifferenceV returns a new set containing elements that are in s but not in
// any of the sets ts.
//
// See SetM.Difference for a non-variadic version.
func (s SetM) DifferenceV(ts ...SetM) (d SetM) {
s:
	for _, e := range s {
		for _, t := range ts {
			if t.HasElement(e) {
				continue s
			}
		}
		d = append(d, e)
	}
	return
}

// Difference2 returns a new set containing elements of s not in c.
//
// The algorithm here starts with a copy of s and removes elements of c.
func (s SetM) Difference2(c SetM) SetM {
	d := s.Copy()
	for _, e := range c {
		d.Remove(e)
	}
	return d
}

// DifferenceR modifies receiver r to be the difference r - c.
//
// Commonly called Subtract in other packages.
func (r *SetM) DifferenceR(c SetM) {
	for _, e := range c {
		r.Remove(e)
	}
}

// Do calls f on each element of s, in random order.
func (s SetM) Do(f func(Element)) {
	for _, i := range rand.Perm(len(s)) {
		f(s[i])
	}
}

// DoWhile calls f on each element of s, in random order, as long as f returns
// true.
//
// DoWhile returns true if f returns true for all elements of s.
// If f returns false for an element, DoWhile returns false immediately
// without calling f on any remaining elements.
func (s SetM) DoWhile(f func(Element) bool) bool {
	for _, i := range rand.Perm(len(s)) {
		if !f(s[i]) {
			return false
		}
	}
	return true
}

// Equal implements set equality similar to other packages, but whereas the
// argument in other packages typically has a set type, here it is the Element
// interface.  This allows SetM.Equal to implement the Element interface, which
// allows sets of sets, for example power sets.
//
// If the dynamic type of the argument is not SetM or if the sets are not
// equal, Equal returns false.  (It doesn't panic.)
func (s SetM) Equal(e Element) bool {
	t, ok := e.(SetM)
	if !ok {
		return false
	}
	if len(s) != len(t) {
		return false
	}
	return s.Contains(t...)
}

// Filter returns the subset containing elements e of s where f(e) is true.
func (s SetM) Filter(f func(Element) bool) (r SetM) {
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return
}

// Flatten flattens a set potentially containing nested sets.
//
// Flatten returns a new non-nested set containing all non-set elements in s
// at any level of nesting.
//
// See SetM.Flatten2 for an alternative algorithm.
func (s SetM) Flatten() (f SetM) {
	for _, e := range s {
		if s2, ok := e.(SetM); ok {
			f.UnionR(s2.Flatten())
		} else {
			f.Add(e)
		}
	}
	return
}

// Flatten2 flattens a set potentially containing nested sets.
//
// It's a version of Flatten implmented with an alternative algorithm.
func (s SetM) Flatten2() (f SetM) {
	var r func(SetM)
	r = func(SetM) {
		for _, e := range s {
			if s2, ok := e.(SetM); ok {
				r(s2)
			} else {
				f.Add(e)
			}
		}
	}
	r(s)
	return
}

// HasElement returns true if set s contains element e.
//
// Called Contains or Exists in some packages, I left this method with the
// name from my set package and used the name Contains for the variadic version.
//
// See SetM.Contains for the variadic version.
func (s SetM) HasElement(e Element) bool {
	for _, ex := range s {
		if e.Equal(ex) {
			return true
		}
	}
	return false
}

// Intersect returns a new set of elements of s also in t.
//
// The algorithm here starts with an empty set and adds elements of s found
// in t.
//
// See Intersect2 for an alternate algorithm.
//
// See IntersectV for a variadic version.
func (s SetM) Intersect(t SetM) (i SetM) {
	for _, e := range s {
		if t.HasElement(e) {
			i = append(i, e)
		}
	}
	return
}

// IntersectV returns a new set of elements of s that are also present in
// all sets ts.
func (s SetM) IntersectV(ts ...SetM) (i SetM) {
s:
	for _, e := range s {
		for _, t := range ts {
			if !t.HasElement(e) {
				continue s
			}
		}
		i = append(i, e)
	}
	return
}

// IsEmpty returns true if s is the empty set.
func (s SetM) IsEmpty() bool { return len(s) == 0 }

// Intersect2 returns a new set of elements of s also in t.
//
// The algorithm here starts with a copy of s and removes elements not found
// in t.
func (s SetM) Intersect2(t SetM) SetM {
	i := s.Copy()
	for _, e := range s {
		if !t.HasElement(e) {
			i.Remove(e)
		}
	}
	return i
}

// IsSubset returns true if s is a subset of t.
//
// A similar predicate named All is in some packages.
func (s SetM) IsSubset(t SetM) bool {
	for _, e := range s {
		if !t.HasElement(e) {
			return false
		}
	}
	return true
}

// IsSuperset returns true if s is a superset of t.
func (s SetM) IsSuperset(t SetM) bool {
	return t.IsSubset(s)
}

// Iter sends elements of s on the returned channel.
//
// The channel is unbuffered and the set is not copied.  Changes to s
// concurrent with channel receives may be reflected in the received values.
//
// The channel is closed after all elements are sent.
func (s SetM) Iter() <-chan Element {
	c := make(chan Element)
	go func() {
		for _, e := range s {
			c <- e
		}
		close(c)
	}()
	return c
}

// IterFunc returns a function that iterates over elements of s in random
// order.
//
// Like SetM.Iter, the set is not copied.  Changes to s concurrent with
// calls to the returned function may be reflected in the returned values.
//
// The ok return will be true for each element of s, then false on any
// call afterwards.
func (s SetM) IterFunc() func() (e Element, ok bool) {
	r := rand.Perm(len(s))
	i := 0
	return func() (Element, bool) {
		if i >= len(s) {
			return nil, false
		}
		e := s[r[i]]
		i++
		return e, true
	}
}

// IterBuffered sends elements of s in random order on the returned channel.
//
// The algorithm here buffers all elements of s before returning.
// Changes to s interleaved with channel receives will not be reflected
// in the received values.
//
// The channel is closed after all elements are sent.
func (s SetM) IterBuffered() <-chan Element {
	c := make(chan Element, len(s))
	for _, i := range rand.Perm(len(s)) {
		c <- s[i]
	}
	close(c)
	return c
}

// Map returns the set of distinct values f(e) for all e in s.
//
// The cardinality of the result m may be less that the cardinality of s.
func (s SetM) Map(f func(Element) Element) (m SetM) {
	for _, e := range s {
		m.Add(f(e))
	}
	return
}

// Pop returns a random element of r and removes it from r.
//
// If r is empty, Pop returns nil, false.
func (r *SetM) Pop() (e Element, ok bool) {
	s := *r
	if len(s) == 0 {
		return
	}
	i := rand.Intn(len(s))
	e = s[i]
	copy(s[i:], s[i+1:])
	last := len(s) - 1
	*r = s[:last:last]
	return e, true
}

// PowerSet returns the power set of a set.
//
// The power set of s is the set of all possible subsets of s.
func (s SetM) PowerSet() SetM {
	r := SetM{SetM{}}
	for _, es := range s {
		var u SetM
		for _, er := range r {
			u = append(u, append(er.(SetM), es))
		}
		r = append(r, u...)
	}
	return r
}

// Remove removes a single element from a set.
//
// Returns true if the element was found and removed.
// Returns false if the element was not found.
func (r *SetM) Remove(e Element) bool {
	for i, ex := range *r {
		if e.Equal(ex) {
			s := *r
			last := len(s) - 1
			s[i], s[last] = s[last], s[i]
			*r = s[:last]
			return true
		}
	}
	return false
}

// RemoveIf removes all elements where f returns true.
//
// Returns true if any element was removed.
// Returns false if no elements were removed.
func (r *SetM) RemoveIf(f func(Element) bool) (removed bool) {
	s := *r
	for i := 0; i < len(s); {
		if f(s[i]) {
			last := len(s) - 1
			s[i], s[last] = s[last], s[i]
			s = s[:last]
			removed = true
		} else {
			i++
		}
	}
	*r = s
	return
}

// String satisfies fmt.Stringer, providing a printable representation of a set.
func (s SetM) String() string {
	r := "{"
	for i, j := range rand.Perm(len(s)) {
		if i > 0 {
			r += " "
		}
		r += fmt.Sprint(s[j])
	}
	return r + "}"
}

// SymmetricDifference returns a new set with elements in s or t but not both.
func (s SetM) SymmetricDifference(t SetM) SetM {
	d := s.Copy()
	for _, e := range t {
		if s.HasElement(e) {
			d.Remove(e)
		} else {
			d = append(d, e)
		}
	}
	return d
}

// Union returns a new set with elements of s or t.
//
// See UnionR for a version that modifies the receiver.
//
// See UnionV for a variadic version.
func (s SetM) Union(t SetM) SetM {
	u := s.Copy()
	for _, e := range t {
		u.Add(e)
	}
	return u
}

// UnionR produces a union by modifying receiver r to include elements of t.
//
// Called Merge in some packages.
func (r *SetM) UnionR(t SetM) {
	for _, e := range t {
		r.Add(e)
	}
}

// UnionV returns a new set with elements of s or any of t.
func (s SetM) UnionV(ts ...SetM) SetM {
	u := s.Copy()
	for _, t := range ts {
		for _, e := range t {
			u.Add(e)
		}
	}
	return u
}
