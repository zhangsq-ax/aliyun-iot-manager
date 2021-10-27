package manager

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

// Manager 阿里云 IoT 平台管理器
type Manager struct {
	client        *iot.Client
	iotInstanceId string
}

// NewManager 创建阿里云 IoT 平台管理器实例
func NewManager(regionId, instanceId, accessKeyId, accessKeySecret string) (*Manager, error) {
	client, err := iot.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	return &Manager{
		client:        client,
		iotInstanceId: instanceId,
	}, nil
}
