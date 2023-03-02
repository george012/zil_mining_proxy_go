package main

import (
	"fmt"
	"github.com/george012/gtbox"
	"github.com/george012/gtbox/gtbox_log"
	"github.com/sirupsen/logrus"
	"zil_mining_proxy_go/config"
)

func main() {
	// 读取配置文件
	configFileInfo := config.ReadConfig()
	logLevel := logrus.InfoLevel

	if config.DebugMod == true {
		logLevel = logrus.DebugLevel
	}
	gtbox.SetupGTBox(config.AppNameLower, config.DebugMod, logLevel, 30, gtbox_log.GTLogSaveHours, 20)

	fmt.Printf("%s", configFileInfo)
	// 建立到Zilliqa节点的连接
	zilliqaNode := NewZilliqaNode(config.ZilliqaNodeURL)

	//// 启动网络代理
	//ln, err := net.Listen("tcp", config.ProxyListenAddress)
	//if err != nil {
	//	panic(err)
	//}
	//defer ln.Close()
	//
	//workerManager := NewWorkerManager()
	//proxy := NewNetworkProxy(ln, zilliqaNode, workerManager)
	//fmt.Printf("代理程序已启动，监听地址：%s\n", config.ProxyListenAddress)
	//
	//// 监听来自矿工的连接请求
	//for {
	//	conn, err := ln.Accept()
	//	if err != nil {
	//		fmt.Println("接受连接请求失败：", err)
	//		continue
	//	}
	//
	//	// 创建新的矿工连接
	//	workerConn := NewWorkerConnection(conn)
	//
	//	// 将矿工连接加入矿工管理器
	//	workerManager.AddWorkerConnection(workerConn)
	//}
}
