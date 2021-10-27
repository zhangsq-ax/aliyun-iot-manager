package options

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// QueryDeviceOptions 查询设备设置选项
type QueryDeviceOptions struct {
	ProductKey string `json:"ProductKey,omitempty"` // 要查询设备所属产品标识，需同时设置 DeviceName 参数
	DeviceName string `json:"DeviceName,omitempty"` // 要查询设备名称，需同时设置 ProductKey 参数
	IotId      string `json:"IotId,omitempty"`      // 要查询设备 ID，如果设置该参数则无需设置 ProductKey 和 DeviceName 参数，同时设置时以 IotId 为准
}

// GenerateQueryDeviceInfoRequest 生成查询设备信息请求
func (opts *QueryDeviceOptions) GenerateQueryDeviceInfoRequest(iotInstanceId string) (*iot.QueryDeviceInfoRequest, error) {
	req := iot.CreateQueryDeviceInfoRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	if opts.ProductKey == "" && opts.DeviceName == "" && opts.IotId == "" {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	} else if opts.IotId != "" {
		req.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != "" {
		req.ProductKey = opts.ProductKey
		req.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	}

	return req, nil
}

func (opts *QueryDeviceOptions) GenerateQueryDeviceDetailRequest(iotInstanceId string) (*iot.QueryDeviceDetailRequest, error) {
	req := iot.CreateQueryDeviceDetailRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	if opts.ProductKey == "" && opts.DeviceName == "" && opts.IotId == "" {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	} else if opts.IotId != "" {
		req.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != "" {
		req.ProductKey = opts.ProductKey
		req.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	}

	return req, nil
}

func (opts *QueryDeviceOptions) GenerateDeleteDeviceRequest(iotInstanceId string) (*iot.DeleteDeviceRequest, error) {
	req := iot.CreateDeleteDeviceRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	if opts.ProductKey == "" && opts.DeviceName == "" && opts.IotId == "" {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	} else if opts.IotId != "" {
		req.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != "" {
		req.ProductKey = opts.ProductKey
		req.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	}

	return req, nil
}

func (opts *QueryDeviceOptions) GenerateEnableThingRequest(iotInstanceId string) (*iot.EnableThingRequest, error) {
	req := iot.CreateEnableThingRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	if opts.ProductKey == "" && opts.DeviceName == "" && opts.IotId == "" {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	} else if opts.IotId != "" {
		req.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != "" {
		req.ProductKey = opts.ProductKey
		req.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	}

	return req, nil
}

func (opts *QueryDeviceOptions) GenerateDisableThingRequest(iotInstanceId string) (*iot.DisableThingRequest, error) {
	req := iot.CreateDisableThingRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	if opts.ProductKey == "" && opts.DeviceName == "" && opts.IotId == "" {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	} else if opts.IotId != "" {
		req.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != "" {
		req.ProductKey = opts.ProductKey
		req.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid condition: %s(ProductKey) %s(DeviceName) %s(IotId)", opts.ProductKey, opts.DeviceName, opts.IotId)
	}

	return req, nil
}
