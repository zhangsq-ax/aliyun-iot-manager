package options

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// DeviceListOptions 获取设备列表设置选项
type DeviceListOptions struct {
	ProductKey  string `json:"ProductKey"`            // 要查询的设备所属产品标识
	PageSize    int    `json:"PageSize,omitempty"`    // 每页记录数，最大值是 50， 默认值是 10
	CurrentPage int    `json:"CurrentPage,omitempty"` // 当前页数，默认为 1
	NextToken   string `json:"NextToken,omitempty"`   // 下一页标识，首次查询无需传入。后续查询需使用的NextToken，要从上一次查询的返回结果中获取
}

// GenerateRequest 生成获取设备列表请求
func (opts *DeviceListOptions) GenerateRequest(iotInstanceId string) *iot.QueryDeviceRequest {
	req := iot.CreateQueryDeviceRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	req.ProductKey = opts.ProductKey
	if opts.PageSize < 1 {
		opts.PageSize = 10
	} else if opts.PageSize > 50 {
		opts.PageSize = 50
	}
	req.PageSize = requests.NewInteger(opts.PageSize)
	if opts.CurrentPage < 1 {
		opts.CurrentPage = 1
	}
	req.CurrentPage = requests.NewInteger(opts.CurrentPage)
	if opts.NextToken != "" {
		req.NextToken = opts.NextToken
	}

	return req
}
