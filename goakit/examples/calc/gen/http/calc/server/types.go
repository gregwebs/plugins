// Code generated by goa v3.9.0, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/calc

package server

import (
	calc "goa.design/plugins/v3/goakit/examples/calc/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v
}
