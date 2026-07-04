/*
*ExcelToJson
added by yh @ 2023/6/25 17:35
注意: go build excel_to_cx.go
*/
package to_json

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"ricebean-ExcelTools/pkg/sys_base"
	"strings"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
)

type ExcelToJson struct {
	mutex    sync.Mutex
	F        *excelize.File
	output   string //输出生成文件路径 ./Cx_output
	FileName string //文件名字

	classBaseName     string
	classBaseInfoName string
	baseInfo_data     string
	ExcelPath         string

	// 中间内容
	file_class_content string
	file_content       string
	file2Content       string
	SheetTableMap      []*ClassJsonSheetTable
}

func NewExcelToJson() *ExcelToJson {
	t := &ExcelToJson{}
	return t
}

// 打开
func (t *ExcelToJson) OpenExcelFile(fileName string, filePath string, output string) {
	t.output = output
	//"j-奖励表.xlsx"
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetlist := f.GetSheetList() //GetSheetMap

	fileName, explain := sys_base.GetFileName(fileName) //strings.Replace(fileName, ".xlsx", "", -1)
	t.FileName = fileName
	t.F = f
	t.ExcelPath = filePath
	t.file_content = ""
	t.DoClassTable(f, sheetlist, explain)
	t.SaveCsharpFile()
}

// DoGoTableClass
func (t *ExcelToJson) DoClassTable(f *excelize.File, sheetlist []string, explain string) {
	//t.WLine("package cfg")
	//t.WLine("/**")
	//t.WLine("由 %s.xlsx excel文件生成 ...", t.FileName)
	//t.WLine("author:yh ")
	//t.WLine("*/")

	n := len(sheetlist)

	for i := 0; i < n; i++ {
		sheetName := sheetlist[i]
		if sheetName == "" {
			continue
		}
		if strings.Contains(sheetName, "说明") {
			fmt.Printf("包含说明 %s 不处理\n", sheetName)
		} else {
			fmt.Printf("  ----开始生成:%s.%s  Path= %s\n", t.FileName, sheetName, f.Path)

			t.DoSheetTable(f, sheetName, i+1 == n)

		}
	}
	t.DoAllIntegrate()

}

// 所有整合
func (t *ExcelToJson) DoAllIntegrate() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.WLine("{")
	for _, table := range t.SheetTableMap {
		t.WLine(table.GetContent())
	}
	t.WLine("}")

	//
}
func (t *ExcelToJson) DoSheetTable(f *excelize.File, sheetName string, isbool bool) {
	// 读取 Sheet1 中的数据
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	//t.SheetTableMap = make(map[string]*ClassJsonSheetTable)
	q := &ClassJsonSheetTable{}
	q.Init(t.FileName, sheetName)
	q.DoBaseInfo(rows)
	//q.DoClass()
	q.DoCfgData(rows, isbool)
	t.SheetTableMap = append(t.SheetTableMap, q)
}
func (t *ExcelToJson) SaveCsharpFile() {
	//dirPath := "./Cx_output"
	dirPath := t.output
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

	outputFile := fmt.Sprintf("%s/%s.json", t.output, t.FileName)
	// 将配置文件写入文件中
	err := ioutil.WriteFile(outputFile, []byte(t.file_content), 0644)
	if err != nil {
		fmt.Println(fmt.Sprintf("失败 配置文件%s err=%s ", outputFile, err))
		return
	}
	fmt.Println("配置文件已生成", outputFile)
	t.file_content = ""
}

func (t *ExcelToJson) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"

}

type ExcelToJsonMain struct {
}

func (t *ExcelToJsonMain) CheckToTime() bool {

	currentTime := time.Now()
	targetTime := time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC)

	if currentTime.After(targetTime) {
		//fmt.Println("当前时间大于指定日期")
		return true
	} else {
		//fmt.Println("当前时间小于或等于指定日期")
		return false
	}
}
func (t *ExcelToJsonMain) Hang() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	//ServerClose()
	fmt.Printf("kill process exit ------- signal:[%v]", s)
	//log4.Info("kill process exit ------- signal:[%v]", s)

}
