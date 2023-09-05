package tool

import (
	"context"
	"fmt"
	"github.com/projectdiscovery/goflags"
	httpxrunner "github.com/projectdiscovery/httpx/runner"
	"log"
)

type httpxEngine struct {
	runner *httpxrunner.Runner
}

func NewHTTPXEngine(runner *httpxrunner.Runner) *httpxEngine {
	return &httpxEngine{
		runner: runner,
	}
}
func (e *httpxEngine) EnumerateHTTPX(ctx context.Context, CIDRs string) ([]string, error) {
	var list []string
	opt := httpxrunner.Options{
		Methods:         "GET",
		InputTargetHost: goflags.StringSlice{CIDRs},
		OnResult: func(r httpxrunner.Result) {
			// handle error
			if r.Err != nil {
				return
			}
			fmt.Printf("%s ", r.Input)
			list = append(list, r.Host)
		},
	}
	runner, err := httpxrunner.New(&opt)
	if err != nil {
		log.Fatal(err)
	}
	e.runner = runner
	defer e.runner.Close()
	e.runner.RunEnumeration()
	return list, nil
}
