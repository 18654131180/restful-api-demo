package apps

import (
	"fmt"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// IOC 容器层: 管理所有的服务的实例

// 1. HostService的实例必须注册过来, HostService才会有具体的实例, 服务启动时注册
// 2. HTTP 暴露模块, 依赖Ioc里面的HostService

var (
	// 40 Service, 怎么办? 写40个定义,
	// 使用Interface{} + 断言进行抽象
	HostService host.Service

	// 维护当前所有的服务
	implApps    = map[string]ImplService{}
	ginApps     = map[string]GinService{}
	restfulApps = map[string]RestfulService{}
	grpcApps    = map[string]GrpcService{}
)

func RegistryImpl(svc ImplService) {
	// 服务实例注册到svcs map当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	implApps[svc.Name()] = svc
	// 更加对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 如果指定了具体类型, 就导致没增加一种类型, 多一个Get方法
// func GetHostImpl(name string) host.Service

// Get 一个Impl服务的实例：implApps
// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言
func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}

	return nil
}

func RegistryGrpc(svc GrpcService) {
	// 服务实例注册到svcs map当中
	if _, ok := grpcApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	grpcApps[svc.Name()] = svc
}

// Get 一个Impl服务的实例：implApps
// 返回一个对象, 任何类型都可以, 使用时, 由使用方进行断言
func GetGrpcApp(name string) interface{} {
	for k, v := range grpcApps {
		if k == name {
			return v
		}
	}

	return nil
}

func RegistryGin(svc GinService) {
	// 服务实例注册到svcs map当中
	if _, ok := ginApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	ginApps[svc.Name()] = svc
}

func RegistryRestful(svc RestfulService) {
	// 服务实例注册到svcs map当中
	if _, ok := restfulApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registried", svc.Name()))
	}

	restfulApps[svc.Name()] = svc
}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitImpl() {
	for _, v := range grpcApps {
		v.Config()
	}

	for _, v := range implApps {
		v.Config()
	}
}

// 已经加载完成的Gin App由哪些
func LoadedGinApps() (names []string) {
	for k := range ginApps {
		names = append(names, k)
	}

	return
}

// 已经加载完成的Gin App由哪些
func LoadedRestApps() (names []string) {
	for k := range restfulApps {
		names = append(names, k)
	}

	return
}

// 已经加载完成的Gin App由哪些
func LoadedGrpcApps() (names []string) {
	for k := range grpcApps {
		names = append(names, k)
	}

	return
}

// 初始化所有的Grpc服务, 把所有实现Grpc接口的实例类，注册到grpcServer里面
func InitGrpc(r *grpc.Server) {
	// 先初始化好所有对象
	for _, v := range grpcApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range grpcApps {
		v.Registry(r)
	}
}

// 用户初始化 注册到Ioc容器里面的所有服务
func InitGin(r gin.IRouter) {
	// 先初始化好所有对象
	for _, v := range ginApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range ginApps {
		v.Registry(r)
	}
}

// 注册 restful web service
// restful 有一个Container, 类适于一个Root Router
func InitRestful(r *restful.Container) {
	// 先初始化好所有对象
	for _, v := range restfulApps {
		v.Config()
	}

	// 完成Http Handler的注册
	for _, v := range restfulApps {
		ws := new(restful.WebService)
		r.Add(ws)
		v.Registry(ws)
	}
}

type ImplService interface {
	Config()
	Name() string
}

// 注册Gin编写的Handler
// 比如 编写了Http服务A, 只需要实现Registry方法, 就能把Handler注册给Root Router
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}

type GrpcService interface {
	Registry(r *grpc.Server)
	Config()
	Name() string
}

type RestfulService interface {
	Registry(ws *restful.WebService)
	Config()
	Name() string
}
