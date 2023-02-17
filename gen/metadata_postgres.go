package gen

import (
	"fmt"
	"gorm-gen/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const sql_postgres = `
SELECT
	A.ordinal_position,
	A.column_name,
	CASE A.is_nullable WHEN 'NO' THEN 0 ELSE 1 END AS nullable,
	A.data_type,
	coalesce(A.character_maximum_length, A.numeric_precision, -1) as length,
	A.numeric_scale,
    CASE WHEN length(B.attname) > 0 THEN 1 ELSE 0 END AS is_pk,
    d.description as remark
 FROM
	information_schema.columns A
        JOIN pg_class c ON c.relname = a.table_name
        LEFT JOIN pg_description d ON d.objoid = c.oid AND d.objsubid = a.ordinal_position

 LEFT JOIN (
	 SELECT
		 pg_attribute.attname
	 FROM
		 pg_index
		 left join pg_class on pg_index.indrelid = pg_class.oid
		 left join pg_attribute on pg_attribute.attrelid = pg_class.oid AND pg_attribute.attnum = ANY (pg_index.indkey)
	 WHERE
		 pg_class.oid = ? :: regclass
	
	 
 ) B ON A.column_name = b.attname
 WHERE
	A.table_schema = ? 
 AND A.table_name = ?
 ORDER BY
	ordinal_position ASC;	
`

// 用于初始化postgres数据库表的元数据
type initMetaData_Postgres struct {
	db *gorm.DB
}

func NewInitMetaDataPostgres(dbConf *utils.DBConfig) *initMetaData_Postgres {
	db := InitPostgresDB(dbConf)
	if utils.Global.Debug {
		db = db.Debug()
	}
	return &initMetaData_Postgres{
		db: db,
	}
}

func InitPostgresDB(dbConf *utils.DBConfig) *gorm.DB {

	dnsFmt := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf(dnsFmt, dbConf.Host, dbConf.User, dbConf.Password, dbConf.DBName, dbConf.Port)

	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// DB = _db.Debug()

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(5)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于5，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	return _db
}

func (f initMetaData_Postgres) QueryMetaData(schema, tableName string) ([]Field, error) {

	rst := []Field{}
	if schema == "" {
		schema = "public"
	}
	//查询数据库获取表的元数据
	err := f.db.Raw(sql_postgres, tableName, schema, tableName).Scan(&rst).Error
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
// func (o initMetaData_Postgres) MapFieldType(f *Field) {

// 	f.FieldName = toCamelName(f.ColumnName)
// 	switch f.DataType {
// 	case "integer":
// 		f.FieldType = "int"
// 	case "bigint":
// 		f.FieldType = "int64"
// 	case "serial":
// 		f.FieldType = "int64"
// 	case "character varying":
// 		f.FieldType = "string"
// 	case "character":
// 		f.FieldType = "string"
// 	case "text":
// 		f.FieldType = "string"
// 	case "boolean":
// 		f.FieldType = "bool"
// 	case "decimal":
// 		f.FieldType = "float64"
// 	case "numeric":
// 		f.FieldType = "float64"
// 	case "timestamp without time zone":
// 		f.FieldType = "time.Time"
// 	case "timestamp with time zone":
// 		f.FieldType = "time.Time"
// 	default:
// 		f.FieldType = "string"
// 	}
// }
