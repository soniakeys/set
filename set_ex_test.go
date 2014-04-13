// Copyright 2014 Sonia Keys
// License MIT: http://opensource.org/licenses/MIT

package set_test

import (
	"fmt"
	"math/rand"

	"github.com/soniakeys/set"
)

// integer type satisfying element interface
type intEle int

func (i intEle) Equal(e set.Element) bool {
	if j, ok := e.(intEle); ok {
		return i == j
	}
	return false
}

func Example() {
	rand.Seed(123)

	var s set.Set
	fmt.Println(s)

	for _, i := range []intEle{1, 4, 4, 2, 2, 3, 4} {
		s.AddElement(i)
	}
	fmt.Println(s)

	s.RemoveElement(intEle(4))
	fmt.Println(s)

	ps := s.PowerSet()
	fmt.Println(ps)

	fmt.Println(ps.HasElement(s))
	// Output:
	// {}
	// {1 2 4 3}
	// {1 3 2}
	// {{1 3 2} {1} {2} {2 3} {3} {} {1 3} {1 2}}
	// true
}
