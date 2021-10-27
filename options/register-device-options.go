package options

import "github.com/aliyun/alibaba-cloud-sdk-go/services/iot"

// RegisterDeviceOptions 注册设备设置选项
type RegisterDeviceOptions struct {
	ProductKey string `json:"productKey"`         // 设备所属产品标识
	DeviceName string `json:"deviceName"`         // 设备名称，设备的唯一标识
	Nickname   string `json:"nickname,omitempty"` // 设备的备注名称
}

// GenerateRequest 生成注册设备请求
func (opts *RegisterDeviceOptions) GenerateRequest(iotInstanceId string) *iot.RegisterDeviceRequest {
	req := iot.CreateRegisterDeviceRequest()
	req.AcceptFormat = "json"
	req.ProductKey = opts.ProductKey
	req.DeviceName = opts.DeviceName
	if opts.Nickname != "" {
		req.Nickname = opts.Nickname
	}
	req.IotInstanceId = iotInstanceId
	return req
}
