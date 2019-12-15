//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDataArk DescribeUhostTmMeta

package udataark

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

type CreateVdiskSnapshotRequest struct {
	request.CommonBase

	// 盘id
	VDiskId *string `required:"true"`

	// 快照名字
	Name *string `required:"true"`

	// 快照描述
	Comment *string
}

type CreateVdiskSnapshotResponse struct {
	response.CommonBase

	SnapshotId string
}

func (c *UDataArkClient) NewCreateVdiskSnapshotRequest() *CreateVdiskSnapshotRequest {
	req := &CreateVdiskSnapshotRequest{}
	c.Client.SetupRequest(req)
	return req
}

func (c *UDataArkClient) CreateVdiskSnapshotRequest(req *CreateVdiskSnapshotRequest) (*CreateVdiskSnapshotResponse, error) {
	var err error
	var res CreateVdiskSnapshotResponse

	err = c.Client.InvokeAction("CreateUserVDiskSnapshot", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
