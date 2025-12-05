package main

import (
	"fmt"
	"server/config"
	"server/model/system"
	"server/plugin/SystemInit"
	"server/plugin/db"
	"server/plugin/log"
	"server/router"
	"time"
)

var cfg = config.LoadConfig()

func init() {
	// 初始化日志系统
	err := log.InitLogger(log.INFO, "./logs/app.log")
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	log.Info("Logger initialized successfully")

	// 执行初始化前等待20s , 让mysql服务完成初始化指令
	time.Sleep(time.Second * 10)
	log.Info("Starting database connections...")
	//初始化redis客户端
	err = db.InitRedisConn()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize Redis: %v", err))
	}
	log.Info("Redis initialized successfully")
	// 初始化数据库
	err = db.InitDatabase()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to initialize database: %v", err))
	}
	log.Info("Database initialized successfully")
}

func main() {
	start()
}

func start() {
	// 启动前先执行数据库内容的初始化工作
	DefaultDataInit()
	// 开启路由监听
	r := router.SetupRouter()
	_ = r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func DefaultDataInit() {
	// 如果系统中不存在用户表则进行初始化
	if !system.ExistUserTable() {
		// 初始化数据库相关数据
		SystemInit.TableInIt()
		// 初始化网站基本配置信息
		SystemInit.BasicConfigInit()
		// 初始化轮播组件信息
		SystemInit.BannersInit()
	}
	// 初始化影视来源列表信息, 并回复恢复定时任务
	SystemInit.SpiderInit()
}
