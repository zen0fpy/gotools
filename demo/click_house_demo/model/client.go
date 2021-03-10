package model

import (
	"fmt"
	"github.com/tal-tech/cds/pkg/ckgroup"
	"github.com/tal-tech/cds/pkg/ckgroup/config"
	"github.com/tal-tech/go-zero/core/conf"
)

type CkConfig struct {
	Host     string
	User     string
	Password string
	Database string
}

type DBConfig struct {
	ClickHouse CkConfig
}

var ck ckgroup.DBGroup

func InitDB(configFile string) {
	var c DBConfig

	conf.MustLoad(configFile, &c)
	dsn := fmt.Sprintf("tcp://%s?username=%s&password=%s&database=%s",
		c.ClickHouse.Host,
		c.ClickHouse.User,
		c.ClickHouse.Password,
		c.ClickHouse.Database,
	)
	ckConfig := config.Config{
		ShardGroups: []config.ShardGroupConfig{
			{ShardNode: dsn},
		}}
	ck = ckgroup.MustCKGroup(ckConfig)
}

func DB() ckgroup.DBGroup {
	return ck
}
