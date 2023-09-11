package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"gitee.com/go-course/restful-api-demo-g7/apps"
	"gitee.com/go-course/restful-api-demo-g7/apps/book"
	"gitee.com/go-course/restful-api-demo-g7/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

// 这个是就是我们GRPC接口的实现类
// 必须要嵌套UnimplementedServiceServer
type service struct {
	db  *sql.DB
	log logger.Logger
	book.UnimplementedServiceServer
}

func (s *service) Config() {
	s.log = zap.L().Named(s.Name())
	s.db = conf.C().MySQL.GetDB()
	return
}

func (s *service) Name() string {
	return book.AppName
}

func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

func init() {
	apps.RegistryGrpc(svr)
}
