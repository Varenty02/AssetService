package main

import (
	appctx "AssetService/component"
	"github.com/gin-gonic/gin"
	httpxrunner "github.com/projectdiscovery/httpx/runner"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"github.com/projectdiscovery/uncover"
	"github.com/projectdiscovery/uncover/sources"
	"log"
)

type Data struct {
}

func main() {

	//Uncover config
	opts := uncover.Options{
		Agents:   []string{"censys"},
		Limit:    1000,
		MaxRetry: 2,
		Timeout:  20,
	}
	unc, err := uncover.New(&opts)
	if err != nil {
		panic(err)
	}
	//subfinder config
	subfinderOpts := &runner.Options{
		Threads:            10,
		Timeout:            30,
		MaxEnumerationTime: 10,
		RemoveWildcard:     true,
		OnlyRecursive:      true,
	}
	log.SetFlags(0)

	subfinder, err := runner.NewRunner(subfinderOpts)
	if err != nil {
		log.Fatalf("failed to create subfinder runner: %v", err)
	}
	//httpx
	//gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)
	options := httpxrunner.Options{}
	httpxRunner, err := httpxrunner.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	// logic code
	appCtx := appctx.NewAppcontext(unc, subfinder, httpxRunner)
	//start server
	r := gin.Default()
	v1 := r.Group("/v1")
	SetupRoute(appCtx, v1)
	r.Run(":3009")
	sources.NewProvider()
}
