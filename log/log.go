package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

var (
	Logger  *zap.SugaredLogger
	ZLogger *zap.Logger
)

// Init log
func Init() {
	var (
		err error
		cfg zap.Config
	)

	cfg = zap.NewDevelopmentConfig()
	//cfg = zap.NewProductionConfig()

	cfg.OutputPaths[0] = "stdout"

	ZLogger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	Logger = ZLogger.Sugar()
}

func Exit() {
	if Logger != nil {
		Logger.Sync()
	}
	if ZLogger != nil {
		ZLogger.Sync()
	}
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}
func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}
func Info(args ...interface{}) {
	Logger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}
func Error(args ...interface{}) {
	Logger.Error(args...)
	Ding("", args)
}
func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
	Ding(template, args)
}
func DPanic(args ...interface{}) {
	Logger.DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	Logger.DPanicf(template, args...)
}
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	Logger.Panicf(template, args...)
}
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}

type DingMarkdown struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Context string `json:"content"`
	} `json:"text"`
}

func Ding(template string, args ...interface{}) {
	var d DingMarkdown
	d.Msgtype = "text"
	d.Text.Context = fmt.Sprintf("err : "+template, args)

	buf, _ := json.Marshal(d)

	_, err := http.DefaultClient.Post("https://oapi.dingtalk.com/robot/send?access_token=3325c63def69133725f88bafc172a60b2e5a1a0e8ae660d1bc446c8c559acb69", "application/json", bytes.NewBuffer(buf))
	fmt.Println(err)
}
