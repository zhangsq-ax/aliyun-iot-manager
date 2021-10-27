package options

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// ProductListOptions 获取产品列表设置选项
type ProductListOptions struct {
	CurrentPage int `json:"currentPage,omitempty"`
	PageSize    int `json:"pageSize,omitempty"`
}

// GenerateRequest 生成获取产品列表请求
func (opts *ProductListOptions) GenerateRequest(iotInstanceId string) *iot.QueryProductListRequest {
	currentPage := opts.CurrentPage
	if currentPage < 1 {
		currentPage = 1
	}
	pageSize := opts.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	req := iot.CreateQueryProductListRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	req.CurrentPage = requests.NewInteger(currentPage)
	req.PageSize = requests.NewInteger(pageSize)

	return req
}
