package impl_test

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"gitee.com/go-course/restful-api-demo-g7/apps/host/impl"
	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)
	ins := host.NewHost()
	ins.Id = "ins-01"
	ins.Name = "test"
	ins.Region = "cn-hangzhou"
	ins.Type = "sm1"
	ins.CPU = 1
	ins.Memory = 2048
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func TestQuery(t *testing.T) {
	should := assert.New(t)

	req := host.NewQueryHostRequest()
	req.Keywords = "接口测试"
	set, err := service.QueryHost(context.Background(), req)
	if should.NoError(err) {
		for i := range set.Items {
			fmt.Println(set.Items[i].Id)
		}
	}
}

func TestDescribe(t *testing.T) {
	should := assert.New(t)

	req := host.NewDescribeHostRequestWithId("ins-09")
	ins, err := service.DescribeHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

func TestUpdate(t *testing.T) {
	should := assert.New(t)

	req := host.NewPutUpdateHostRequest("ins-09")
	req.Name = "更新测试02"
	req.Region = "rg 02"
	req.Type = "small"
	req.CPU = 1
	req.Memory = 2048
	req.Description = "测试更新"
	ins, err := service.UpdateHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

// {
// 	"id": "ins-10",
// 	"vendor": 0,
// 	"region": "cn-hangzhou",
// 	"create_at": 1651290785519,
// 	"expire_at": 0,
// 	"type": "sm1",
// 	"name": "接口测试",
// 	"description": "",
// 	"status": "",
// 	"tags": null,
// 	"update_at": 0,
// 	"sync_at": 0,
// 	"accout": "",
// 	"public_ip": "10.1.1.2",
// 	"private_ip": "",
// 	"cpu": 1,
// 	"memory": 2048,
// 	"gpu_amount": 0,
// 	"gpu_spec": "",
// 	"os_type": "",
// 	"os_name": "",
// 	"serial_number": ""
// }
func TestPatch(t *testing.T) {
	should := assert.New(t)

	req := host.NewPatchUpdateHostRequest("ins-09")
	req.Description = "Patch更新模式测试"
	ins, err := service.UpdateHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

func init() {
	// 测试用例的配置文件
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 需要初始化全局Logger,
	// 为什么不设计为默认打印, 因为性能
	fmt.Println(zap.DevelopmentSetup())

	// host service 的具体实现
	service = impl.NewHostServiceImpl()
}
