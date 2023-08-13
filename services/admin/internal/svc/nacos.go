package svc

import (
	"encoding/json"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"github.com/laoningmeng/go-zero-admin/services/admin/internal/config"
)

type ServerConf struct {
	Mysql struct {
		Host     string
		Port     int32
		DBName   string
		Username string
		Password string
	}
}

func getConfFromNacos(c config.Config) (ServerConf, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Nacos.Host, uint64(c.Nacos.Port)),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(c.Nacos.NamespaceId),
		constant.WithTimeoutMs(uint64(c.Nacos.Timeout)),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(c.Nacos.LodDir),
		constant.WithCacheDir(c.Nacos.CacheDir),
		constant.WithLogLevel("info"),
		constant.WithUsername(c.Nacos.Username),
		constant.WithPassword(c.Nacos.Password),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: c.Nacos.DataId,
		Group:  c.Nacos.Group,
	})
	var serverConf ServerConf
	err = json.Unmarshal([]byte(content), &serverConf)
	if err != nil {
		return ServerConf{}, err
	}
	return serverConf, nil
}
