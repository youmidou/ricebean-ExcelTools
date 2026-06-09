/*
*ExcelToGo
added by yh @ 2023/6/25 17:35
注意: go build excel_to_cx.go
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"ricebean-ExcelTools/toos/to_csharp"
	"ricebean-ExcelTools/toos/to_go"
	"time"

	_ "github.com/xuri/excelize/v2"
)

func main() {
	m := &to_go.ExcelToGoMain{}
	dirPath := "./excel"
	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果不存在则创建文件夹
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}
		fmt.Println("Directory created successfully!")
	}

	files, err := os.Open(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer files.Close()

	fileNames, err := files.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	//go build excel_to_cx.go
	fmt.Printf("----------*.xlsx生成规则--------------------\n")
	fmt.Printf("--第一行 字段属性为 默认空:客户端和服务端使用;c:只有客户端用;s:只有服务端使用;\n")
	fmt.Printf("--第二行 字段属性名字 如果为空字符串这个列将不生成配置\n")
	fmt.Printf("--第三行 字段类型 如果为空 将默认为string,其它类型有 string,float\n")
	fmt.Printf("--第四行 字段介绍名称\n")
	fmt.Printf("--第五行 字段属性使用介绍\n")
	fmt.Printf("--第六行 配置第一行数据开始\n")
	fmt.Printf("added by yh @ 2023/6/25 17:35 408309839@qq.com \n")
	fmt.Printf("\n")
	if m.CheckToTime() {
		fmt.Printf("--少年有报错联系管理员...\n")
		m.Hang()
		return
	}
	//sss
	output := "./bin/ExcelConfig"
	//output := "../Hotfix/ExcelConfig"
	//output := "./bin/cfg_csharp" //./Cx_output
	m.DeleteFiles(output, ".cs")

	for _, fileName := range fileNames {
		if filepath.Ext(fileName) == ".xlsx" {
			filePath := filepath.Join(dirPath, fileName)
			t := &to_csharp.ExcelToCsharp{}
			t.Init()
			t.OpenExcelFile(fileName, filePath, output)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("----------生成所有C#配置完成---------------------\n")
	time.Sleep(2 * time.Second)
	//Hang()
}
