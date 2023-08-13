package model

import (
	"encoding/json"
	"fmt"
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

var ProviderSet = wire.NewSet(NewNacosConf, NewDB, NewUserModel, logger2.NewZapLogger)

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

func NewNacosConf(c config.Config) *NacosConf {
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
	var serverConf NacosConf
	err = json.Unmarshal([]byte(content), &serverConf)
	if err != nil {
		panic("fail to get conf from nacos")
	}
	return &serverConf

}

// NewDB  先用go-zero 的db， 后面再用gorm吧
func NewDB(conf *NacosConf) *DB {
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
	if err != nil {
		panic("fail to connect db")
	}
	return &DB{db: db}
}
