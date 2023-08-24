package model

import (
	"encoding/json"
	"fmt"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/svc"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"log"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	logger2 "github.com/laoningmeng/go-zero-admin/common/logger"
	"github.com/laoningmeng/go-zero-admin/services/admin/internal/config"
)

var ProviderSet = wire.NewSet(NewNacosConf, NewDB, NewUserModel, NewRoleModel, NewRuleModel, logger2.NewZapLogger)

type DB struct {
	db *gorm.DB
}
type NacosConf struct {
	Mysql struct {
		Host     string `json:"host"`
		Port     int32  `json:"port"`
		DBName   string `json:"db_name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

func (n *NacosConf) ListenOnChange(c config.Config, reload func(data string)) {
	client, err := n.client(c)
	if err != nil {
		panic(err)
	}
	go func() {
		err := client.ListenConfig(vo.ConfigParam{
			DataId: c.Nacos.DataId,
			Group:  c.Nacos.Group,
			OnChange: func(namespace, group, dataId, data string) {
				reload(data)
			},
		})
		if err != nil {
			fmt.Println(err)
		}
	}()
}
func (n *NacosConf) client(c config.Config) (config_client.IConfigClient, error) {
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
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
}
func (n *NacosConf) GetConf(c config.Config) *NacosConf {
	client, err := n.client(c)
	if err != nil {
		panic(err)
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: c.Nacos.DataId,
		Group:  c.Nacos.Group,
	})
	err = json.Unmarshal([]byte(content), n)
	if err != nil {
		panic("fail to get conf from nacos")
	}
	return n
}

func NewNacosConf(c config.Config) *NacosConf {
	nacos := &NacosConf{}
	client, err := nacos.client(c)
	if err != nil {
		panic(err)
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: c.Nacos.DataId,
		Group:  c.Nacos.Group,
	})
	var serverConf NacosConf
	err = json.Unmarshal([]byte(content), &serverConf)
	if err != nil {
		panic("fail to get conf from nacos")
	}
	return &serverConf

}

// NewDB  先用go-zero 的db， 后面再用gorm吧
func NewDB(conf *NacosConf, s *svc.ServiceContext) *DB {
	c := s.Config
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         newLogger,
	})
	conf.ListenOnChange(c, func(data string) {
		confNew := conf.GetConf(c)
		dsnNew := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			confNew.Mysql.Username,
			confNew.Mysql.Password,
			confNew.Mysql.Host,
			confNew.Mysql.Port,
			confNew.Mysql.DBName)
		dbNew, err := gorm.Open(mysql.Open(dsnNew), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			Logger:         newLogger,
		})
		if err != nil {
			fmt.Println("ERR:", err)
		}
		db = dbNew

	})
	if err != nil {
		fmt.Println("fail to connect db")
	}
	return &DB{db: db}
}
