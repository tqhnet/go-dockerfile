package envConfig

// mongodb 配置
type mongo struct {
	Conn            string
	Database        string
	CtxTimeout      uint   // ctx超时时间
	LocalThreshold  uint   // 与mongo操作耗时小于3秒
	MaxConnIdleTime uint   // 连接最大空闲时间
	MaxPoolSize     uint64 // 连接池最大数量
}

type qiniuKodo struct {
	Bucket    string
	AccessKey string
	SecretKey string
	BaseUrl   string
	Zone      string
}

type config struct {
	// rest 服务相关配置
	baseConfig
	projectConfig
}

// 基础配置，连接环境相关
type baseConfig struct {
	//Rest struct {
	//	Domain string
	//	Host   string
	//	Port   int
	//}
	// 目录
	Path struct {
		Log string
		Tmp string
	}
	// 日志
	Log struct {
		Level        string // 日志等级
		OutStd       bool   // 输出到控制台
		OutFile      bool   // 输出到文件
		ShowFileLine bool   // 打印函数所在文件及其行数
	}
	//// mongodb 配置
	//Mongo struct {
	//	Backend mongo // cutframe库
	//}
	//// redis 配置
	//Redis struct {
	//	Host        string
	//	Password    string
	//	MaxIdle     int
	//	MaxActive   int
	//	IdleTimeout int
	//}
	//// rabbitmq 配置
	//Rabbitmq struct {
	//	Host string
	//}
	// etcd相关配置
	//Etcd struct {
	//	Addr     string
	//	Username string
	//	Password string
	//	BasePath string
	//}
	//// 七牛云
	//Qiniu struct {
	//	Backend qiniuKodo
	//}
	//// 极光推送
	//Jpush struct {
	//	AppKey       string
	//	MasterSecret string
	//}
}

// 项目配置
type projectConfig struct {
	BlackList []string
}

var Default config

func init() {
	baseConfigInit()
	projectConfigInit()
}

func baseConfigInit() {
	// path
	Default.Path.Log = "/db/tmp/test/log/"
	Default.Path.Tmp = "/tmp/tmp"

	// 日志等级
	Default.Log.Level = "trace"
	Default.Log.OutStd = false
	Default.Log.OutFile = false
	Default.Log.ShowFileLine = false

}

func projectConfigInit() {
	Default.BlackList = []string{"yika"}
}
