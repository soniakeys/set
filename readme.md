# Set

A little demonstration.

Go maps serve pretty well as sets in many cases, using Go-comparabilty for
the test of element equality.  Sometimes though you want elements that aren’t
Go comparable, or you want a different definition of equality.  This package
demonstrates sets as lists of elements where the Element interface has a
single method Equal, allowing more generality in types as set elements and
more general defintions of equality.
