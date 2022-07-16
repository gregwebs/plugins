// Code generated by goa v3.7.11, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/v3/cors/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/cors/examples/calc

package server

import (
	calc "goa.design/plugins/v3/cors/examples/calc/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	v := &calc.AddPayload{}
	v.A = a
	v.B = b

	return v
}
