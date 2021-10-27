package manager

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
	"github.com/zhangsq-ax/aliyun-iot-manager/options"
)

// ProductList 查询产品列表
func (m *Manager) ProductList(opts *options.ProductListOptions) (*iot.DataInQueryProductList, error) {
	if opts == nil {
		opts = &options.ProductListOptions{}
	}
	req := opts.GenerateRequest(m.iotInstanceId)
	res, err := m.client.QueryProductList(req)
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return &res.Data, nil
}

// QueryProduct 查询产品详情
func (m *Manager) QueryProduct(productKey string) (*iot.DataInQueryProduct, error) {
	req := iot.CreateQueryProductRequest()
	req.AcceptFormat = "json"
	req.IotInstanceId = m.iotInstanceId
	req.ProductKey = productKey

	res, err := m.client.QueryProduct(req)
	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s - %s", res.Code, res.ErrorMessage)
	}

	return &res.Data, nil
}
