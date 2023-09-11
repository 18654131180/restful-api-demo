# Book管理系统样例

微服务开发的时候, A, B, C

都需要做分页, 都需要更新模式, 把这些公共功能，抽象公共库

## 公共库

import 把需要功能导入进行

## 功能的protobuf定义

import protobuf 公共库的时候, 需要版本对应上

比如我们公开库 demo1 v1, 生成的代码就是 v1, 依赖的v1, 我们依赖的protobuf 也是v1

这些库的位置: Page公共protobuf:  ${GOMODCACHE}/github.com/infraboard/mcube@version  pb/page
```proto
syntax = "proto3";

package page;
option go_package = "github.com/infraboard/mcube/pb/page";

message PageRequest {
    uint64 page_size = 1;
    uint64 page_number = 2;
    int64 offset = 3;
}
```

这些库的位置: Request protobuf: UpdateMode  ${GOMODCACHE}/github.com/infraboard/mcube@version pb/request
```proto
syntax = "proto3";

package mcube.request;
option go_package = "github.com/infraboard/mcube/pb/request";

enum UpdateMode {
    PUT = 0;
    PATCH = 1;
}
```

## 第一解决方案

把依赖的protobuf的对应的版本存放本地的 /usr/local/include 下面，通过-I=/usr/local/include 指定外部依赖库的位置

类似于你把这个公共的protobuf 作为了一个全局的版本, 如果你本地有2个项目，依赖2个版本, 这个时候 放在全局就无法解决

## 把依赖放入项目里面

把依赖的protobuf copy到项目内部, 然后通过 -I=common/pb, 不同项目就可以依赖指定的版本, 互补干扰

项目A: mcube@v1.1
项目B: mcube@v1.2

有没有这个的工具, 能把对应版本的外部库的protobuf定义 copy项目里面?

## protobuf 外部依赖copy

项目依赖的protobuf的版本是 v1.8.6 (go.mod), 需要copy就是 v1.8.6 protobuf

1. 找到项目依赖的外部公共库的版本
```sh
go list -m github.com/infraboard/mcube
```

2. 找到该库在本地位置: /e/Golang/pkg/mod /github.com/infraboard/mcube @v1.4.5
+ go mod 存在的位置
+ 包名称
+ 版本号

```proto
import "github.com/infraboard/mcube/pb/request/request.proto";
```

3. 确定目标copy位置:
copy 的目的地址: common/pb/ github.com/infraboard/mcube/pb
+ 搜索位置: -I=common/pb
+ 包前缀: github.com/infraboard/mcube/pb

4. copy对应的版本的protobuf文件到当前目录
```sh
cp -r /e/Golang/pkg/mod/github.com/infraboard/mcube\@v1.8.6/pb/* common/pb/github.com/infraboard/mcube/pb
```

5. 清理多余的Go文件
```
rm -rf common/pb/github.com/infraboard/mcube/pb/*/*.go
```

没有这样的工具, 把上面的逻辑做成Make:
```sh
MCUBE_MODULE := "github.com/infraboard/mcube"
MCUBE_VERSION :=$(shell go list -m ${MCUBE_MODULE} | cut -d' ' -f2)
MCUBE_PKG_PATH := ${MOD_DIR}/${MCUBE_MODULE}@${MCUBE_VERSION}

pb: ## Copy mcube protobuf files to common/pb
	@mkdir -pv common/pb/github.com/infraboard/mcube/pb
	@cp -r ${MCUBE_PKG_PATH}/pb/* common/pb/github.com/infraboard/mcube/pb
	@sudo rm -rf common/pb/github.com/infraboard/mcube/pb/*/*.go
```

6. 使用这些依赖protobuf来生成代码:
```sh
protoc -I=common/pb
```

## 为protobuf编译添加脚手架


