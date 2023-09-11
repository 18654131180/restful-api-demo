package impl

import (
	"database/sql"
	"github.com/18654131180/restful-api-demo/apps/host"
	"github.com/18654131180/restful-api-demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
// 这样写, 会造成 conf.C()并准备好, 造成conf.C().MySQL.GetDB()该方法pannic
// var impl = NewHostServiceImpl()

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

// 服务的名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}
