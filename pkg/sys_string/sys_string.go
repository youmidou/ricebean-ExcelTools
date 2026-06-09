package sys_string

import (
	"strconv"
	"strings"
)

func Split(str string, sep string) []string {
	// 使用空格作为分隔符分割字符串
	result := strings.Split(str, sep)
	return result
}

// "300,500,1000,2000"
func SplitToInt32(str string, sep string) ([]int32, error) {
	// 使用空格作为分隔符分割字符串
	result := strings.Split(str, sep)
	// 创建一个整数数组
	var intArray []int32
	// 使用循环将字符串转换为整数并添加到整数数组中
	for _, s := range result {
		if s == "" {
			continue
		}
		// 将字符串转换为整数
		num, _ := strconv.ParseInt(s, 10, 32)
		// 将整数添加到整数数组
		intArray = append(intArray, int32(num))
	}

	return intArray, nil
}

func StringToInt(strNum string) (int, error) {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return 0, err
	}
	return num, err

}
