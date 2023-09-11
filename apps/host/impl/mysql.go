package impl

import (
	"database/sql"

	"gitee.com/go-course/restful-api-demo-g7/apps"
	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
// 这样写, 会造成 conf.C()并准备好, 造成conf.C().MySQL.GetDB()该方法pannic
// var impl = NewHostServiceImpl()

// 把对象的注册和对象的注册这2个逻辑独立出来
var impl = &HostServiceImpl{}

// NewHostServiceImpl 保证调用该函数之前, 全局conf对象已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子Loggger
		// 封装的Zap让其满足 Logger接口
		// 为什么要封装:
		// 		1. Logger全局实例
		// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
		// 		3. 加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

// 只需要保证 全局对象Config和全局Logger已经加载完成
func (i *HostServiceImpl) Config() {
	// Host service 服务的子Loggger
	// 封装的Zap让其满足 Logger接口
	// 为什么要封装:
	// 		1. Logger全局实例
	// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
	// 		3. 加入日志轮转功能的集合
	i.l = zap.L().Named("Host")
	i.db = conf.C().MySQL.GetDB()
}

// 服务服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

// _ import app 自动执行注册逻辑
func init() {
	//  对象注册到ioc层
	apps.RegistryImpl(impl)
}

// 注册HostService的实例到IOC
// apps.HostService = impl.NewHostServiceImpl()

// mysql 的驱动加载的实现方式
// sql 这个库, 是一个框架, 驱动是 引入依赖的时候加载的
// 我们把 app模块，比作一个驱动, ioc比作框架
// _ import app， 该app就注册到ioc层
