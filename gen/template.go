package gen

import (
	_ "embed"
	"os"
	"path"
	"strings"
	"text/template"
)

//go:embed gorm/model.tpl
var gorm_model string

//go:embed gorm/service.tpl
var gorm_service string

//go:embed sqlx/do.tpl
var sqlx_do string

//go:embed sqlx/entity.tpl
var sqlx_entity string

//go:embed sqlx/dao.tpl
var sqlx_dao string

//go:embed sqlx/service.tpl
var sqlx_service string

var gorm_tpls map[string]*template.Template = make(map[string]*template.Template)
var sqlx_tpls map[string]*template.Template = make(map[string]*template.Template)

func init() {

	gorm_modeltpl, err := template.New("gorm_model").Funcs(funcMap).Parse(gorm_model)
	if err != nil {
		panic(err)
	}
	gorm_tpls["model"] = gorm_modeltpl

	gorm_servicetpl, err := template.New("gorm_service").Funcs(funcMap).Parse(gorm_service)
	if err != nil {
		panic(err)
	}
	gorm_tpls["service"] = gorm_servicetpl

	sqlx_daotpl, err := template.New("sqlx_dao").Funcs(funcMap).Parse(sqlx_dao)
	if err != nil {
		panic(err)
	}
	sqlx_tpls["dao"] = sqlx_daotpl

	sqlx_dotpl, err := template.New("sqlx_do").Funcs(funcMap).Parse(sqlx_do)
	if err != nil {
		panic(err)
	}
	sqlx_tpls["do"] = sqlx_dotpl

	sqlx_entitytpl, err := template.New("sqlx_entity").Funcs(funcMap).Parse(sqlx_entity)
	if err != nil {
		panic(err)
	}
	sqlx_tpls["entity"] = sqlx_entitytpl

	sqlx_servicetpl, err := template.New("sqlx_service").Funcs(funcMap).Parse(sqlx_service)
	if err != nil {
		panic(err)
	}
	sqlx_tpls["service"] = sqlx_servicetpl
}

// 查询表结构元数据
func queryMetaData(tt *TableTemplate, metaQueryer MetaQueryer) (err error) {

	tt.Fields, err = metaQueryer.QueryMetaData(tt.Schema, tt.TableName)
	if err != nil {
		Logger.Error(err)
		return err
	}

	//根据查询的表元数据，生成对应struct中的FieldType,FieldName和主键PrimaryKey
	err = mapTableFieldType(tt)
	if err != nil {
		Logger.Error(err)
		return err

	}
	for _, f := range tt.Fields {
		if f.IsPk {
			tt.PrimaryKey = append(tt.PrimaryKey, f)
		}
	}
	return nil
}

// 根据TableTemplate配置生成代码文件
func GenerateFile(o *TableTemplate, metaQueryer MetaQueryer) error {

	err := queryMetaData(o, metaQueryer)
	if err != nil {
		Logger.Error(err)
		return err
	}

	tpls := gorm_tpls
	if o.GenerateType == "sqlx" {
		tpls = sqlx_tpls
	}

	for k, v := range tpls {

		if o.CreateFileTypes != nil && !arrayContains(o.CreateFileTypes, k) {
			continue
		}
		err := os.MkdirAll(path.Join(o.GeneratePath, k), 0666)
		if err != nil {
			Logger.Error(err)
			return err
		}

		f, err := os.Create(path.Join(o.GeneratePath, k, strings.ToLower(o.ObjectName)+".go"))

		if err != nil {
			Logger.Error(err)
			return err
		}

		defer f.Close()
		err = v.Execute(f, o)
		if err != nil {
			Logger.Error(err)
			return err
		}

	}
	return nil

}

func arrayContains[T comparable](arr []T, itm T) bool {

	for _, c := range arr {
		if c == itm {
			return true
		}
	}
	return false
}

// 将下划线分割的字段名转换为驼峰形式
func toCamelName(name string) string {
	itms := strings.Split(name, "_")
	var newItm []string = make([]string, len(itms))
	for i, itm := range itms {
		if itm == "" {
			newItm[i] = ""
		} else {
			newItm[i] = strings.ToUpper(itm[0:1]) + strings.ToLower(itm[1:])
		}
	}
	return strings.Join(newItm, "")
}
