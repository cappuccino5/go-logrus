package main

import "github.com/sirupsen/logrus"

func main(){
	InitConfig()
	logrus.Info("init  success")

	logrus.Debug("init  debug")

	logrus.Error("init  error")
}