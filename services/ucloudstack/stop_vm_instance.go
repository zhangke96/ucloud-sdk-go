// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// StopVMInstanceRequest is request schema for StopVMInstance action
type StopVMInstanceRequest struct {
	request.CommonBase

	// [公共参数] 地域。枚举值：cn，表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，表示中国；
	// Zone *string `required:"true"`

	// 虚拟机 ID
	VMID *string `required:"true"`
}

// StopVMInstanceResponse is response schema for StopVMInstance action
type StopVMInstanceResponse struct {
	response.CommonBase

	// 返回信息描述
	Message string

	// 虚拟机 ID
	VMID string
}

// NewStopVMInstanceRequest will create request of StopVMInstance action.
func (c *UCloudStackClient) NewStopVMInstanceRequest() *StopVMInstanceRequest {
	req := &StopVMInstanceRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// StopVMInstance - 关闭UCloudStack虚拟机
func (c *UCloudStackClient) StopVMInstance(req *StopVMInstanceRequest) (*StopVMInstanceResponse, error) {
	var err error
	var res StopVMInstanceResponse

	reqCopier := *req

	err = c.Client.InvokeAction("StopVMInstance", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
