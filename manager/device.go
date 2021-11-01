package manager

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
	"github.com/zhangsq-ax/aliyun-iot-manager/options"
	"github.com/zhangsq-ax/aliyun-iot-manager/response"
)

// RegisterDevice 注册设备
func (m *Manager) RegisterDevice(opts *options.RegisterDeviceOptions) (*response.DeviceInfo, error) {
	req := opts.GenerateRequest(m.iotInstanceId)
	res, err := m.client.RegisterDevice(req)
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return response.DeviceInfoFrom(res.Data)
}

// QueryDeviceInfo 查询设备基本信息
func (m *Manager) QueryDeviceInfo(opts *options.QueryDeviceOptions) (*response.DeviceInfo, error) {
	req, err := opts.GenerateQueryDeviceInfoRequest(m.iotInstanceId)
	if err != nil {
		return nil, err
	}
	res, err := m.client.QueryDeviceInfo(req)
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return response.DeviceInfoFrom(res.Data)
}

// QueryDeviceDetail 查询设备详细信息
func (m *Manager) QueryDeviceDetail(opts *options.QueryDeviceOptions) (*iot.DataInQueryDeviceDetail, error) {
	req, err := opts.GenerateQueryDeviceDetailRequest(m.iotInstanceId)
	if err != nil {
		return nil, err
	}
	res, err := m.client.QueryDeviceDetail(req)
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return &res.Data, nil
}

// DeviceList 查询设备列表
func (m *Manager) DeviceList(opts *options.DeviceListOptions) (*response.DeviceList, error) {
	req := opts.GenerateRequest(m.iotInstanceId)
	res, err := m.client.QueryDevice(req)
	if err != nil {
		return nil, err
	}

	return response.FromQueryDeviceResponse(res)
}

// EnableThing 解除设备禁用状态
func (m *Manager) EnableThing(opts *options.QueryDeviceOptions) error {
	req, err := opts.GenerateEnableThingRequest(m.iotInstanceId)
	if err != nil {
		return err
	}
	res, err := m.client.EnableThing(req)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return nil
}

// DisableThing 禁用设备
func (m *Manager) DisableThing(opts *options.QueryDeviceOptions) error {
	req, err := opts.GenerateDisableThingRequest(m.iotInstanceId)
	if err != nil {
		return err
	}
	res, err := m.client.DisableThing(req)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return nil
}

// BatchGetDeviceState 批量获取设备状态
func (m *Manager) BatchGetDeviceState(opts *options.BatchDeviceStateOptions) (*[]iot.DeviceStatus, error) {
	req, err := opts.GenerateRequest(m.iotInstanceId)
	if err != nil {
		return nil, err
	}

	res, err := m.client.BatchGetDeviceState(req)
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return &res.DeviceStatusList.DeviceStatus, nil
}

// BatchUpdateDeviceNickname 批量修改设备备注名称
func (m *Manager) BatchUpdateDeviceNickname(opts *options.BatchUpdateDeviceNicknameOptions) error {
	req, err := opts.GenerateRequest(m.iotInstanceId)
	if err != nil {
		return err
	}
	res, err := m.client.BatchUpdateDeviceNickname(req)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return nil
}

// DeleteDevice 删除设备
func (m *Manager) DeleteDevice(opts *options.QueryDeviceOptions) error {
	req, err := opts.GenerateDeleteDeviceRequest(m.iotInstanceId)
	if err != nil {
		return err
	}

	res, err := m.client.DeleteDevice(req)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return nil
}
