package response

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

type DeviceList struct {
	NextToken string           `json:"NextToken,omitempty"`
	Page      int              `json:"Page"`
	PageCount int              `json:"PageCount"`
	PageSize  int              `json:"PageSize"`
	List      []iot.DeviceInfo `json:"List"`
}

func FromQueryDeviceResponse(res *iot.QueryDeviceResponse) (*DeviceList, error) {
	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	list := &DeviceList{}
	err = json.Unmarshal(b, list)

	list.List = res.Data.DeviceInfo

	return list, err
}
