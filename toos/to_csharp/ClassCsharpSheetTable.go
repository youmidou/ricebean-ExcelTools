package to_csharp

import (
	"ExcelToolGenerationConfig/TudouFramework/sys_base"
	"fmt"
)

const (
	t_int      = "int"
	t_ints     = "int[]"
	t_int2s    = "int[][]"
	t_float    = "float"
	t_floats   = "float[]"
	t_float2s  = "float[][]"
	t_string   = "string"
	t_strings  = "string[]"
	t_string2s = "string[][]"
)

type ClassCsharpSheetTable struct {
	FileName            string //文件名字
	SheetName           string //表名字
	file_class_content  string //
	file_init_content   string //
	file_content        string //最终str
	classBaseInfoName   string
	ctypeNameList       []string
	CtypeList           []string
	CtypeAnnotationList []string
	rows                [][]string
	file_content2       string
}

func (t *ClassCsharpSheetTable) Init(FileName string, SheetName string, rows [][]string) {
	t.FileName = FileName
	t.SheetName = SheetName
	t.rows = rows

	t.DoBaseInfo(rows)
	t.DoClass()
	t.DoCfgData(rows)
}

/*
*
0:规则第一行不要 备注用
1:第二行 属性名字 当为空字符串 本列不做数据导入
2：第三行 类型默认 int
3:属性说明
4:属性功能使用说明
*/
func (t *ClassCsharpSheetTable) DoBaseInfo(rows [][]string) {
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

// 获取IdKey
func (t *ClassCsharpSheetTable) GetIdKeyType() string {
	//返回值
	ctype := t.GetcTypeName(0)
	switch ctype {
	case "int":
		return "int"
	case "string":
		return "string"
	}
	return ctype
}

func (t *ClassCsharpSheetTable) GetCType(index int) string {
	//返回名字 默认int
	ctype := ""
	if index < len(t.CtypeList) {
		ctype = t.CtypeList[index]
	}
	return ctype
}

func (t *ClassCsharpSheetTable) GetcTypeName(index int) string {
	//返回名字 默认int
	ctype := ""
	if index < len(t.CtypeList) {
		ctype = t.CtypeList[index]
	}
	switch ctype {
	case t_int:
		ctype = "int"
	case t_float:
		ctype = "float"
	case t_string:
		ctype = "string"

		/*
			case t_ints:
				ctype = "int[]"
			case t_int2s:
				ctype = "int[][]"
			case t_floats:
				ctype = "float[]"
			case t_float2s:
				ctype = "float[][]"
			case t_strings:
				ctype = "string[]"
			case t_string2s:
				ctype = "string[][]"
		*/
	default:
		ctype = "string"

	}
	return ctype
}

// 默认值
func (t *ClassCsharpSheetTable) GetcTypeValue(index int) string {
	//返回值
	//返回名字 默认int
	ctype := ""
	if index < len(t.CtypeList) {
		ctype = t.CtypeList[index]
	}
	switch ctype {
	case t_int:
		return "0"
	case t_string:
		return "\"\""
	case t_float:
		return "0f"
		/*
			case t_ints:
				return "[]int"
			case t_int2s:
				return "0"
			case t_floats:
				return "0f"
			case t_float2s:
				return "0f"
			case t_strings:
				return "\"\""
			case t_string2s:
				return "\"\""
		*/
	default:
		return "\"\""

	}
	return ctype
}

// 获取注释
func (t *ClassCsharpSheetTable) GetcTypeAnnotation(index int) string {
	zhushi := ""
	if index < len(t.CtypeAnnotationList) {
		zhushi = t.CtypeAnnotationList[index]
	}
	return zhushi
}
func (t *ClassCsharpSheetTable) DoClass() {
	baseInfo_data := ""

	t.WLine("public class " + t.classBaseInfoName)
	t.WLine("{")
	//添加参数
	for index, cname := range t.ctypeNameList {
		if cname == "" {
			continue
		}
		ctype := t.GetcTypeName(index)
		zhushi := t.GetcTypeAnnotation(index)
		t.WLine("   /** %s */", zhushi)
		t.WLine("	public %s %s{ get; set; }", ctype, cname)
		//t.WLine("	public %s %s;//%s", ctype, cname, zhushi)

		_temp := ","
		//拿到参数串
		if index == len(t.ctypeNameList)-1 {
			_temp = ""
		}
		baseInfo_data += ctype + " " + cname + _temp
	}

	//t.WLine("}")
	//--------------------------------------------------------------
	//构建函数  public FileNameInfo()
	t.WLine("	public %s(%s)", t.classBaseInfoName, baseInfo_data)
	t.WLine("	{") //参数赋值
	for _, cname := range t.ctypeNameList {
		if cname == "" {
			continue
		}
		t.WLine("	  this.%s = %s;", cname, cname)
	}
	t.WLine("	}")

	t.WLine("	public %s(){}", t.classBaseInfoName)
	//--------------------------------------------------------------
	//克隆
	t.WLine("	public %s Clone()", t.classBaseInfoName)
	t.WLine("	{")

	t.WLine("		return new %s(){", t.classBaseInfoName)
	baseInfo_data = ""
	//添加参数
	for _, cname := range t.ctypeNameList {
		if cname == "" {
			continue
		}
		//zhushi := t.GetcTypeAnnotation(index)
		//t.WLine("   /** %s */", zhushi)
		t.WLine("			%s = this.%s,", cname, cname)
	}
	t.WLine("		};")

	t.WLine("	}")
	t.WLine("}")
}

func (t *ClassCsharpSheetTable) WLine(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content += aline + "\n"
}
func (t *ClassCsharpSheetTable) GetContent() string {
	return t.file_content
}
func (t *ClassCsharpSheetTable) WLine2(format string, a ...any) {
	aline := fmt.Sprintf(format, a...)
	t.file_content2 += aline + "\n"
}
func (t *ClassCsharpSheetTable) GetDataContent() string {
	return t.file_content2
}
func (t *ClassCsharpSheetTable) DoCfgData(rows [][]string) {

	t.classBaseInfoName = fmt.Sprintf("%s_%s_Item", t.FileName, t.SheetName)
	keyIdType := t.GetIdKeyType()
	t.WLine2("	public static Dictionary<%s,%s> %s = new Dictionary<%s,%s>()",
		keyIdType, t.classBaseInfoName, t.SheetName,
		keyIdType, t.classBaseInfoName)

	t.WLine2("	{")
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

					if sys_base.StringIsNullOrEmpty(colCell) {
						colCell = t.GetcTypeValue(index)
					}
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
					if ctype == t_float {
						colCell = row[index] + "f"
					}
				}
			case t_string:
				if index > len(row)-1 {
					colCell = fmt.Sprintf("\"\"")
				} else {
					colCell = fmt.Sprintf("\"%s\"", row[index])
				}
				/*
					case t_ints:
						var data string
						if index > len(row)-1 {
							//没有值根据类型 给默认值
							colCell = "null" //t.GetcTypeValue(index)
						} else {
							colCell = row[index]
							if !sys_base.StringIsNullOrEmpty(colCell) {
								//解析二维数组
								//var dv1 []int32
								result := strings.Split(colCell, ",")
								data = t.GetSList(result)

							}
						}
						colCell = data
					case t_int2s:
						var datav []string
						var datav2 string
						if index > len(row)-1 {
							//没有值根据类型 给默认值
							colCell = "null" //t.GetcTypeValue(index)
						} else {
							colCell = row[index]
							if !sys_base.StringIsNullOrEmpty(colCell) {
								//解析二维数组
								v1s := strings.Split(colCell, ";")
								for _, v1 := range v1s {
									if v1 == "" {
										continue
									}
									var dv1 string
									result := strings.Split(v1, ",")
									dv1 = t.GetSList(result)
									datav = append(datav, dv1)
								}
								datav2 = t.GetSList(datav)

							}
						}
						colCell = datav2
					case t_strings:
						var data string
						if index > len(row)-1 {
							//没有值根据类型 给默认值
							colCell = "null" //t.GetcTypeValue(index)
						} else {
							colCell = row[index]
							if !sys_base.StringIsNullOrEmpty(colCell) {
								//解析二维数组
								//var dv1 []int32
								result := strings.Split(colCell, ",")
								data = t.GetStringList(result)
							}
						}
						colCell = data

						case t_string2s:
							var datav []string
							var datav2 string
							if index > len(row)-1 {
								//没有值根据类型 给默认值
								colCell = "null" //t.GetcTypeValue(index)
							} else {
								colCell = row[index]
								if !sys_base.StringIsNullOrEmpty(colCell) {
									//解析二维数组
									v1s := strings.Split(colCell, ";")
									for _, v1 := range v1s {
										if v1 == "" {
											continue
										}
										var dv1 string
										result := strings.Split(v1, ",")
										dv1 = t.GetStringList(result)
										datav = append(datav, dv1)
									}
									datav2 = t.GetStringList2(datav)

								}
							}
							colCell = datav2

				*/
			/*
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
			*/
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
			canshuzhi += colCell + _t
		}

		//keyIdType := t.GetIdKeyType()
		if row[0] != "" {
			if keyIdType == t_int {
				t.WLine2("		[%s] = new %s(%s),", row[0], t.classBaseInfoName, canshuzhi)
			} else {
				t.WLine2("		[\"%s\"] = new %s(%s),", row[0], t.classBaseInfoName, canshuzhi)
			}
		}

		/*
			if x == 20 {
				break
			}
		*/

	}
	t.WLine2("	};")
}
func (t *ClassCsharpSheetTable) GetSList(result []string) string {
	cstr := "new int[]{"
	l := len(result)
	for i, s := range result {
		if s == "" {
			continue
		}
		if i != l-1 {
			cstr += s + ","
		} else {
			cstr += s
		}
	}

	cstr += "}"
	return cstr
}
func (t *ClassCsharpSheetTable) GetStringList(result []string) string {
	cstr := "new string[]{"
	l := len(result)
	for i, s := range result {
		if s == "" {
			continue
		}
		if i != l-1 {
			cstr += fmt.Sprintf("\"%s\",", s)
		} else {
			cstr += fmt.Sprintf("\"%s\"", s)
		}
	}

	cstr += "}"
	return cstr
}
func (t *ClassCsharpSheetTable) GetStringList2(result []string) string {
	cstr := "new string[]{"
	l := len(result)
	for i, s := range result {
		if s == "" {
			continue
		}
		if i != l-1 {
			cstr += fmt.Sprintf("%s", s)
		} else {
			cstr += s
		}
	}

	cstr += "}"
	return cstr
}
