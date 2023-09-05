package tool

import (
	"AssetService/common"
	"bytes"
	"context"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"io"
	"strings"
)

type subFinder struct {
	runner *runner.Runner
}

func NewSubFinder(runner *runner.Runner) *subFinder {
	return &subFinder{
		runner: runner,
	}
}
func (finder *subFinder) Enumerate(ctx context.Context, domain string) ([]string, error) {
	output := &bytes.Buffer{}
	// To run subdomain enumeration on a single domain
	if err := finder.runner.EnumerateSingleDomainWithCtx(ctx, domain, []io.Writer{output}); err != nil {
		return nil, common.NewErrorResponse(err, "failed to enumerate single domain", "failed to enumerate single domain", "EnumerateFailed")
	}
	return strings.Split(output.String(), "\n"), nil
}
