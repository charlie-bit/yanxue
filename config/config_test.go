package config_test

import (
	"fmt"
	"testing"

	"github.com/charlie-bit/yanxue/config"
)

func TestNewConfig2(t *testing.T) {
	fmt.Println(config.NewConfig("config/setting.yml"))
}
