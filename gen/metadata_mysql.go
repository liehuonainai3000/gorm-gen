package gen

import (
	"fmt"
	"gorm-gen/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const mysql_sql = `
select 
column_name ,
ordinal_position as "ordinal_position",

case when IS_NULLABLE='YES' then 1 else 0 end as "nullable",
CHARACTER_MAXIMUM_LENGTH as "length",
data_type as "data_type",
numeric_scale as "numeric_scale",
case when COLUMN_KEY='PRI' then 1 else 0 end as "is_pk"

from information_schema.columns where table_schema = ?
and table_name = ?
 ORDER BY
	ordinal_position ASC;	
`

// 用于初始化Mysql数据库表的元数据
type initMetaData_Mysql struct {
	db *gorm.DB
}

func NewInitMetaDataMysql(dbConf *utils.DBConfig) *initMetaData_Mysql {
	db := InitMysqlDB(dbConf)
	if utils.Global.Debug {
		db = db.Debug()
	}
	return &initMetaData_Mysql{
		db: db,
	}
}

func InitMysqlDB(dbConf *utils.DBConfig) *gorm.DB {

	dnsFmt := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dnsFmt, dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(5)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于5，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	return _db
}

func (f initMetaData_Mysql) QueryMetaData(schema, tableName string) ([]Field, error) {

	rst := []Field{}
	if schema == "" {
		schema = "public"
	}
	//查询数据库获取表的元数据
	err := f.db.Raw(mysql_sql, schema, tableName).Scan(&rst).Error
	if err != nil {
		Logger.Errorf("Query Database err :%+v", err)
		return nil, err
	}

	if len(rst) == 0 {
		return nil, fmt.Errorf("schema:%s,tablename:%s not found ", schema, tableName)
	}

	return rst, nil

}

// 数据库类型对应关系
// func (o initMetaData_Mysql) MapFieldType(f *Field) {

// 	f.FieldName = toCamelName(f.ColumnName)
// 	switch f.DataType {
// 	case "int":
// 		f.FieldType = "int"
// 	case "bigint":
// 		f.FieldType = "int64"
// 	case "serial":
// 		f.FieldType = "int64"
// 	case "varchar":
// 		f.FieldType = "string"
// 	case "char":
// 		f.FieldType = "string"
// 	case "text":
// 		f.FieldType = "string"
// 	case "boolean":
// 		f.FieldType = "bool"
// 	case "decimal":
// 		f.FieldType = "float64"
// 	case "numeric":
// 		f.FieldType = "float64"
// 	case "timestamp":
// 		f.FieldType = "time.Time"
// 	default:
// 		f.FieldType = "string"
// 	}
// }
