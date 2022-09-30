// Code generated by goa v3.9.0, DO NOT EDIT.
//
// archiver HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/archiver/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/archiver

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
	archiver "goa.design/plugins/v3/goakit/examples/fetcher/archiver/gen/archiver"
	archiverviews "goa.design/plugins/v3/goakit/examples/fetcher/archiver/gen/archiver/views"
)

// BuildArchiveRequest instantiates a HTTP request object with method and path
// set to call the "archiver" service "archive" endpoint
func (c *Client) BuildArchiveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ArchiveArchiverPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("archiver", "archive", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeArchiveRequest returns an encoder for requests sent to the archiver
// archive server.
func EncodeArchiveRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*archiver.ArchivePayload)
		if !ok {
			return goahttp.ErrInvalidType("archiver", "archive", "*archiver.ArchivePayload", v)
		}
		body := NewArchiveRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("archiver", "archive", err)
		}
		return nil
	}
}

// DecodeArchiveResponse returns a decoder for responses returned by the
// archiver archive endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeArchiveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ArchiveResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("archiver", "archive", err)
			}
			p := NewArchiveMediaViewOK(&body)
			view := "default"
			vres := &archiverviews.ArchiveMedia{Projected: p, View: view}
			if err = archiverviews.ValidateArchiveMedia(vres); err != nil {
				return nil, goahttp.ErrValidationError("archiver", "archive", err)
			}
			res := archiver.NewArchiveMedia(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("archiver", "archive", resp.StatusCode, string(body))
		}
	}
}

// BuildReadRequest instantiates a HTTP request object with method and path set
// to call the "archiver" service "read" endpoint
func (c *Client) BuildReadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*archiver.ReadPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("archiver", "read", "*archiver.ReadPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ReadArchiverPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("archiver", "read", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeReadResponse returns a decoder for responses returned by the archiver
// read endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeReadResponse may return the following errors:
//   - "not_found" (type *goa.ServiceError): http.StatusNotFound
//   - "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeReadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ReadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("archiver", "read", err)
			}
			p := NewReadArchiveMediaOK(&body)
			view := "default"
			vres := &archiverviews.ArchiveMedia{Projected: p, View: view}
			if err = archiverviews.ValidateArchiveMedia(vres); err != nil {
				return nil, goahttp.ErrValidationError("archiver", "read", err)
			}
			res := archiver.NewArchiveMedia(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ReadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("archiver", "read", err)
			}
			err = ValidateReadNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("archiver", "read", err)
			}
			return nil, NewReadNotFound(&body)
		case http.StatusBadRequest:
			var (
				body ReadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("archiver", "read", err)
			}
			err = ValidateReadBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("archiver", "read", err)
			}
			return nil, NewReadBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("archiver", "read", resp.StatusCode, string(body))
		}
	}
}
