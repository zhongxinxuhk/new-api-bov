package controller

import (
	"strings"

	"github.com/zhongxinxuhk/new-api-bov/setting/system_setting"
)

func paymentReturnPath(suffix string) string {
	base := strings.TrimRight(system_setting.ServerAddress, "/")
	return base + suffix
}
