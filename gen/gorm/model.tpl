package model

import "{{GenImportTime .Fields}}"
//{{.Remark}}  
type {{.ObjectName}} struct {
    {{- range $k,$v := .Fields}} 
    //{{$v.Remark}}
    {{$v.FieldName}}  {{$v.FieldType}}  `gorm:"column:{{$v.ColumnName}} {{- if eq $v.IsPk true}};PRIMARY_KEY{{end}}" json:"{{$v.FieldName}},omitempty"`
    {{- end}}
}

func (u {{.ObjectName}}) TableName() string {
	return "{{ if IsNotBlank .Schema}}{{.Schema}}.{{end}}{{.TableName}}"
}
