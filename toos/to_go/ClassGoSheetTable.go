package to_go

import (
	"fmt"
)

const (
	t_int      = "int"
	t_ints     = "int[]"
	t_int2s    = "int[][]"
	t_int64    = "int64"
	t_int64s   = "int64[]"
	t_int642s  = "int64[][]"
	t_float    = "float"
	t_floats   = "[]float"
	t_float2s  = "[][]float"
	t_string   = "string"
	t_strings  = "[]string"
	t_string2s = "[][]string"
)

type ClassGoSheetTable struct {
	FileName            string //文件名字
	SheetName           string //表名字
	file_class_content  string //
	file_init_content   string //
	file_content        string //最终str
	classBaseInfoName   string
	CtypeNameList       []string
	CtypeList           []string
	CtypeAnnotationList []string
}

func (t *ClassGoSheetTable) Init(FileName string, SheetName string) {
	t.FileName = FileName
	t.SheetName = SheetName

}

/*
*
0:规则第一行不要 备注用
1:第二行 属性名字 当为空字符串 本列不做数据导入
2：第三行 类型默认 int
3:属性说明
4:属性功能使用说明
*/
func (t *ClassGoSheetTable) DoBaseInfo(rows [][]string) {
	//基本信息
	t.classBaseInfoName = fmt.Sprintf("%s_%s_Item", t.FileName, t.SheetName)
	for x, row := range rows {
		switch x {
		case 0:
		case 1:
			t.CtypeNameList = row
		case 2: //类型
			t.CtypeList = row
		case 3: //注释
			t.CtypeAnnotationList = row
		case 4:
			//=====================================================
		default:
			break
		}
	}
}

// 获取IdKey
func (t *ClassGoSheetTable) GetIdKeyType() string {
	//返回值
	ctype := t.GetcTypeName(0)
	switch ctype {
	case "int":
		return "Int32"
	case "string":
		return "string"
	}
	return ctype
}

func (t *ClassGoSheetTable) GetcTypeName(index int) string {
	//返回名字 默认int
	ctype := ""
	if index < len(t.CtypeList) {
		ctype = t.CtypeList[index]
	}
	switch ctype {
	case t_int:
		ctype = "int32"
	case t_ints:
		ctype = "[]int32"
	case t_int2s:
		ctype = "[][]int32"
	case t_int64:
		ctype = "int64"
	case t_int64s:
		ctype = "[]int64"
	case t_int642s:
		ctype = "[][]int64"
	case t_float:
		ctype = "float64"
	case t_floats:
		ctype = "[]float64"
	case t_float2s:
		ctype = "[][]float64"
	case t_string:
		ctype = "string"
	case t_strings:
		ctype = "[]string"
	case t_string2s:
		ctype = "[][]string"
	default:
		ctype = "int32" //"string"

	}
	return ctype
}

// 默认值
func (t *ClassGoSheetTable) GetcTypeValue(index int) string {
	//返回值
	//返回名字 默认int
	ctype := ""
	if index < len(t.CtypeList) {
		ctype = t.CtypeList[index]
	}
	switch ctype {
	case t_int:
		return "0"
	case t_ints:
		return "0"
	case t_int2s:
		return "0"
	case t_int64:
		return "0"
	case t_int64s:
		return "0"
	case t_int642s:
		return "0"
	case t_float:
		return "0f"
	case t_floats:
		return "0f"
	case t_float2s:
		return "0f"
	case t_string:
		return "\"\""
	case t_strings:
		return "\"\""
	case t_string2s:
		return "\"\""
	default:
		return "\"\""

	}
	return ctype
}

// 获取注释
func (t *ClassGoSheetTable) GetcTypeAnnotation(index int) string {
	zhushi := ""
	if index < len(t.CtypeAnnotationList) {
		zhushi = t.CtypeAnnotationList[index]
	}
	return zhushi
}
func (t *ClassGoSheetTable) DoClass() {
	baseInfo_data := ""

	t.WLine("type %s struct {", t.classBaseInfoName)
	//添加参数
	for index, cname := range t.CtypeNameList {
		if cname == "" {
			continue
		}
		ctype := t.GetcTypeName(index)
		zhushi := t.GetcTypeAnnotation(index)
		t.WLine("   /* %s */", zhushi)
		t.WLine("	%s %s `json:\"%s\"`", cname, ctype, cname)
		//t.WLine("	public %s %s;//%s", ctype, cname, zhushi)

		_temp := ","
		//拿到参数串
		if index == len(t.CtypeNameList)-1 {
			_temp = ""
		}
		baseInfo_data += cname + " " + ctype + _temp
	}

	t.WLine("}")
	//--------------------------------------------------------------
	//构建函数  public FileNameInfo()
	t.WLine("func (t *%s) Init(%s) {", t.classBaseInfoName, baseInfo_data)
	//参数赋值
	for _, cname := range t.CtypeNameList {
		if cname == "" {
			continue
		}
		t.WLine("	  t.%s = %s", cname, cname)
	}
	t.WLine("	}")
	//--------------------------------------------------------------
	//克隆
	t.WLine("func (t *%s) Clone() *%s {", t.classBaseInfoName, t.classBaseInfoName)

	t.WLine("	return &%s{", t.classBaseInfoName)
	baseInfo_data = ""
	//添加参数
	for _, cname := range t.CtypeNameList {
		if cname == "" {
			continue
		}
		//zhushi := t.GetcTypeAnnotation(index)
		//t.WLine("   /** %s */", zhushi)
		t.WLine("		%s:t.%s,", cname, cname)
	}
	t.WLine("	}")
	t.WLine("}")
}

func (t *ClassGoSheetTable) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}

func (t *ClassGoSheetTable) GetContent() string {
	return t.file_content
}
