package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitee.com/go-course/restful-api-demo-g7/apps"
	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// HttpService构造函数
func NewRestfulService() *RestfulService {
	// new restful router实例, 并没有加载Handler
	r := restful.DefaultContainer

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.RestfulAddr(),
		Handler:           r,
	}
	return &RestfulService{
		server: server,
		l:      zap.L().Named("HTTP Service"),
		r:      r,
	}
}

type RestfulService struct {
	server *http.Server
	l      logger.Logger
	r      *restful.Container
}

func (s *RestfulService) Start() error {
	// 加载Handler, 把所有的模块的Handler注册给了restfuls Router
	apps.InitRestful(s.r)

	// 已加载App的日志信息
	apps := apps.LoadedRestApps()
	s.l.Infof("loaded rest apps :%v", apps)

	// 该操作时阻塞的, 简单端口，等待请求
	// 如果服务的正常关闭,
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service stoped success")
			return nil
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}

	return nil
}

func (s *RestfulService) Stop() {
	s.l.Info("start graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Warnf("shut down http service error, %s", err)
	}
}
