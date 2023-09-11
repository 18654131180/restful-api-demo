package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"

	"gitee.com/go-course/restful-api-demo-g7/apps"
	"gitee.com/go-course/restful-api-demo-g7/conf"

	"gitee.com/go-course/restful-api-demo-g7/protocol"

	// 注册所有的服务实例
	_ "gitee.com/go-course/restful-api-demo-g7/apps/all"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

// 程序的启动时 组装都在这里进行
// 1.
// StartCmd represents the base command when called without any subcommands
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}

		// 初始化全局日志Logger
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 加载我们Host Service的实体类
		// host service 的具体实现
		// service := impl.NewHostServiceImpl()

		// 注册HostService的实例到IOC
		// 采用: _ "gitee.com/go-course/restful-api-demo-g7/apps/host/impl" 完成注册
		// apps.HostService = impl.NewHostServiceImpl()

		// 如何执行HostService的config方法
		// 因为 apps.HostService是一个host.Service接口, 并没有 保存实例初始化(Config)的方法
		apps.InitImpl()

		// // 提供一个Gin Router
		// g := gin.Default()
		// // 注册IOC的所有http handler
		// apps.InitGin(g)
		// g.Run(conf.C().App.HttpAddr())
		svc := newManager()

		ch := make(chan os.Signal, 1)
		// channel是一种复合数据结构, 可以当初一个容器, 自定义的struct make(chan int, 1000), 8bytes * 1024  1Kb
		// 如果没close gc是不会回收的
		defer close(ch)

		// Go为了并发编程设计的(CSP), 依赖Channel作为数据通信的信道
		// 出现了一个思路模式的转变:
		//    单兵作战(只有一个Groutine) --> 团队作战(多个Groutine 采用Channel来通信)
		//    main { for range channel }  这个时候的channel仅仅想到于一个缓存, 必须选择带缓存区的channl
		//    signal.Notify 当中一个Goroutine, g1
		//    go svc.WaitStop(ch) 第二Goroutine, g2
		//    g1 -- ch1 --> g2
		//    g1 <-- ch2 -- g2
		//    g1 数据发送给ch1, g2 读取channle中的数据, chanel 只要生成好了就能用, 如果channle关闭
		//    设计channel 使用数据的发送方负责关闭, 相当于表示挂电话
		//    for range   由range帮忙处理了 chnanel 关闭后， read的中断处理
		//    for v,err := <-ch { if(err == io.EOF) break }

		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go svc.WaitStop(ch)

		//  gprc 服务启动在后台
		go svc.grpc.Start()

		// restful api 启动在后头
		go svc.rest.Start()

		return svc.Start()
	},
}

// 有2个服务, 一个http, 一个gprc
//
func newManager() *manager {
	return &manager{
		rest: protocol.NewRestfulService(),
		http: protocol.NewHttpService(),
		grpc: protocol.NewGRPCService(),
		l:    zap.L().Named("CLI"),
	}
}

// 用于管理所有需要启动的服务
// 1. HTTP服务的启动
type manager struct {
	rest *protocol.RestfulService
	http *protocol.HttpService
	grpc *protocol.GRPCService
	l    logger.Logger
}

func (m *manager) Start() error {
	return m.http.Start()
}

// 处理来自外部的中断信号, 比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Infof("received signal: %s", v)

			// 先关闭内部调用
			if err := m.grpc.Stop(); err != nil {
				m.l.Error(err)
			}

			// 在关闭外部调用
			m.http.Stop()
		}
	}
}

// 问题:
// 1. http API, Grpc API 需要启动, 消息总线也需要监听, 比如负责注册于配置,  这些模块都是独立
//    都需要在程序启动时，进行启动, 都写在start start膨胀到不易维护
// 2. 服务的优雅关闭怎么办? 外部都会发送一个Terminal(中断)信号给程序, 程序时需要处理这个信号
//    需要实现程序优雅关闭的逻辑的处理: 由先后顺序 (从外到内完成资源的释放逻辑处理)
//    1. api 层的关闭 (HTTP, GRPC)
//    2. 消息总线关闭
//    3. 关闭数据库链接
//    4. 如果 使用了注册中心, 完成注册中心的注销操作
//    5. 退出完毕

// 还没有初始化Logger实例
// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)

	// 更加Config里面的日志配置，来配置全局Logger对象
	lc := conf.C().Log

	// 解析日志Level配置
	// DebugLevel: "debug",
	// InfoLevel:  "info",
	// WarnLevel:  "warning",
	// ErrorLevel: "error",
	// FatalLevel: "fatal",
	// PanicLevel: "panic",
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}

	// 使用默认配置初始化Logger的全局配置
	zapConfig := zap.DefaultConfig()

	// 配置日志的Level基本
	zapConfig.Level = level

	// 程序没启动一次, 不必都生成一个新日志文件
	zapConfig.Files.RotateOnStartup = false

	// 配置日志的输出方式
	switch lc.To {
	case conf.ToStdout:
		// 把日志打印到标准输出
		zapConfig.ToStderr = true
		// 并没在把日志输入输出到文件
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}

	// 配置日志的输出格式:
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}

	// 把配置应用到全局Logger
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
