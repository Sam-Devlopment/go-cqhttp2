// Package base provides base config for go-cqhttp
package base

import (
	"flag"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Mrs4s/go-cqhttp/global/config"
)

// command flags
var (
	LittleC  string // config file
	LittleD  bool   // daemon
	LittleH  bool   // help
	LittleWD string // working directory
)

// config file flags
var (
	Debug               bool // 是否开启 debug 模式
	RemoveReplyAt       bool // 是否删除reply后的at
	ExtraReplyData      bool // 是否上报额外reply信息
	IgnoreInvalidCQCode bool // 是否忽略无效CQ码
	SplitURL            bool // 是否分割URL
	ForceFragmented     bool // 是否启用强制分片
	SkipMimeScan        bool // 是否跳过Mime扫描
	UseSSOAddress       bool // 是否使用服务器下发的新地址进行重连
	LogForceNew         bool // 是否在每次启动时强制创建全新的文件储存日志
	FastStart           bool // 是否为快速启动

	PostFormat   string                 // 上报格式 string or array
	Proxy        string                 // 存储 proxy_rewrite,用于设置代理
	PasswordHash [16]byte               // 存储QQ密码哈希供登录使用
	AccountToken []byte                 // 存储 AccountToken 供登录使用
	Reconnect    *config.Reconnect      // 重连配置
	LogAging     = time.Hour * 24 * 365 // 日志时效
)

// Parse parse flags
func Parse() {
	flag.StringVar(&LittleC, "c", config.DefaultConfigFile, "configuration filename")
	flag.BoolVar(&LittleD, "d", false, "running as a daemon")
	flag.BoolVar(&LittleH, "h", false, "this help")
	flag.StringVar(&LittleWD, "w", "", "cover the working directory")
	d := flag.Bool("D", false, "debug mode")
	flag.Parse()

	config.DefaultConfigFile = LittleC // cover config file
	if *d {
		Debug = true
	}
}

// Init read config from yml file
func Init() {
	conf := config.Get()
	{ // bool config
		if conf.Output.Debug {
			Debug = true
		}
		IgnoreInvalidCQCode = conf.Message.IgnoreInvalidCQCode
		SplitURL = conf.Message.FixURL
		RemoveReplyAt = conf.Message.RemoveReplyAt
		ExtraReplyData = conf.Message.ExtraReplyData
		ForceFragmented = conf.Message.ForceFragment
		SkipMimeScan = conf.Message.SkipMimeScan
		UseSSOAddress = conf.Account.UseSSOAddress
	}
	{ // others
		Proxy = conf.Message.ProxyRewrite
		Reconnect = conf.Account.ReLogin
		if conf.Message.PostFormat != "string" && conf.Message.PostFormat != "array" {
			log.Warnf("post-format 配置错误, 将自动使用 string")
			PostFormat = "string"
		} else {
			PostFormat = conf.Message.PostFormat
		}
		if conf.Output.LogAging > 0 {
			LogAging = time.Hour * 24 * time.Duration(conf.Output.LogAging)
		}
	}
}