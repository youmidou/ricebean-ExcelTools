/*
*ExcelToGo
added by yh @ 2023/6/25 17:35
注意: go build excel_to_cx.go
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"ricebean-ExcelTools/toos/to_go"
	"ricebean-ExcelTools/toos/to_json"
	"time"
	//_ "github.com/xuri/excelize/v2"
)

func main() {
	// 定义命令行参数：参数名、默认值、参数说明
	go_output_dir := flag.String("go_output_dir", "./config/cfg_go", "Go 文件的输出目录")
	json_output_dir := flag.String("json_output_dir", "./config/cfg_json", "JSON 文件的输出目录")

	// 必须调用 Parse() 来解析传入的参数
	flag.Parse()

	// 注意：flag.String 返回的是指针，使用时需要加 * 取值
	fmt.Printf("Go 输出目录: %s\n", *go_output_dir)
	fmt.Printf("JSON 输出目录: %s\n", *json_output_dir)

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
	///
	go_output := "./config/cfg_go"     //"./bin/cfg_go"
	json_output := "./config/cfg_json" //"./bin/cfg_json"
	if 0 < len(*go_output_dir) {
		go_output = *go_output_dir
	}
	if 0 < len(*json_output_dir) {
		json_output = *json_output_dir
	}

	m.DeleteFiles(go_output, ".go")
	m.DeleteFiles(json_output, ".json")

	for _, fileName := range fileNames {
		if filepath.Ext(fileName) == ".xlsx" {
			filePath := filepath.Join(dirPath, fileName)
			t := to_go.NewExcelToGo()
			t.OpenExcelFile(fileName, filePath, go_output)

			js := to_json.NewExcelToJson()
			js.OpenExcelFile(fileName, filePath, json_output)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("----------生成所有C#配置完成---------------------\n")
	time.Sleep(2 * time.Second)
	//Hang()
}
