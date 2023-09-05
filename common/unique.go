package common

import "AssetService/response"

func GetUniqueData(input []response.Data) []response.Data {
	seenIps := make(map[string]bool)
	uniqueDatas := []response.Data{}

	for _, res := range input {
		if !seenIps[res.IP] {
			seenIps[res.IP] = true
			uniqueDatas = append(uniqueDatas, res)
		}
	}

	return uniqueDatas
}
