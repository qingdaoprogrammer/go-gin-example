package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

//App配置信息
type App struct {
	//页大小
	PageSize int
	//JWT密码
	JwtSecret string
	//Runtime路径
	RuntimeRootPath string

	//上传图片前缀
	ImagePrefixUrl string
	//图片保存地址
	ImageSavePath string
	//最大图片大小MB
	ImageMaxSize int
	//图片扩展名
	ImageAllowExts []string

	//日志保存路径
	LogSavePath string
	//日志文件名称
	LogSaveName string
	//日志扩展名
	LogFileExt string
	//文件日期格式
	TimeFormat string
}

var AppSetting *App = &App{}

//服务器配置信息
type Server struct {
	//运行环境release、debug
	RunMode string
	//运行端口号
	HttpPort int
	//读取超时时间
	ReadTimeout time.Duration
	//写入超时时间
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

//数据库配置信息
type Database struct {
	//数据库类型默认mysql
	Type string
	//用户名
	User string
	//密码
	Password string
	//主机地址
	Host string
	//数据库名
	Name string
	//表前缀
	TablePrefix string
}

var DatabaseSetting = &Database{}

var (
	Cfg *ini.File

	/*RunMode string

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
	TablePrefix string*/
)

//初始化配置信息
func SetUp() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	//将MB转换成字节
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}
	//将时间单位转换成秒
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	/*LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()*/
}

//读取基础配置信息
/*func LoadBase() {
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
}*/
