package conf

import "go.uber.org/zap"

// 基础配置
const (
	//APIServer前缀
	APIPrefix = "/awecloud/bmq/api"
	//APIServer的端口设置
	APIPort = 8080
)

var (
	Options     *Config
	Logger      *zap.Logger
	LoggerLevel string
	RunMode     string
)

// 公共配置
type Config struct {
	// ECTD
	EtcdEndPoint    []string
	EtcdEndUsername string
	EtcdEndPassword string
	EtcdTls         bool
	EtcdCert        string
	EtcdCertKey     string
	EtcdCa          string
}
