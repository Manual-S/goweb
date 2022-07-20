// Package distributed 编写具体的业务逻辑
package distributed

import (
	"goweb/framework"
	"goweb/framework/contract"
	"log"
	"os"
	"syscall"
	"time"
)

type LocalDistributedService struct {
	contract.Distributed
	container framework.Container // 服务容器
}

func (l *LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (string, error) {

	lockFileName := ""
	lockFile, err := os.OpenFile(lockFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("os.OpenFile error %v", err)
		return "", err
	}

	err = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		// 没有抢占到锁
	}

	go func() {
		// 在一段时间内，选举有效，其他节点在这段时间不能再进行抢占
	}()

	// 抢占到文件锁 将抢占到的appid写入文件
	_, err = lockFile.WriteString(appID)
	if err != nil {
		log.Printf("", err)
		return "", err
	}
	return appID, nil
}
