package main

import (

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

//本包仅用于main函数初始化logrus和日志轮转
//main函数调用InitLog初始化完成后，其他包无需import本包，其他包只需要import logrus即可
var logConf lumberjack.Logger //oplog 配置
/*
   "log":{
       "filename":"./ggm.log",
       "maxsize":10,
       "maxage":1,
       "maxbackups":3,
       "compress":"true",
       "localtime":"true"
   }

	logConf = lumberjack.Logger{
		Filename:   v.GetString("log.filename"),
		MaxSize:    v.GetInt("log.maxsize"),
		MaxAge:     v.GetInt("log.maxage"),
		MaxBackups: v.GetInt("log.maxbackups"),
		LocalTime:  v.GetBool("log.localtime"),
		Compress:   v.GetBool("log.compress"),
	}
*/
func init() {
	logConf = lumberjack.Logger{
		MaxSize:    10,
		MaxAge:     1,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   true,
	}
}
func InitConfig() {
	fmtr := new(logrus.TextFormatter)
	fmtr.FullTimestamp = true                    // 显示完整时间
	fmtr.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	fmtr.DisableTimestamp = false                // 禁止显示时间
	fmtr.DisableColors = false                   // 禁止颜色显示

	hook := NewHook()
	hook.Field = "line"
	logrus.SetFormatter(fmtr)
	logrus.AddHook(hook)
	logrus.SetLevel(logrus.DebugLevel)
		jack := &logConf //日志路径，轮转，压缩等
		jack.Filename = "./def.log"
		logrus.SetOutput(jack)

}
