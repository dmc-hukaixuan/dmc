/*
	所有的配置文件，按照配置文件的类别导入
*/
package database

import (
	"fmt"

	"github.com/spf13/viper"
)

// 系统信息
type SystemInfo struct {
	SystemName  string
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	CompanyName string `mapstructure:"companyName" json:"companyName" yaml:"companyName"`
}

// 日志保存地址
type LogConfig struct {
	Path  string
	Level string
}

type ConfigObject struct {
	DbType string
	// gorm
	Log     LogConfig
	SysInfo SystemInfo
	MySQL   Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PgSQL   Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
}

var c ConfigObject

func init() {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&c)
}

func GetConfig() ConfigObject {
	return c
}
