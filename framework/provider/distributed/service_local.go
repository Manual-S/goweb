// Package distributed 编写具体的业务逻辑
package distributed

import (
	"errors"
	"goweb/framework"
	"goweb/framework/contract"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

type LocalDistributedService struct {
	contract.Distributed
	container framework.Container // 服务容器
}

func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("params error")
	}

	container := params[0].(framework.Container)

	return &LocalDistributedService{container: container}, nil
}

func (l *LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (string, error) {

	disServer := l.container.MustMake(contract.DirectoryKey).(contract.DirectoryInf)
	runtimeFolder := disServer.RuntimeFolder()
	lockFileName := filepath.Join(runtimeFolder, "distribute_"+serviceName)
	lockFile, err := os.OpenFile(lockFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("os.OpenFile error %v", err)
		return "", err
	}

	err = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		// 没有抢占到锁
		selectAppID, err := ioutil.ReadAll(lockFile)
		if err != nil {
			log.Printf("ioutil.ReadAll error %v", err)
			return "", err
		}
		return string(selectAppID), nil
	}

	go func() {
		// 在一段时间内，选举有效，其他节点在这段时间不能再进行抢占

		defer func() {
			// 释放文件锁
			syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)

			lockFile.Close()

			os.Remove(lockFileName)
		}()

		time := time.NewTimer(holdTime)
		<-time.C
	}()

	// 抢占到文件锁 将抢占到的appid写入文件
	_, err = lockFile.WriteString(appID)
	if err != nil {
		log.Printf("", err)
		return "", err
	}
	return appID, nil
}
