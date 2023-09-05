package tool

import (
	"AssetService/common"
	"AssetService/response"
	"context"
	"encoding/json"
	"github.com/projectdiscovery/uncover"
	"github.com/projectdiscovery/uncover/sources"
)

func removeDuplicates(input []string) []string {
	// Tạo một map để lưu trữ các giá trị duy nhất.
	uniqueMap := make(map[string]bool)
	var uniqueSlice []string

	for _, item := range input {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueSlice = append(uniqueSlice, item)
		}
	}

	return uniqueSlice
}

type uncoverEngine struct {
	unc *uncover.Service
}

func NewUncoverEngine(unc *uncover.Service) *uncoverEngine {
	return &uncoverEngine{
		unc: unc,
	}
}
func (uncEn *uncoverEngine) ExecuteCallBack(ctx context.Context, ips []string) ([]response.Response, error) {
	var listResponse []response.Response
	for _, ip := range ips {
		res, subnets, err := uncEn.fetchData(ctx, ip)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		subnets = removeDuplicates(subnets)
		for _, subnet := range subnets {
			res.SameSubnet = append(res.SameSubnet, response.SameSubnet{
				Prefix: subnet,
			})
		}

		listResponse = append(listResponse, *res)
	}
	return listResponse, nil
}
func (uncEn *uncoverEngine) fetchData(ctx context.Context, ip string) (*response.Response, []string, error) {
	var data response.Data
	var listData []response.Data
	var subnets []string
	uncEn.unc.Options.Queries = []string{ip}
	result := func(result sources.Result) {
		_ = json.Unmarshal([]byte(result.RawData()), &data)
		listData = append(listData, data)
		subnets = append(subnets, data.ASN.Prefix)
	}
	if err := uncEn.unc.ExecuteWithCallback(ctx, result); err != nil {
		return nil, nil, err
	}
	newData := common.GetUniqueData(listData)
	res := &response.Response{
		Ip:   ip,
		Data: newData,
	}
	return res, subnets, nil
}
