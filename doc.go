// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

// Sets as lists of interfaces.
//
// An interface is defined called Element with a single method Equal.
// This allows generality in element types and general definitions of equality.
// A set type is defined as an Element slice.  You lose O(1) access in
// comparison to maps for example, but you gain this greater generality.
//
// Usually your application will allow something more efficient.  If your
// element types are Go-comparable and the Go definition of equality for
// the type is what you want, then Go maps are hard to beat as set types.
// The math/big.Int type makes a servicable bitset in many cases.
// If your element types are ordered, then some ordered data structure
// is likely a good choice.  Just staying with the standard library,
// sort.Search can be used to maintain an ordered set.  Outside the standard
// library there are packages out there for map-based sets, bitsets, and
// ordered trees.
//
// But just to demonstrate another possibility, this package shows an example
// of sets as lists of interfaces with no requirements of ordering or
// Go-comparability.
//
// Set type
//
// The package defines two set types.  The first, just called Set, has a
// minimal number of methods, just to show the basic mathematical concepts
// represented and how they are implemented with a list of Elements.
//
// For what it's worth, the Set type has 100% test coverage.
//
// SetM type
//
// The SetM type has the same underlying type as the Set type of this package,
// just more methods.  The M of SetM is for "more methods."  It resulted from
// looking at a number of popular set packages and accumulating a rough
// superset of methods.
//
//   github.com/deckarep/golang-set
//   github.com/dropbox/godropbox/container/set
//   github.com/jchauncey/kubeclient/util/sets
//   github.com/luci/luci-go/common/stringset
//   github.com/nlandolfi/set
//   github.com/shopsmart/set
//   github.com/Workiva/go-datastructures/set
//   gopkg.in/fatih/set.v0
//
// The result shows lots of capabilities, but shouldn't be taken as anything
// like an "ultimate" set API.  It's easy to imagine more and more methods
// but the API is awflully big as it is.  A number of the methods are trivial.
//
// The SetM type is yet untested.
package set
