package common

type successRes struct {
	Res       interface{} `json:"res"`
	SubDomain interface{} `json:"sub_domain"`
	//CIDRIp    interface{} `json:"cidr_ip,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

//filter đối soát thứ client gửi lên có phải thứ mà lưu vào server hay không
//omitempty nếu có thi lay neu khong co thì bo qua

func NewSuccessResponse(res, subDomain, paging, filter interface{}) *successRes {
	return &successRes{Res: res, SubDomain: subDomain, Paging: paging, Filter: filter}
}
func SimpleSuccessResponse(res, subDomain interface{}) *successRes {
	return NewSuccessResponse(res, subDomain, nil, nil)
}
