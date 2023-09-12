package util

import (
	"os"
	"strconv"
	"strings"
)

func SetEnvStr(env string, value string) string {
	if e := os.Getenv(env); e != "" {
		return e
	}
	return value
}

func SetEnvInt(env string, value int) int {
	if e := os.Getenv(env); e != "" {
		if res, err := strconv.Atoi(e); err != nil {
			panic("环境变量参数格式错误，无法注入！")
		} else {
			return res
		}
	}
	return value
}

func SetEnvBool(env string, value bool) bool {
	if e := os.Getenv(env); e != "" {
		if res, err := strconv.ParseBool(e); err != nil {
			panic("环境变量参数格式错误，无法注入！")
		} else {
			return res
		}
	}
	return value
}

func SetEnvSlice(env string, value []string) []string {
	var arr []string
	if e := os.Getenv(env); e != "" {
		arr = strings.Split(e, ",")
	}
	return arr
}
