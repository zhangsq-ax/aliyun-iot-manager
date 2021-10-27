package response

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

type DeviceInfo struct {
	DeviceName   string `json:"DeviceName"`   // 设备名称
	DeviceSecret string `json:"DeviceSecret"` // 设备密钥
	IotId        string `json:"IotId"`        // 物联网平台为该设备颁发的设备 ID，作为该设备的唯一标识
	Nickname     string `json:"Nickname"`     // 备注名称
	ProductKey   string `json:"ProductKey"`   // 设备所属的产品标识

	/*FirmwareVersion string `json:"FirmwareVersion,omitempty"` // 设备的固件版本号
	GmtActive       string `json:"GmtActive,omitempty"`       // 设备激活时间，GMT 格式
	GmtCreate       string `json:"GmtCreate,omitempty"`       // 设备创建时间，GMT 格式
	GmtOnline       string `json:"GmtOnline,omitempty"`       // 设置最近一次上线时间，GMT 格式
	IpAddress       string `json:"IpAddress,omitempty"`       // 设备 IP 地址
	NodeType        int    `json:"NodeType,omitempty"`        // 节点类型：0 - 设备  1 - 网关
	Owner           bool   `json:"Owner,omitempty"`           // 是否是该设备拥有者
	ProductName     string `json:"ProductName,omitempty"`     // 设备所属产品名称
	Region          string `json:"Region,omitempty"`          // 设备所在地区
	Status          string `json:"Status,omitempty"`          // 设备状态：ONLINE - 在线  OFFLINE - 离线  UNACTIVE - 未激活  DISABLE - 已禁用
	UtcActive       string `json:"UtcActive,omitempty"`       // 设备激活时间，UTC 格式
	UtcCreate       string `json:"UtcCreate,omitempty"`       // 设备创建时间，UTC 格式
	UtcOnline       string `json:"UtcOnline,omitempty"`       // 设备最近一次上线时间，UTC 格式*/
}

func DeviceInfoFrom(data iot.Data) (*DeviceInfo, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	deviceInfo := &DeviceInfo{}
	err = json.Unmarshal(b, deviceInfo)

	return deviceInfo, err
}
