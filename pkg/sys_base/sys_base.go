package sys_base

import "strings"

func GetFileName(fileName string) (string, string) {
	str := strings.Replace(fileName, ".xlsx", "", -1)
	//去除@以后字符串
	parts := strings.Split(str, "@")
	result := parts[0]
	explain := ""
	if len(parts) == 2 {
		explain = parts[1]
	}

	return result, explain
}
