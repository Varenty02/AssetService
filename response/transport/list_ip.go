package transport

import (
	"AssetService/common"
	appctx "AssetService/component"
	"AssetService/response/business"
	"AssetService/response/tool"
	"errors"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

func ListIps(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uncover := ctx.GetUncover()
		engine := tool.NewUncoverEngine(uncover)
		runner := ctx.GetRunner()
		subFinder := tool.NewSubFinder(runner)
		httpxrunner := ctx.GetHTTPXRunner()
		httpxEn := tool.NewHTTPXEngine(httpxrunner)
		biz := business.NewAssetBiz(engine, subFinder, httpxEn)
		queryIP := c.Query("ip")
		queryDomain := c.Query("domain")
		ips, domain, err := handle(queryIP, queryDomain)
		if err != nil {
			panic(err)
		}
		res, err := biz.ExecuteUncoverBiz(c.Request.Context(), ips)
		if err != nil {
			panic(err)
		}
		subs, err := biz.EnumerateSubDomain(c.Request.Context(), domain)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res, subs))
	}
}

func handle(queryIP, queryDomain string) ([]string, string, error) {
	if queryIP != "" && queryDomain != "" {
		ipAddr := net.ParseIP(queryIP)
		if ipAddr != nil {
			ips, err := net.LookupIP(queryDomain)
			if err == nil {
				for _, ip := range ips {
					if ip.Equal(ipAddr) {
						return []string{ip.String()}, queryDomain, nil
					}
				}
			}
			return nil, "", common.NewCustomError(err, "Error looking up IP for domain\n", "IpNotMatchError")

		}
		return nil, "", common.NewCustomError(errors.New("Invalid ip address"), "Invalid ip address", "InvalidIp")
	}
	if queryIP != "" {
		return []string{queryIP}, "", nil
	}
	if queryDomain != "" {
		ips, err := net.LookupIP(queryDomain)
		if err != nil {
			panic(err)
		}
		stringArray := make([]string, len(ips))

		for i, ip := range ips {
			stringArray[i] = ip.String()
		}
		return stringArray, queryDomain, nil
	}
	return nil, "", common.NewCustomError(errors.New("Missing ip address or domain"), "Missing ip address or domain", "MissIpOrDomain")
}
