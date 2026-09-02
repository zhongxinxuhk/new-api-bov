package service

import (
	"strings"

	"github.com/zhongxinxuhk/new-api-bov/constant"
)

func CoverTaskActionToModelName(platform constant.TaskPlatform, action string) string {
	return strings.ToLower(string(platform)) + "_" + strings.ToLower(action)
}
