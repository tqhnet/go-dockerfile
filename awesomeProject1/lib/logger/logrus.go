package logger

import (
	"awesomeProject1/config"
	//"backend/config"
	"bytes"
	//"cutframe/config"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

// 自定义日志格式
type logFormatter struct {
	showFileLine bool // 是否显示行数
}

func (m *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s] [%s] ", timestamp, entry.Level)
	if m.showFileLine {
		// 以超链接的方式显示调用文件及其行数
		pwd, _ := os.Getwd()
		// 获取相对项目的路径
		file := strings.Replace(entry.Caller.File, pwd+"/", "", 1)
		newLog += fmt.Sprintf("\x1b[0;34m%s:%d\x1b[0m ", file, entry.Caller.Line)
	}
	newLog += fmt.Sprintf("%s\n", entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	// 日志输出位置
	var out = []io.Writer{}
	if config.Config.Log.OutStd {
		// 输出到控制台
		out = append(out, os.Stdout)
	}
	if config.Config.Log.OutFile {
		// 输出到文件
		// 日志文件按天切割
		fileName := config.Config.Path.Log + "%Y%m%d_access.log"
		fileOut, _ := rotatelogs.New(fileName)
		out = append(out, fileOut)
		fmt.Println("输出目录：" + fileName)
	}
	// 输出控制台和文件都为false时，则会输出到控制台
	if len(out) > 0 {
		mw := io.MultiWriter(out...)
		Logger.SetOutput(mw)
	}

	// 设置日志格式
	formatter := logFormatter{showFileLine: config.Config.Log.ShowFileLine}
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&formatter)

	// 设置日志等级
	switch config.Config.Log.Level {
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
		break
	case "fatal":
		Logger.SetLevel(logrus.FatalLevel)
		break
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
		break
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
		break
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
		break
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
		break
	case "trace":
		Logger.SetLevel(logrus.TraceLevel)
		break
	default:
		Logger.SetLevel(logrus.TraceLevel)
	}
}
