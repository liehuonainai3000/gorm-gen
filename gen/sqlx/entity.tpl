package entity

import "{{GenImportTime .Fields}}"

type {{.ObjectName}} struct {
    {{- range $k,$v := .Fields}} 
    //{{$v.Remark}}
    {{$v.FieldName}}  {{$v.FieldType}}  `json:"{{$v.ColumnName}},omitempty"`
    {{- end}}
}