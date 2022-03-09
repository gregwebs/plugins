// Code generated by goa v3.6.1, DO NOT EDIT.
//
// health go-kit HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package client

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/http/health/client"
)

// DecodeShowResponse returns a go-kit DecodeResponseFunc suitable for decoding
// health show responses.
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder) kithttp.DecodeResponseFunc {
	dec := client.DecodeShowResponse(decoder, false)
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		return dec(resp)
	}
}
