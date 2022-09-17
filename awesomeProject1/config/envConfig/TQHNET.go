package envConfig

var TQHNET config

func init() {
	print("加载tqhnet")
	//TQHNET.DevStatus.ACCON_ONLINE =1
	// path
	TQHNET.Path.Log = "/Users/macpro/Documents/xiaojing/go-test/"

	// 日志等级
	TQHNET.Log.Level = "trace"
	TQHNET.Log.OutStd = true
	TQHNET.Log.OutFile = true
	TQHNET.Log.ShowFileLine = true
}
