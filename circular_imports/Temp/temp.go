package Temp

import (
	"github.com/madhukirans/golang-examples/circular_imports/a"
	"github.com/madhukirans/golang-examples/circular_imports/b"
)

type XXX struct {
	a.A
}

type X struct {
	a.A
	b.B
}

