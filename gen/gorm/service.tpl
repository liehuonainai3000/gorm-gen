package service

import (
	"{{.PackageBasePath}}/model"
	"gorm.io/gorm"
)
{{$serviceName := LowerFirstChar .ObjectName}}

type {{$serviceName}}Service struct {
	db *gorm.DB
}

// 获得Service实例
func New{{.ObjectName}}Service(db *gorm.DB) *{{$serviceName}}Service {
	return &{{$serviceName}}Service{
		db: db,
	}
}

// 数据插入
func (s *{{$serviceName}}Service) Insert(mo *model.{{.ObjectName}}) error {
	return s.db.Create(mo).Error
}

// 按主键更新
func (s *{{$serviceName}}Service) UpdateByPk(mo map[string]any) error {
	return s.db.Model(&model.{{.ObjectName}}{}){{- range $k,$v := .PrimaryKey}}.Where("{{$v.ColumnName}}=?",mo["{{$v.FieldName}}"]){{end}}.Updates(mo).Error
}

// 按主键删除
func (s *{{$serviceName}}Service) DeleteByPk(mo *model.{{.ObjectName}}) error {
	return s.db{{- range $k,$v := .PrimaryKey}}.Where("{{$v.ColumnName}}=?",mo.{{$v.FieldName}}){{end}}.Delete(&model.{{.ObjectName}}{}).Error
}

// 按主键查询一条记录
func (s *{{$serviceName}}Service) GetByPk(mo *model.{{.ObjectName}}) (*model.{{.ObjectName}},error) {
	rst := &model.{{.ObjectName}}{}
	err := s.db.Model(&model.{{.ObjectName}}{}){{- range $k,$v := .PrimaryKey}}.Where("{{$v.ColumnName}}=?",mo.{{$v.FieldName}}){{end}}.Scan(rst).Error
	if err != nil{
		return nil,err
	}
	return rst,nil
}

// 按条件查询数据列表
func (s *{{$serviceName}}Service) QueryList(mo map[string]any) ([]model.{{.ObjectName}}, error) {

	rst := []model.{{.ObjectName}}{}
	stmt := s.db.Model(&model.{{.ObjectName}}{})
	{{- range $k,$v := .Fields}}
	if mo["{{$v.FieldName}}"] != nil {
		stmt.Where("{{$v.ColumnName}}=?",mo["{{$v.FieldName}}"])
	}
	{{- end}}
	err := stmt.Scan(&rst).Error
	if err != nil {
		return nil, err
	}
	return rst, nil
}

// 执行sql查询，返回map切片数据格式
func (s *{{$serviceName}}Service) ExecuteQuery(sql string, params ...any) (rst map[string]any, err error) {
	rst = make(map[string]any)
	err = s.db.Raw(sql, params...).Scan(&rst).Error
	return
}

// 执行sql更新
func (s *{{$serviceName}}Service) ExecuteUpdate(sql string, params ...any) error {
	return s.db.Raw(sql, params...).Error
}
