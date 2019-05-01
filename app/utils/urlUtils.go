package utils

import (
	"strings"
)

func GetParam(url,keywords string) string {
	//url : /user/login?username=jiangshipan&password=19980502
	questIndex := strings.Index(url, "?")
	//打散成数组
	rs := []rune(url)
	paramStr := ""
	if questIndex != -1 {
		//判断url长度
		paramStr = string(rs[questIndex+1 : len(url)])
		//参数数组
		paramterArray := strings.Split(paramStr, "&")
		//生成参数字典
		for i := 0; i < len(paramterArray); i++ {
			str := paramterArray[i]
			if len(str) > 0 {
				temp := strings.Split(str, "=")
				if temp[0] == keywords {
					return temp[1]
				}
			}
		}
	}
	//找不到
	return ""
}