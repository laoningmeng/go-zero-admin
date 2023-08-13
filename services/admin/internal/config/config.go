package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	Nacos  struct {
		Host        string
		Port        int32
		NamespaceId string
		DataId      string
		Group       string
		Username    string
		Password    string
		LodDir      string
		CacheDir    string
		Timeout     int32
	}
}
