package do

type {{.ObjectName}} struct {
    {{- range $k,$v := .Fields}} 
    //{{$v.Remark}}
    {{$v.FieldName}}  any  `db:"{{$v.ColumnName}}" json:"{{$v.ColumnName}}"`
    {{- end}}

}