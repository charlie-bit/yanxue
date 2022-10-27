package common

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/db"
	"math/rand"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
)

func SendCode(phone string) error {
	// 1. 生成一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 获取配置文件
	// 2. 调用阿里云sdk，完成发送
	client, _err := CreateClient(tea.String(config.Cfg.SMSKey), tea.String(config.Cfg.SMSSecret))
	if _err != nil {
		return _err
	}

	bCode, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		TemplateCode:  tea.String(config.Cfg.SMSTemplateCode),
		SignName:      tea.String(config.Cfg.SMSSignName),
		TemplateParam: tea.String(string(bCode)),
	}

	sendSmsResponse, _err := client.SendSms(sendSmsRequest)
	if _err != nil {
		return _err
	}

	if *sendSmsResponse.Body.Code == "OK" {
		zap.S().Infof("发送给手机号： %s 的短信验证码成功【%s】", phone, code)
		// 将验证码保存进redis数据库
		db.RedisClient.Set(context.Background(), "sms_"+phone, code, time.Minute)
		return nil
	}

	return errors.New(*sendSmsResponse.Body.Message)
}

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	smsconfig := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}

	// 访问的域名
	smsconfig.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(smsconfig)
	return _result, _err
}
