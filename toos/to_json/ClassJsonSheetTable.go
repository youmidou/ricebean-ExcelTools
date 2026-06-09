package to_json

import (
	"fmt"
	"ricebean-ExcelTools/pkg/sys_base"
	"ricebean-ExcelTools/pkg/sys_json"
	"ricebean-ExcelTools/pkg/sys_string"
	"strings"
	"sync"
)

type ClassJsonSheetTable struct {
	mutex               sync.Mutex
	FileName            string //文件名字
	SheetName           string //表名字
	file_class_content  string //
	file_init_content   string //
	file_content        string //最终str
	classBaseInfoName   string
	ctypeNameList       []string
	CtypeList           []string
	CtypeAnnotationList []string
	//base                *ExcelToGo
}

func (t *ClassJsonSheetTable) Init(FileName string, SheetName string) {
	t.FileName = FileName
	t.SheetName = SheetName
	//t.base = base
}

/*
*
0:规则第一行不要 备注用
1:第二行 属性名字 当为空字符串 本列不做数据导入
2：第三行 类型默认 int
3:属性说明
4:属性功能使用说明
*/
func (t *ClassJsonSheetTable) DoBaseInfo(rows [][]string) {
	//基本信息
	t.classBaseInfoName = fmt.Sprintf("%s_%s_Item", t.FileName, t.SheetName)
	for x, row := range rows {
		switch x {
		case 0:
		case 1:
			t.ctypeNameList = row
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

func (t *ClassJsonSheetTable) GetCType(index int) string {
	//返回名字 默认int
	ctype := ""
	if index < len(t.CtypeList) {
		ctype = t.CtypeList[index]
	}
	return ctype
}
func (t *ClassJsonSheetTable) GetCTypeName(index int) string {
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
		ctype = "string"

	}
	return ctype
}

// 默认值
func (t *ClassJsonSheetTable) GetcTypeValue(index int) string {
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
func (t *ClassJsonSheetTable) GetcTypeAnnotation(index int) string {
	zhushi := ""
	if index < len(t.CtypeAnnotationList) {
		zhushi = t.CtypeAnnotationList[index]
	}
	return zhushi
}
func (t *ClassJsonSheetTable) DoClass() {
	baseInfo_data := ""

	t.WLine("\"%s\":{", t.classBaseInfoName)
	//添加参数
	for index, cname := range t.ctypeNameList {
		if cname == "" {
			continue
		}
		ctype := t.GetCTypeName(index)
		//zhushi := t.GetcTypeAnnotation(index)
		//t.WLine("   /** %s */", zhushi)
		//t.WLine("	%s %s ", cname, ctype)
		//t.WLine2("	public %s %s;//%s", ctype, cname, zhushi)
		t.WLine("	\"%s\"={", cname)
		t.WLine("	}")

		_temp := ","
		//拿到参数串
		if index == len(t.ctypeNameList)-1 {
			_temp = ""
		}
		baseInfo_data += ctype + " " + cname + _temp
	}

	t.WLine("}")
	/*
		//构建函数  public FileNameInfo()
		t.WLine("func (t *%s) Init(%s) {", t.classBaseInfoName, baseInfo_data)
		//参数赋值
		for _, cname := range t.ctypeNameList {
			if cname == "" {
				continue
			}
			t.WLine("	  t.%s = %s", cname, cname)
		}
	*/
	t.WLine("	}")

}

// 生成json结构体
func (t *ClassJsonSheetTable) DoCfgData(rows [][]string, isbool bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.WLine("	\"%s\":{", t.SheetName)
	l := len(rows)

	for x, row := range rows {
		if x < 5 {
			continue
		}
		if row == nil {
			continue
		}
		canshuzhi := ""
		_t := ","

		//添加参数
		for index, cname := range t.ctypeNameList {
			colCell := ""
			if cname == "" {
				continue
			}
			ctype := t.GetCType(index)

			switch ctype {
			case t_int:
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						colCell = t.GetcTypeValue(index)
					}
				}
			case t_ints:
				var data []int32
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = "" //t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						//解析二维数组
						var dv1 []int32
						dv1, _ = sys_string.SplitToInt32(colCell, ",")
						data = dv1

					}
				}
				str, _ := sys_json.MarshalToString(data)
				colCell = str
			case t_int2s:
				var data [][]int32
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = "" //t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						//解析二维数组
						v1s := strings.Split(colCell, ";")
						for _, v1 := range v1s {
							if v1 == "" {
								continue
							}
							var dv1 []int32
							dv1, _ = sys_string.SplitToInt32(v1, ",")
							data = append(data, dv1)
						}

					}
				}
				str, _ := sys_json.MarshalToString(data)
				colCell = str

			case t_string:
				if index > len(row)-1 {
					colCell = fmt.Sprintf("\"\"")
				} else {
					colCell = fmt.Sprintf("\"%s\"", row[index])
				}
			case t_float:
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if colCell == "" {
						colCell = t.GetcTypeValue(index)
					}
					/*
						if ctype == t_float {
							colCell = row[index] + "f"
						}*/
				}
			case t_floats:
				var data []int32
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = "" //t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						//解析二维数组
						var dv1 []int32
						dv1, _ = sys_string.SplitToInt32(colCell, ",")
						data = dv1

					}
				}
				str, _ := sys_json.MarshalToString(data)
				colCell = str
			case t_float2s:
				var data [][]int32
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = "" //t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						//解析二维数组
						v1s := strings.Split(colCell, ";")
						for _, v1 := range v1s {
							if v1 == "" {
								continue
							}
							var dv1 []int32
							dv1, _ = sys_string.SplitToInt32(v1, ",")
							data = append(data, dv1)
						}

					}
				}
				str, _ := sys_json.MarshalToString(data)
				colCell = str
			case t_strings:
				var data [][]string
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = "" //t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						//解析二维数组
						v1s := strings.Split(colCell, ";")
						for _, v1 := range v1s {
							if v1 == "" {
								continue
							}
							var dv1 []string
							dv1 = strings.Split(v1, ",")
							data = append(data, dv1)
						}

					}
				}
				str, _ := sys_json.MarshalToString(data)
				colCell = str
			case t_string2s:
				var data [][]string
				if index > len(row)-1 {
					//没有值根据类型 给默认值
					colCell = "" //t.GetcTypeValue(index)
				} else {
					colCell = row[index]
					if !sys_base.StringIsNullOrEmpty(colCell) {
						//解析二维数组
						v1s := strings.Split(colCell, ";")
						for _, v1 := range v1s {
							if v1 == "" {
								continue
							}
							var dv1 []string
							dv1 = strings.Split(v1, ",")
							data = append(data, dv1)
						}

					}
				}
				str, _ := sys_json.MarshalToString(data)
				colCell = str
			default:
				if index > len(row)-1 {
					colCell = fmt.Sprintf("\"\"")
				} else {
					colCell = fmt.Sprintf("\"%s\"", row[index])
				}
				//log.Fatalf("有错误 ctype =%s 错误", ctype)
			}

			if index == len(t.ctypeNameList)-1 {
				_t = ""
			}

			tv := fmt.Sprintf("\"%s\":%s", cname, colCell)
			canshuzhi += tv + _t

		}
		if x+1 == l {
			t.WLine("		\"%s\":{%s}", row[0], canshuzhi)
		} else {
			t.WLine("		\"%s\":{%s},", row[0], canshuzhi)
		}

	}
	if isbool {
		t.WLine("	}")
	} else {
		t.WLine("	},")
	}

}
func (t *ClassJsonSheetTable) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}

func (t *ClassJsonSheetTable) GetContent() string {
	return t.file_content
}
