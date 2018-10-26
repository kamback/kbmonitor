//自动加载配置
package config

import (
	"github.com/BurntSushi/toml"
	"fmt"
)

type config struct {
	Logs logsConf
	Ifdb InfluxdbConf
}

type logsConf struct {
	Path string `toml:"path"` //日志文件路径
}

type InfluxdbConf struct {
	Addr      string `toml:"addr"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	Database  string `toml:"database"`
	Precision string `toml:"precision"`
	Table     string `toml:"table"`
}

var Config config

func init() {
	tomlPath := "config/config.toml"
	if _, err := toml.DecodeFile(tomlPath, &Config); err != nil {
		fmt.Println(err)
		return
	}
}
