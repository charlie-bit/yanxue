package model_test

import (
	"fmt"
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/log"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/uhttp"
	"os"
	"testing"
	"time"
)

func TestPhoneCode_Insert(t *testing.T) {
	cfg, err := config.NewConfig("../config/setting.yml")
	if err != nil {
		fmt.Printf("load config failed err: %v", err)
		os.Exit(1)
		return
	}

	if err = db.InitMysql(cfg.Env); err != nil {
		log.Error(err.Error())
		return
	}

	req := uhttp.GetHTTPClient()

	var body struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		ZpData  []struct {
			Type  string `json:"type"`
			Array []struct {
				Country string `json:"country"`
				Code    string `json:"code"`
			} `json:"array"`
		} `json:"zpData"`
	}

	_, err = req.Get("https://signup.zhipin.com/wapi/zpuser/countryCode", nil, nil, &body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, datum := range body.ZpData {
		for _, val := range datum.Array {
			var code = &model.PhoneCode{}
			code.CreatedAt = time.Now()
			code.Country = val.Country
			code.Code = val.Code
			code.Types = datum.Type

			err = code.Insert()
			if err != nil {
				return
			}
		}
	}
}
