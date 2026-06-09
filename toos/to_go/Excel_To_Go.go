/*
*ExcelToGo
added by yh @ 2023/6/25 17:35
注意: go build excel_to_cx.go
*/
package to_go

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"ricebean-ExcelTools/pkg/sys_base"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	_ "github.com/xuri/excelize/v2"
)

func NewExcelToGo() *ExcelToGo {
	t := &ExcelToGo{}
	return t
}

type ExcelToGo struct {
	F        *excelize.File
	output   string //输出生成文件路径 ./Cx_output
	FileName string //文件名字

	classBaseName     string
	classBaseInfoName string
	baseInfoData      string
	ExcelPath         string

	// 中间内容
	file_class_content string
	file_content       string
	SheetTableMap      []*ClassGoSheetTable
	excel_fileName     string
}

// 打开
func (t *ExcelToGo) OpenExcelFile(fileName string, filePath string, output string) {
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
	t.excel_fileName = fileName
	fileName, explain := sys_base.GetFileName(fileName) //strings.Replace(fileName, ".xlsx", "", -1)
	t.FileName = fileName
	t.F = f
	t.ExcelPath = filePath
	t.file_content = ""
	t.DoClassTable(f, sheetlist, explain)
	t.SaveCsharpFile()
}

// DoGoTableClass
func (t *ExcelToGo) DoClassTable(f *excelize.File, sheetlist []string, explain string) {
	t.WLine("package cfg_go")
	t.WLine("/**")
	// 获取当前时间
	//currentTime := time.Now()
	// 格式化为 "年:月:日 00:00" 的格式
	//formattedTime := currentTime.Format("2006.01.02 15:04")
	t.WLine("由 %s.xlsx %s excel文件生成 ...", t.FileName, explain)
	t.WLine("author:yh ")
	t.WLine("*/")

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
			t.DoSheetTable(f, sheetName)

		}
	}
	t.DoAllIntegrate()

}

// 所有整合
func (t *ExcelToGo) DoAllIntegrate() {
	t.WLine("type %s struct{", t.FileName)
	for _, table := range t.SheetTableMap {
		keyIdType := table.GetIdKeyType()
		t.WLine("	%s map[%s]*%s", table.SheetName, keyIdType, table.classBaseInfoName)
	}
	t.WLine("}")

	for _, table := range t.SheetTableMap {
		t.WLine(table.GetContent())
	}

	//生成json 配置文件

}
func (t *ExcelToGo) DoSheetTable(f *excelize.File, sheetName string) {
	// 读取 Sheet1 中的数据
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	q := &ClassGoSheetTable{}
	q.Init(t.FileName, sheetName)
	q.DoBaseInfo(rows)
	q.DoClass()
	t.SheetTableMap = append(t.SheetTableMap, q)
}
func (t *ExcelToGo) SaveCsharpFile() {
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

	outputFile := fmt.Sprintf("%s/%s.go", t.output, t.FileName)
	// 将配置文件写入文件中
	err := ioutil.WriteFile(outputFile, []byte(t.file_content), 0644)
	if err != nil {
		fmt.Println(fmt.Sprintf("失败 配置文件%s err=%s ", outputFile, err))
		return
	} else {
		fmt.Println("配置文件已生成", outputFile)
	}

	fmt.Println("配置文件已生成", outputFile)
	t.file_content = ""
}

func (t *ExcelToGo) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"

}

type ExcelToGoMain struct {
}

func (t *ExcelToGoMain) CheckToTime() bool {

	currentTime := time.Now()
	targetTime := time.Date(2027, 5, 2, 0, 0, 0, 0, time.UTC)

	if currentTime.After(targetTime) {
		//fmt.Println("当前时间大于指定日期")
		return true
	} else {
		//fmt.Println("当前时间小于或等于指定日期")
		return false
	}
}
func (t *ExcelToGoMain) Hang() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	//ServerClose()
	fmt.Printf("kill process exit ------- signal:[%v]", s)
	//log4.Info("kill process exit ------- signal:[%v]", s)

}
func (t *ExcelToGoMain) DeleteFiles(directoryPath string, suffix string) {
	// 使用Walk函数遍历目录
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过目录本身
		if path == directoryPath {
			return nil
		}
		// 检查文件扩展名是否为.cs（不区分大小写）
		if !info.IsDir() && strings.EqualFold(filepath.Ext(path), suffix) {
			err := os.Remove(path)
			if err != nil {
				fmt.Println("无法删除文件:", err)
			} else {
				fmt.Println("已删除文件:", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录时出错:", err)
	}
}
