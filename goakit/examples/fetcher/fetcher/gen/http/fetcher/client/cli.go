// Code generated by goa v3.9.0, DO NOT EDIT.
//
// fetcher HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package client

import (
	goa "goa.design/goa/v3/pkg"
	fetcher "goa.design/plugins/v3/goakit/examples/fetcher/fetcher/gen/fetcher"
)

// BuildFetchPayload builds the payload for the fetcher fetch endpoint from CLI
// flags.
func BuildFetchPayload(fetcherFetchURL string) (*fetcher.FetchPayload, error) {
	var err error
	var url_ string
	{
		url_ = fetcherFetchURL
		err = goa.MergeErrors(err, goa.ValidateFormat("url_", url_, goa.FormatURI))

		if err != nil {
			return nil, err
		}
	}
	v := &fetcher.FetchPayload{}
	v.URL = url_

	return v, nil
}
