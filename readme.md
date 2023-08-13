
# go-zero-admin

一个简单的基于go-zero框架的练手项目

## 1. 用到的微服务组件
* consul
* nacos
* jaeger
* prometheus
* grafana
* mysql

## 2. 一些依赖

* wire： 解决项目的依赖关系，默认go-zero 采用的链式上下文，不过随着模型增加，需要频繁改动，所以采用wire 解决这个问题
* goctl： go-zero的代码生成工具，这里需要安装一下
* gorm: 虽然go-zero 也有自己的model，并不是特别好用，更贴近原生sql，但是又没有原生灵活这里采用gorm替换model


## 3. quick start

```shell
make env
```
这里有一点需要注意，测试环境consul ，我开启了acl验证，所以启动的容器的时候，需要手动生成一下secretId
进入到consul容器内部，输入：
```shell
consul acl bootstrap
```

得到如下信息，保存起来，并修改consul.Token
```shell

AccessorID:       f5d599b2-63d4-1388-5bb7-5ae16cd01237
SecretID:         6f91b479-7245-e4e5-14ea-10b58fc2f888
Description:      Bootstrap Token (Global Management)
Local:            false
Create Time:      2023-06-06 00:03:51.776925006 +0000 UTC
Policies:
   00000000-0000-0000-0000-000000000001 - global-management
```

其实consul 也提供了key-value的配置中心的功能，但是这方面没有nacos做的好，所以这里又单独nacos做了配置中心

这里需要写入一下mysql的配置到nacos
```json
{
    "mysql":{
        "host":"127.0.0.1",
        "username":"root",
        "password":"123456",
        "port":3306,
        "db_name":"go_zero_admin"
    }
}

```

如果产生新的依赖关系,需要重新生成wire_gen
```shell
make wire
```
重新生成proto
```shell
make admin
```

构建程序
```shell
make admin_build
```

执行程序
```shell
cd services/admin
./server -f etc/admin.yaml
```

## 4. 微服务组件的一些访问地址

consul: http://localhost:8500/

nacos:  http://localhost:8848/nacos 账户 nacos 密码 nacos

grafana: http://localhost:3000/ 账户 admin 密码 admin

jaeger: http://localhost:16686/search

prometheus: http://localhost:9090/