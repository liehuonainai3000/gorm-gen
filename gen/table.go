package gen

type TableTemplate struct {
	//数据库代码，代表数据库配置，需要指定
	DBCode string `json:"db_code" validate:"required,not-blank"`
	//包路径,需要指定
	PackageBasePath string `json:"package_base_path" validate:"required,not-blank"`
	//生成路径,需要指定
	GeneratePath string `json:"generate_path" validate:"required,not-blank"`
	//表名,需要指定
	TableName string `json:"table_name" validate:"required,not-blank"`
	//模式,需要指定
	Schema string `json:"schema"`
	//对象说明,需要指定
	Remark string `json:"remark" validate:"required"`
	//对象名,需要指定
	ObjectName string `json:"object_name" validate:"required,not-blank"`

	//生成类型：gorm/sqlx,默认gorm
	GenerateType string `json:"generate_type"`

	//创建文件类型，目前有：model/service,如果为空则全部生成
	CreateFileTypes []string `json:"create_file_types"`

	//字段列表，通过MetaQueryer生成
	Fields []Field
	//主键，通过MetaQueryer生成
	PrimaryKey []Field
}

// 表字段信息
type Field struct {
	//排序，非必须
	OrdinalPosition int
	//表字段名称，必须
	ColumnName string
	//是否可空
	Nullable bool
	//表字段数据类型，必须
	DataType string

	//表字段长度
	Length int
	//表字段精度
	NumericScale int
	//是否主键，必须
	IsPk bool
	//字段备注
	Remark string

	//生成struct的字段名称，通过InitMetaData生成
	FieldName string
	//生成struct的字段类型，通过InitMetaData生成
	FieldType string
}
