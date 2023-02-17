package gen

import (
	"errors"
	"fmt"
	"gorm-gen/utils"
)

// 初始化数据表的元数据，不同数据库需要不同的实现
type MetaQueryer interface {
	//查询表中的字段信息
	QueryMetaData(schema, tablename string) ([]Field, error)
}

var metaQueryers map[string]MetaQueryer = make(map[string]MetaQueryer)

func RegisteMetaQuery(dbType string, metaQuery MetaQueryer) {
	metaQueryers[dbType] = metaQuery
}

func InitMetaQueryers() {

	metaQueryers = make(map[string]MetaQueryer)
	var mq MetaQueryer
	for k, v := range utils.Global.DBConfigs {
		if v.DBType == "pg" {
			mq = NewInitMetaDataPostgres(&v)
		} else if v.DBType == "mysql" {
			mq = NewInitMetaDataMysql(&v)

		}

		metaQueryers[k] = mq
	}
}

// 根据指定的数据库代码返回数据库元数据生成器
func GetMetaQueryer(dbCode string) (m MetaQueryer, err error) {

	v, ok := metaQueryers[dbCode]

	if !ok {
		return nil, errors.New("No metaqueryer found:" + dbCode)
	}
	return v, nil
}

func mapTableFieldType(t *TableTemplate) error {

	dbCfg, ok := utils.Global.DBConfigs[t.DBCode]

	if !ok {
		return fmt.Errorf("db config not found : %s", t.DBCode)
	}
	fieldMap, ok := utils.Global.FieldTypeMap[dbCfg.DBType]

	if !ok {
		return fmt.Errorf("fieldType map not found : %s", dbCfg.DBType)
	}

	for i, f := range t.Fields {
		f.FieldName = toCamelName(f.ColumnName)
		f.FieldType, ok = fieldMap[f.DataType]
		if !ok {
			f.FieldType = "string"
		}
		t.Fields[i] = f
	}

	return nil
}
