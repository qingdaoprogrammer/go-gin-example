package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	//数据库类型
	DatabaseType string
	//数据库用户名
	DatabaseUser string
	//数据库用户密码
	DatabasePassword string
	//数据库地址
	DatabaseHost string
	//数据库名称
	DatabaseName string
	//数据库中表前缀
	TablePrefix string
)

//初始化配置信息
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()
}

//读取基础配置信息
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//读取服务器配置信息
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server: %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8009)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

//读取App安全配置信息
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	DatabaseType = sec.Key("TYPE").MustString("")
	DatabaseHost = sec.Key("HOST").MustString("")
	DatabaseUser = sec.Key("USER").MustString("")
	DatabasePassword = sec.Key("PASSWORD").MustString("")
	DatabaseName = sec.Key("NAME").MustString("")
	TablePrefix = sec.Key("TABLE_PREFIX").MustString("")
}
