// Code generated by goa v3.7.3, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/plugins/v3/i18n/examples/calc/design -o
// $(GOPATH)/src/goa.design/plugins/i18n/examples/calc

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	calcc "goa.design/plugins/v3/i18n/examples/calc/gen/http/calc/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `calc add
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` calc add --a 1 --b 2` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddAFlag = calcAddFlags.String("a", "REQUIRED", "")
		calcAddBFlag = calcAddFlags.String("b", "REQUIRED", "")
	)
	calcFlags.Usage = calcUsage
	calcAddFlags.Usage = calcAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "calc":
			svcf = calcFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "calc":
			switch epn {
			case "add":
				epf = calcAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calc":
			c := calcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = calcc.BuildAddPayload(*calcAddAFlag, *calcAddBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service exposes public endpoints to do basic mathematical calculations.
Usage:
    %[1]s [globalflags] calc COMMAND [flags]

COMMAND:
    add: Add adds up the two integer parameters and returns the results.

Additional help:
    %[1]s calc COMMAND --help
`, os.Args[0])
}
func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] calc add -a INT -b INT

Add adds up the two integer parameters and returns the results.
    -a INT: 
    -b INT: 

Example:
    %[1]s calc add --a 1 --b 2
`, os.Args[0])
}
