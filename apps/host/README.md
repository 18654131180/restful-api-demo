# Host 服务模块


## IMPL

这个模块写完后, Host Service 的具体实现, 上层业务就基于Service进行编程, 面向接口
```
http
 |
Host Service (interface impl)
 |
impl(基于MySQL实现)
```

Host Service定义 并把实现编写完成, 使用方式有多种用途:
+ 用于内部模块调用, 基于他封装更高一层的业务逻辑, 比如发布服务
+ Host Service对外暴露: http协议(暴露给用户)
+ Host Service对外暴露: Grpc(暴露给内部服务)
+ ...