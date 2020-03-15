package main

import (
	"flag"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	"bbs-go/app"
	"bbs-go/common/config"
	"bbs-go/model"
)

var configFile = flag.String("config", "./wkycloud.yaml", "配置文件路径")

func init() {
	flag.Parse() //解析命令行参数

	config.InitConfig(*configFile) // 初始化配置
	initLogrus()                   // 初始化日志
	//simple是一个封装了一些功能的类库，类似与公共工具API。TODO 这个simple库以后可以换成自己的，方便追加新的功能
	err := simple.OpenMySql(config.Conf.MySqlUrl, 10, 20, config.Conf.ShowSql, model.Models...) // 连接数据库
	if err != nil {
		logrus.Error(err)
	}
}

//logrus是一个结构化的日志库，完全兼容标准的日志
func initLogrus() {
	file, err := os.OpenFile(config.Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}
}

func main() {
	app.StartOn()
	app.InitIris()
}
