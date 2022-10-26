package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	nhttp "net/http"
	"os"
	"strings"

	"github.com/charlie-bit/yanxue/pkg/constant"
)

var Cfg *Config

type Config struct {
	BaseConfig `yaml:"db"`
	Stat       `yaml:",inline"`
}

type BaseConfig struct {
	// redis config
	Redis Redis `yaml:"redis"`

	// mysql
	Mysql Mysql `yaml:"mysql"`

	Env string `yaml:"env"`
}

// NewConfig creates a Config from file.
func NewConfig(name string) (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = constant.Stag
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var number = 2
	out := make(map[string]*Config, number)
	err = yaml.Unmarshal(data, out)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config file failed  {path: %v, err: %v}", name, err.Error())
	}

	c := out[env]
	if c == nil {
		return nil, fmt.Errorf("can't found env config in file  {env: %v, path: %v}", env, name)
	}

	// 配置文件强制改变环境变量
	if c.Env == "" {
		c.Env = env
	}

	Cfg = c

	return c, nil
}

// Stat config
type Stat struct {
	Addr      string `yaml:"addr"`
	HealthURI string `yaml:"health_uri"`
}

// Ping handles pinging the endpoint and returns an error if the
// agent is in an unhealthy state.
func (s *Stat) Ping() error {
	index := strings.Index(s.Addr, ":")
	host := "localhost"
	if index >= 0 {
		host += s.Addr[index:]
	}
	resp, err := nhttp.Get("http://" + host + s.HealthURI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, resp.Body)
	if resp.StatusCode != nhttp.StatusOK {
		return fmt.Errorf("server returned non-200 status code[%v]", resp.StatusCode)
	}
	return nil
}
