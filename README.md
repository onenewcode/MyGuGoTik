# 目录架构
```shell
├─docs # 文档说明
├─src # GuGoTik 源代码
│  ├─constant # 常量，用于定义服务预设信息等
│  │  ├─config # 为项目提供配置读取说明
│  │  └─strings # 提供编码后的信息，用于定义常量信息
│  ├─idl # idl 说明文件
│  ├─models # 用于存储通用数据模型
│  ├─rpc # gRPC 生成文件
│  ├─services # 服务，下为横向扩展的服务
│  │  ├─auth # Auth 鉴权 / 登录服务
│  │  └─health # *唯一一个非横向扩展服务，用于注册到其他服务中，提供 consul 健康检查的功能
│  ├─storage # 存储模块，暂时缺少 RabbitMQ 对接模块，需要由视频相关业务开发组制作
│  │  ├─database # 数据库模块，对接 PostgreSQL
│  │  ├─file # 二进制存储模块，目前只有 fs 模块
│  │  └─redis # Redis 模块，对接 Redis
│  ├─utils # 通用问题
│  │  ├─consul # Consul 服务，用于向 Consul 注册服务
│  │  ├─interceptor # 拦截器，用于切片某一个方法或过程
│  │  ├─logging # 日志
│  │  └─trace # 链路追踪
│  └─web # 网页服务
│      ├─about # *About 服务，非正式业务，仅供测试
│      ├─auth # Auth 服务，提供 /douyin/user * 接口
│      ├─authmw # Auth 鉴权中间件，非服务
│      ├─middleware # Middle Ware 中间件，为除了 Auth MW 以外的中间件服务
│      └─models # 网站模型
└─test # 单元测试
    ├─rpc # GRPC 单元测试
    └─web # 网页单元测试
```
# 项目搭建
```shell
├─.idea
├─config
└─src
    ├─constant
    │  └─config
    ├─models
    ├─storage
    │  └─database
    └─utils
        └─logging

```
在idl文件下执行以下命令
protoc --proto_path=. --go_out=./../../.. --go-grpc_out=./../../.. ./*.proto


```shell
├─config
└─src
    ├─constant
    │  ├─config
    │  └─strings
    ├─idl
    ├─models
    ├─rpc
    │  ├─auth
    │  ├─chat
    │  ├─comment
    │  ├─favorite
    │  ├─feed
    │  ├─health
    │  ├─publish
    │  ├─recommend
    │  ├─relation
    │  └─user
    ├─storage
    │  └─database
    └─utils
        └─logging

```

# 存在的问题
## consul
而consul是注册中心访问服务提供者健康检查url。

**问题在于**：服务器无法主动与内网IP建立连接(连路由都做不到)，也就是说除非你本地主机拥有公网IP, 否则无法直接ping通。

在windos安装，
输入consul --version进行测试，最后输入consul agent -dev,访问http://localhost:8500/