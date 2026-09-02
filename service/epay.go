package service

import (
	"github.com/zhongxinxuhk/new-api-bov/setting/operation_setting"
	"github.com/zhongxinxuhk/new-api-bov/setting/system_setting"
)

func GetCallbackAddress() string {
	if operation_setting.CustomCallbackAddress == "" {
		return system_setting.ServerAddress
	}
	return operation_setting.CustomCallbackAddress
}
