package appctx

import (
	httpxrunner "github.com/projectdiscovery/httpx/runner"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"github.com/projectdiscovery/uncover"
)

type appContext struct {
	unc         *uncover.Service
	runner      *runner.Runner
	httpxrunner *httpxrunner.Runner
}
type AppContext interface {
	GetUncover() *uncover.Service
	GetRunner() *runner.Runner
	GetHTTPXRunner() *httpxrunner.Runner
}

func NewAppcontext(unc *uncover.Service, runner *runner.Runner, httpxrunner *httpxrunner.Runner) *appContext {
	return &appContext{
		unc:         unc,
		runner:      runner,
		httpxrunner: httpxrunner,
	}
}

func (appCtx appContext) GetUncover() *uncover.Service {
	return appCtx.unc
}
func (appCtx appContext) GetRunner() *runner.Runner {
	return appCtx.runner
}
func (appCtx appContext) GetHTTPXRunner() *httpxrunner.Runner {
	return appCtx.httpxrunner
}
