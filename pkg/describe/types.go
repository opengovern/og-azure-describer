package describe

import (
	"github.com/kaytu-io/kaytu-azure-describer/pkg/describe/enums"
	"github.com/kaytu-io/kaytu-azure-describer/pkg/source"
)

type DescribeJob struct {
	JobID         uint // DescribeResourceJob ID
	ScheduleJobID uint
	ParentJobID   uint // DescribeSourceJob ID
	ResourceType  string
	SourceID      string
	AccountID     string
	DescribedAt   int64
	SourceType    source.Type
	CipherText    string
	TriggerType   enums.DescribeTriggerType
	RetryCounter  uint
}

type LambdaDescribeWorkerInput struct {
	WorkspaceId      string      `json:"workspaceId"`
	DescribeEndpoint string      `json:"describeEndpoint"`
	KeyARN           string      `json:"keyARN"`
	DescribeJob      DescribeJob `json:"describeJob"`
}
