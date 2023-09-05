package business

import (
	"AssetService/response"
	"context"
)

type UncoverEngine interface {
	ExecuteCallBack(ctx context.Context, ips []string) ([]response.Response, error)
}
type SubFinder interface {
	Enumerate(ctx context.Context, domain string) ([]string, error)
}
type HTTPXRunner interface {
	EnumerateHTTPX(ctx context.Context, CIDRs string) ([]string, error)
}
type assetBiz struct {
	engine      UncoverEngine
	subFinder   SubFinder
	httpxRunner HTTPXRunner
}

func NewAssetBiz(engine UncoverEngine, subFinder SubFinder, httpxRunner HTTPXRunner) *assetBiz {
	return &assetBiz{
		engine:      engine,
		subFinder:   subFinder,
		httpxRunner: httpxRunner,
	}
}
func (biz *assetBiz) ExecuteUncoverBiz(ctx context.Context, ips []string) ([]response.Response, error) {
	datas, err := biz.engine.ExecuteCallBack(ctx, ips)
	if err != nil {
		return nil, err
	}
	for _, data := range datas {
		for i := range data.SameSubnet {
			data.SameSubnet[i].Ips, _ = biz.httpxRunner.EnumerateHTTPX(ctx, data.SameSubnet[i].Prefix)
		}
	}
	return datas, nil
}
func (biz *assetBiz) EnumerateSubDomain(ctx context.Context, domain string) ([]string, error) {
	subs, err := biz.subFinder.Enumerate(ctx, domain)
	if err != nil {
		return nil, err
	}
	return subs, nil
}
