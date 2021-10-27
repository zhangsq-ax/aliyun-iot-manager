package options

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// BatchDeviceStateOptions 批量获取设备状态设置选项
type BatchDeviceStateOptions struct {
	ProductKey string    `json:"ProductKey,omitempty"` // 要查询设备所属产品标识，需同时设置 DeviceName 参数
	DeviceName *[]string `json:"DeviceName,omitempty"` // 要查看运行状态的设备的名称列表，需同时设置 ProductKey 参数
	IotId      *[]string `json:"IotId,omitempty"`      // 要查看运行状态的设备 ID 列表
}

// GenerateRequest 生成批量获取设备状态请求
func (opts *BatchDeviceStateOptions) GenerateRequest(iotInstanceId string) (*iot.BatchGetDeviceStateRequest, error) {
	req := iot.CreateBatchGetDeviceStateRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	if opts.ProductKey == "" && opts.DeviceName == nil && opts.IotId == nil {
		return nil, fmt.Errorf("invalid conditions")
	} else if opts.IotId != nil {
		req.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != nil {
		req.ProductKey = opts.ProductKey
		req.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid conditions")
	}

	return req, nil
}
