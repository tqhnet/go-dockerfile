package envConfig

var Dev config

func init() {
	print("加载dev")
	Dev.Path.Log = "/app/"

	// 日志等级
	Dev.Log.Level = "trace"
	Dev.Log.OutStd = true
	Dev.Log.OutFile = true
	Dev.Log.ShowFileLine = true
}
