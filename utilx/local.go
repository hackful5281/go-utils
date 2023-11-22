package utilx

import (
	"log"
	"net"
	"os"
)

const (
	ENV_PRODUCTION = "production"
	ENV_DEVELOP    = "develop"
)

var (
	_project   = ""
	_namespace = ""
	_runenv    = ""
)

func init() {
	_project = os.Getenv("PORJECT_NAME")
	_namespace = os.Getenv("NAMESPACE")
	_runenv = os.Getenv("RUN_ENV")
}

// GetLocalIp 获取本地IP
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Println("获取本地IP异常")
		log.Println(err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetProjectName 获取当前服务的项目名
func GetProjectName() string {
	return _project
}

// GetRunEnv 获取当前服务运行的环境
func GetRunEnv() string {
	return _runenv
}

// GetNamespace 获取当前的服务所在的命名空间
func GetNamespace() string {
	return _namespace
}
