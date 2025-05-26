package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
)

func InitLogger() {
	encode := getEncoder()
	loggingInfo := getLogWriterInfo()
	logLevel := zapcore.DebugLevel
	coreInfo := zapcore.NewCore(encode, loggingInfo, logLevel)
	logger := zap.New(coreInfo)
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	//生成一个默认的配置文件，然后再由这个配置文件，生成一个JSON记录器
	produnctionConfig := zap.NewProductionEncoderConfig()
	produnctionConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	produnctionConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(produnctionConfig)
}

func getLogWriterInfo() zapcore.WriteSyncer {
	logpath, _ := os.Getwd()
	logpath = strings.TrimSuffix(logpath, "/internal/common") + "/log/log.log"
	//lumberjack是一个日志滚动记录器。写入lumberjack的日志达到一定的条件后会进行存档（普通文件的形式，或压缩文件的形式），然后新建另一个同名文件（原文件存档时会重命名）继续记录。但是lumberjack本身并不包含日志的基础功能，
	//例如日志等级、日志格式化等。理论上可以向lumberjack写入任意文本，并实现滚动记录。一般情况下，lumberjack配合其他日志库，实现日志的滚动(rolling)记录。
	l := &lumberjack.Logger{
		Filename:   logpath,
		MaxSize:    18000, //文件最大容量MB
		MaxBackups: 7,     //最大备份数量
		MaxAge:     7,     //日志存活时间
		Compress:   true,  //是否对旧文件进行压缩
	}
	var ws io.Writer
	ws = io.MultiWriter(l, os.Stdout)
	return zapcore.AddSync(ws)
}
