package protocol

import (
	"net"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"gitee.com/go-course/restful-api-demo-g7/apps"
	"gitee.com/go-course/restful-api-demo-g7/conf"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	log := zap.L().Named("GRPC Service")

	// rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer()

	return &GRPCService{
		svr: grpcServer,
		l:   log,
		c:   conf.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	apps.InitGrpc(s.svr)

	// 已加载App的日志信息
	apps := apps.LoadedGrpcApps()
	s.l.Infof("loaded grpc apps :%v", apps)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GrpcAddr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GrpcAddr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	s.svr.GracefulStop()
	return nil
}
