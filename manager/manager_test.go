package manager

import (
	"encoding/json"
	"github.com/zhangsq-ax/aliyun-iot-manager/options"
	"os"
	"testing"
)

var manager *Manager
var productId string

var (
	RegionId        = os.Getenv("ALIYUN_REGION_ID")
	InstanceId      = os.Getenv("ALIYUN_IOT_INSTANCE_ID")
	AccessKeyId     = os.Getenv("ALIYUN_ACCESS_KEY_ID")
	AccessKeySecret = os.Getenv("ALIYUN_ACCESS_KEY_SECRET")
	ProductKey      = os.Getenv("ALIYUN_IOT_PRODUCT_KEY")
	DeviceName      = os.Getenv("ALIYUN_IOT_DEVICE_NAME")
)

func TestNewManager(t *testing.T) {
	var err error
	manager, err = NewManager(RegionId, InstanceId, AccessKeyId, AccessKeySecret)
	if err != nil {
		t.Error(err)
	}
}

func TestManager_ProductList(t *testing.T) {
	r, err := manager.ProductList(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(r.Total)

	productId = r.List.ProductInfo[0].ProductKey
}

func TestManager_QueryProduct(t *testing.T) {
	r, err := manager.QueryProduct(productId)
	if err != nil {
		t.Error(err)
	}
	t.Log(r.ProductName, r.ProductKey, r.ProductSecret)
}

func TestManager_RegisterDevice(t *testing.T) {
	deviceInfo, err := manager.RegisterDevice(&options.RegisterDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
		Nickname:   "测试设备",
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(deviceInfo.ProductKey)
	t.Log(deviceInfo.DeviceName)
	t.Log(deviceInfo.DeviceSecret)
	t.Log(deviceInfo.IotId)
	t.Log(deviceInfo.Nickname)
}

func TestManager_QueryDeviceInfo(t *testing.T) {
	deviceInfo, err := manager.QueryDeviceInfo(&options.QueryDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(deviceInfo.ProductKey)
	t.Log(deviceInfo.DeviceName)
	t.Log(deviceInfo.DeviceSecret)
	t.Log(deviceInfo.IotId)
	t.Log(deviceInfo.Nickname)
}

func TestManager_QueryDeviceDetail(t *testing.T) {
	deviceDetail, err := manager.QueryDeviceDetail(&options.QueryDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("ProductKey", deviceDetail.ProductKey)
	t.Log("ProductName", deviceDetail.ProductName)
	t.Log("DeviceName", deviceDetail.DeviceName)
	t.Log("DeviceSecret", deviceDetail.DeviceSecret)
	t.Log("IotId", deviceDetail.IotId)
	t.Log("Nickname", deviceDetail.Nickname)
	t.Log("NodeType", deviceDetail.NodeType)
	t.Log("FirmwareVersion", deviceDetail.FirmwareVersion)
	t.Log("IpAddress", deviceDetail.IpAddress)
	t.Log("Owner", deviceDetail.Owner)
	t.Log("Region", deviceDetail.Region)
	t.Log("Status", deviceDetail.Status)
	t.Log("GmtCreate", deviceDetail.GmtCreate)
	t.Log("GmtActive", deviceDetail.GmtActive)
	t.Log("GmtOnline", deviceDetail.GmtOnline)
	t.Log("UtcCreate", deviceDetail.UtcCreate)
	t.Log("UtcActive", deviceDetail.UtcActive)
	t.Log("UtcOnline", deviceDetail.UtcOnline)
}

func TestManager_DeviceList(t *testing.T) {
	list, err := manager.DeviceList(&options.DeviceListOptions{
		ProductKey: ProductKey,
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(list.PageSize)
	t.Log(list.PageCount)
	t.Log(list.Page)
	t.Log(list.NextToken)
	t.Log(len(list.List))
	t.Log(list.List[0])

	l, err := json.Marshal(list)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(l))
}

func TestManager_BatchGetDeviceState(t *testing.T) {
	status, err := manager.BatchGetDeviceState(&options.BatchDeviceStateOptions{
		ProductKey: ProductKey,
		DeviceName: &[]string{
			DeviceName,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	s := *status
	t.Log(s)
	t.Log(s[0])
	t.Log(s[0].Status)
}

func TestManager_DisableThing(t *testing.T) {
	err := manager.DisableThing(&options.QueryDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
	})
	if err != nil {
		t.Error(err)
	}

	status, err := manager.BatchGetDeviceState(&options.BatchDeviceStateOptions{
		ProductKey: ProductKey,
		DeviceName: &[]string{
			DeviceName,
		},
	})
	if err != nil {
		t.Error(err)
	}

	s := *status
	if s[0].Status != "DISABLE" {
		t.Errorf("Failed to disable device: %s", s[0].Status)
	}
}

func TestManager_EnableThing(t *testing.T) {
	err := manager.EnableThing(&options.QueryDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
	})
	if err != nil {
		t.Error(err)
	}

	status, err := manager.BatchGetDeviceState(&options.BatchDeviceStateOptions{
		ProductKey: ProductKey,
		DeviceName: &[]string{
			DeviceName,
		},
	})
	if err != nil {
		t.Error(err)
	}

	s := *status
	if s[0].Status != "UNACTIVE" {
		t.Errorf("Failed to enable device: %s", s[0].Status)
	}
}

func TestManager_BatchUpdateDeviceNickname(t *testing.T) {
	err := manager.BatchUpdateDeviceNickname(&options.BatchUpdateDeviceNicknameOptions{
		&options.UpdateDeviceNicknameOptions{
			ProductKey: ProductKey,
			DeviceName: DeviceName,
			Nickname:   "非测试设备",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	info, err := manager.QueryDeviceInfo(&options.QueryDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(info.Nickname)
	if info.Nickname != "非测试设备" {
		t.Error("Failed to update device nickname")
	}
}

func TestManager_DeleteDevice(t *testing.T) {
	err := manager.DeleteDevice(&options.QueryDeviceOptions{
		ProductKey: ProductKey,
		DeviceName: DeviceName,
	})
	if err != nil {
		t.Error(err)
	}
}
