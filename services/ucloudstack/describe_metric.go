// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeMetricRequest is request schema for DescribeMetric action
type DescribeMetricRequest struct {
	request.CommonBase

	// [公共参数] 地域。枚举值：cn，表示中国；
	// Region *string `required:"true"`

	// [公共参数] 可用区。枚举值：zone-01，中国；
	// Zone *string `required:"true"`

	// 开始时间。使用unix时间戳
	BeginTime *string `required:"true"`

	// 结束时间。使用Unix时间戳
	EndTime *string `required:"true"`

	// 监控指标。获取虚拟机监控信息调用举例，MetricName.0="CPUUtilization"、MetricName.0="MemUsage"。虚拟机监控指标枚举值：BlockProcessCount，表示阻塞进程数；CPUUtilization，表示CPU使用率；DiskReadOps，表示磁盘读次数；DiskWriteOps，表示磁盘写次数；IORead，表示磁盘读吞吐；IOWrite，表示磁盘写吞吐；LoadAvg，表示平均负载1分钟；MemUsage，表示内存使用率；NetPacketIn，表示网卡入包量；NetPacketOut，表示网卡出包量；NICIn，表示网卡入带宽；NICOut，表示网卡出带宽；SpaceUsage，表示空间使用率；TCPConnectCount，表示TCP连接数；
	MetricName []string `required:"true"`

	// 资源ID
	ResourceID *string `required:"true"`
}

// DescribeMetricResponse is response schema for DescribeMetric action
type DescribeMetricResponse struct {
	response.CommonBase

	// 返回信息列表
	Infos []MetricInfo

	// 返回信息描述
	Message string

	// 返回监控信息条数
	TotalCount int
}

// NewDescribeMetricRequest will create request of DescribeMetric action.
func (c *UCloudStackClient) NewDescribeMetricRequest() *DescribeMetricRequest {
	req := &DescribeMetricRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeMetric - 获取资源监控信息
func (c *UCloudStackClient) DescribeMetric(req *DescribeMetricRequest) (*DescribeMetricResponse, error) {
	var err error
	var res DescribeMetricResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeMetric", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
