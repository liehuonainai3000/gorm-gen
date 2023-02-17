package gen

import (
	"fmt"
	"strings"
	"text/template"
)

type funcList struct{}

var funcs funcList

var funcMap template.FuncMap = template.FuncMap{
	"GenFieldHolders": funcs.GenFieldHolders,
	"GenColumnNames":  funcs.GenColumnNames,
	"GenUpdateStmt":   funcs.GenUpdateStmt,
	"GenPKWhere":      funcs.GenPKWhere,
	"ToLower":         funcs.ToLower,
	"LowerFirstChar":  funcs.LowerFirstChar,
	"GenImportTime":   funcs.GenImportTime,
	"IsNotBlank":      funcs.IsNotBlank,
}

// 生成sql字段列表，以逗号分隔
func (o *funcList) GenColumnNames(fields []Field) string {

	var names []string
	for _, field := range fields {
		names = append(names, field.ColumnName)
	}

	return strings.Join(names, ",")
}

// 生成sql字段值占位符,用于insert 的values, 如： :name,:age
func (o *funcList) GenFieldHolders(fields []Field) string {

	var names []string
	for _, field := range fields {
		names = append(names, ":"+field.ColumnName)
	}

	return strings.Join(names, ",")
}

// 生成update语句片段
func (o *funcList) GenUpdateStmt(idx int, f Field) string {
	s := fmt.Sprintf(" %s=:%s ", f.ColumnName, f.ColumnName)

	if idx > 0 {
		s = "," + s
	}
	return s
}

func (o *funcList) GenPKWhere(idx int, pk Field) string {

	s := fmt.Sprintf(" %s=:%s ", pk.ColumnName, pk.ColumnName)
	if idx > 0 {
		s = " and " + s
	}
	return s

}

// 判断字段列表中是否有time.Time类型,如果有则导入time包
func (o *funcList) GenImportTime(fields []Field) string {

	hasTime := false
	for _, field := range fields {
		if field.FieldType == "time.Time" {

			hasTime = true
			break
		}
	}
	if hasTime {
		return "time"
	} else {
		return ""
	}
}

func (o *funcList) IsNotBlank(str string) bool {
	return strings.Trim(str, " ") != ""
}

func (o *funcList) LowerFirstChar(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToLower(str[0:1]) + str[1:]
}

func (o *funcList) ToLower(str string) string {
	return strings.ToLower(str)
}
