//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UFS DescribeUFSVolume

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUFSVolumeRequest is request schema for DescribeUFSVolume action
type DescribeUFSVolumeRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 文件系统ID
	VolumeId *string `required:"false"`

	// 文件列表起始
	Offset *int `required:"false"`

	// 文件列表长度
	Limit *int `required:"false"`
}

// DescribeUFSVolumeResponse is response schema for DescribeUFSVolume action
type DescribeUFSVolumeResponse struct {
	response.CommonBase

	// 文件系统总数
	TotalCount int

	// 文件系统详细信息列表
	DataSet []UFSVolumeInfo
}

// NewDescribeUFSVolumeRequest will create request of DescribeUFSVolume action.
func (c *UFSClient) NewDescribeUFSVolumeRequest() *DescribeUFSVolumeRequest {
	req := &DescribeUFSVolumeRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUFSVolume - 获取文件系统列表
func (c *UFSClient) DescribeUFSVolume(req *DescribeUFSVolumeRequest) (*DescribeUFSVolumeResponse, error) {
	var err error
	var res DescribeUFSVolumeResponse

	err = c.client.InvokeAction("DescribeUFSVolume", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
