package env

import (
	"bufio"
	"bytes"
	"errors"
	"goweb/framework/contract"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type EnvService struct {
	contract.EnvInf

	folder     string            // 代表.env文件所在的目录
	envMapping map[string]string // 所有的配置信息
}

func NewEnvService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		err := errors.New("params error")
		log.Printf("params error %v", err)
		return nil, err
	}
	envMapping := make(map[string]string)

	folder := params[0].(string)

	file := path.Join(folder, ".env")

	envFile, err := os.Open(file)
	if err == nil {
		// 读取env配置文件失败也没有问题
		br := bufio.NewReader(envFile)
		defer envFile.Close()

		for {
			envLine, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			s := bytes.SplitN(envLine, []byte{'='}, 2)
			if len(s) < 2 {
				continue
			}
			envMapping[string(s[0])] = string(s[1])
		}
	}

	// 替换所有的变量
	for _, value := range os.Environ() {
		pair := strings.SplitN(value, "=", 2)
		if len(pair) < 2 {
			continue
		}
		envMapping[pair[0]] = pair[1]
	}
	envMapping["APP_ENV"] = contract.EnvDevelopment
	return &EnvService{
		folder:     folder,
		envMapping: envMapping,
	}, nil
}

// AppEnv 获取当前的环境
func (e *EnvService) AppEnv() string {
	return e.Get("APP_ENV")
}

// IsExist  判断一个环境变量是否被设置
func (e *EnvService) IsExist(string) bool {
	return false
}

// Get 获取某个环境变量 如果没有设置就返回""
func (e *EnvService) Get(key string) string {
	value, ok := e.envMapping[key]
	if ok {
		return value
	}
	return ""
}

// All 获取所有的环境变量
func (e *EnvService) All() map[string]string {
	return e.envMapping
}
