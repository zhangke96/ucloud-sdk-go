// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// AllocateEIPRequest is request schema for AllocateEIP action
type AllocateEIPRequest struct {
	request.CommonBase

	// [公共参数] 地域。枚举值：cn,表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，表示中国；
	// Zone *string `required:"true"`

	// 带宽，默认值1，默认范围1~100
	Bandwidth *int `required:"true"`

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType *string `required:"true"`

	// 名称
	Name *string `required:"true"`

	// 线路。目前支持Bgp
	OperatorName *string `required:"true"`

	// 购买时长。默认值1。小时不生效，月范围【1，11】，年范围【1，5】。
	Quantity *int `required:"false"`
}

// AllocateEIPResponse is response schema for AllocateEIP action
type AllocateEIPResponse struct {
	response.CommonBase

	// 申请的EIP的ID
	EIPID string

	// 返回信息描述。
	Message string
}

// NewAllocateEIPRequest will create request of AllocateEIP action.
func (c *UCloudStackClient) NewAllocateEIPRequest() *AllocateEIPRequest {
	req := &AllocateEIPRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// AllocateEIP - 申请UCloudStack外网IP
func (c *UCloudStackClient) AllocateEIP(req *AllocateEIPRequest) (*AllocateEIPResponse, error) {
	var err error
	var res AllocateEIPResponse

	reqCopier := *req

	err = c.Client.InvokeAction("AllocateEIP", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
