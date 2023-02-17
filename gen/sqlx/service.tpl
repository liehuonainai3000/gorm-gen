package service

import (
	"{{.PackageBasePath}}/dao"
	"{{.PackageBasePath}}/do"
	"{{.PackageBasePath}}/entity"

	"github.com/jmoiron/sqlx"
)
{{$serviceName := LowerFirstChar .ObjectName}}

type {{$serviceName}}Service struct {
	{{ToLower .ObjectName}}Dao dao.{{.ObjectName}}Dao
}

func New{{.ObjectName}}Service(db *sqlx.DB) *{{$serviceName}}Service {
	return &{{$serviceName}}Service{
		{{ToLower .ObjectName}}Dao: *dao.New{{.ObjectName}}Dao(db),
	}
}

func (s *{{$serviceName}}Service) Insert(o *do.{{.ObjectName}}) error {
	return s.{{ToLower .ObjectName}}Dao.Insert(o)
}

func (s *{{$serviceName}}Service) UpdateByPK(o *do.{{.ObjectName}}) error {
	return s.{{ToLower .ObjectName}}Dao.UpdateByPK(o)
}

func (s *{{$serviceName}}Service) DeleteByPK(o *do.{{.ObjectName}}) error {
	return s.{{ToLower .ObjectName}}Dao.DeleteByPK(o)
}

func (s *{{$serviceName}}Service) GetByPK(o *do.{{.ObjectName}}) (*entity.{{.ObjectName}}, error) {
	return s.{{ToLower .ObjectName}}Dao.GetByPK(o)
}

func (s *{{$serviceName}}Service) Select(o *do.{{.ObjectName}}) ([]entity.{{.ObjectName}}, error) {
	return s.{{ToLower .ObjectName}}Dao.Select(o)
}
