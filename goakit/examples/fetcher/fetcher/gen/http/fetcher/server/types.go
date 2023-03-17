// Code generated by goa v3.11.2, DO NOT EDIT.
//
// fetcher HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package server

import (
	goa "goa.design/goa/v3/pkg"
	fetcher "goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/fetcher"
	fetcherviews "goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/fetcher/views"
)

// FetchResponseBody is the type of the "fetcher" service "fetch" endpoint HTTP
// response body.
type FetchResponseBody struct {
	// HTTP status code returned by fetched service
	Status int `form:"status" json:"status" xml:"status"`
	// The href to the corresponding archive in the archiver service
	ArchiveHref string `form:"archive_href" json:"archive_href" xml:"archive_href"`
}

// FetchBadRequestResponseBody is the type of the "fetcher" service "fetch"
// endpoint HTTP response body for the "bad_request" error.
type FetchBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// FetchInternalErrorResponseBody is the type of the "fetcher" service "fetch"
// endpoint HTTP response body for the "internal_error" error.
type FetchInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewFetchResponseBody builds the HTTP response body from the result of the
// "fetch" endpoint of the "fetcher" service.
func NewFetchResponseBody(res *fetcherviews.FetchMediaView) *FetchResponseBody {
	body := &FetchResponseBody{
		Status:      *res.Status,
		ArchiveHref: *res.ArchiveHref,
	}
	return body
}

// NewFetchBadRequestResponseBody builds the HTTP response body from the result
// of the "fetch" endpoint of the "fetcher" service.
func NewFetchBadRequestResponseBody(res *goa.ServiceError) *FetchBadRequestResponseBody {
	body := &FetchBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewFetchInternalErrorResponseBody builds the HTTP response body from the
// result of the "fetch" endpoint of the "fetcher" service.
func NewFetchInternalErrorResponseBody(res *goa.ServiceError) *FetchInternalErrorResponseBody {
	body := &FetchInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewFetchPayload builds a fetcher service fetch endpoint payload.
func NewFetchPayload(url_ string) *fetcher.FetchPayload {
	v := &fetcher.FetchPayload{}
	v.URL = url_

	return v
}
