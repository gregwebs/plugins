// Code generated by goa v3.8.2, DO NOT EDIT.
//
// archiver go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/archiver/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/archiver

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/goakit/examples/fetcher/archiver/gen/http/archiver/server"
)

// EncodeArchiveResponse returns a go-kit EncodeResponseFunc suitable for
// encoding archiver archive responses.
func EncodeArchiveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeArchiveResponse(encoder)
}

// DecodeArchiveRequest returns a go-kit DecodeRequestFunc suitable for
// decoding archiver archive requests.
func DecodeArchiveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeArchiveRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeReadResponse returns a go-kit EncodeResponseFunc suitable for encoding
// archiver read responses.
func EncodeReadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeReadResponse(encoder)
}

// DecodeReadRequest returns a go-kit DecodeRequestFunc suitable for decoding
// archiver read requests.
func DecodeReadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeReadRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeReadError returns a go-kit EncodeResponseFunc suitable for encoding
// errors returned by the archiver read endpoint.
func EncodeReadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) kithttp.ErrorEncoder {
	enc := server.EncodeReadError(encoder, formatter)
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		enc(ctx, w, err)
	}
}
