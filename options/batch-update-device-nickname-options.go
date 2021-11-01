package options

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// UpdateDeviceNicknameOptions 更新设备备注名称选项
type UpdateDeviceNicknameOptions struct {
	ProductKey string `json:"ProductKey,omitempty"` // 要查询设备所属产品标识，需同时设置 DeviceName 参数
	DeviceName string `json:"DeviceName,omitempty"` // 要查看运行状态的设备的名称列表，需同时设置 ProductKey 参数
	IotId      string `json:"IotId,omitempty"`      // 要查看运行状态的设备 ID 列表
	Nickname   string `json:"Nickname"`             // 新的设备备注名称
}

func (opts *UpdateDeviceNicknameOptions) GenerateDeviceNicknameInfo() (*iot.BatchUpdateDeviceNicknameDeviceNicknameInfo, error) {
	info := &iot.BatchUpdateDeviceNicknameDeviceNicknameInfo{}
	if opts.ProductKey == "" && opts.DeviceName == "" && opts.IotId == "" {
		return nil, fmt.Errorf("invalid conditions")
	} else if opts.IotId != "" {
		info.IotId = opts.IotId
	} else if opts.ProductKey != "" && opts.DeviceName != "" {
		info.ProductKey = opts.ProductKey
		info.DeviceName = opts.DeviceName
	} else {
		return nil, fmt.Errorf("invalid conditions")
	}
	info.Nickname = opts.Nickname

	return info, nil
}

// BatchUpdateDeviceNicknameOptions 批量更新设备备注名称选项
type BatchUpdateDeviceNicknameOptions []*UpdateDeviceNicknameOptions

func (opts *BatchUpdateDeviceNicknameOptions) GenerateRequest(iotInstanceId string) (*iot.BatchUpdateDeviceNicknameRequest, error) {
	req := iot.CreateBatchUpdateDeviceNicknameRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = iotInstanceId
	deviceNicknameInfo := make([]iot.BatchUpdateDeviceNicknameDeviceNicknameInfo, 0)
	for _, opt := range *opts {
		info, err := opt.GenerateDeviceNicknameInfo()
		if err != nil {
			return nil, err
		}
		deviceNicknameInfo = append(deviceNicknameInfo, *info)
	}
	req.DeviceNicknameInfo = &deviceNicknameInfo

	return req, nil
}
