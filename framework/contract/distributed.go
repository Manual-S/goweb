package contract

import "time"

const DistributedKey = "DistributedKey"

type Distributed interface {
	// Select 实现一个分布式的选举器
	Select(serviceName string, appID string, holdTime time.Duration) (string, error)
}
