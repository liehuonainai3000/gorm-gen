package gen

import (
	"gorm-gen/utils"
	"testing"
)

func TestGenerate(t *testing.T) {

	utils.InitGlobalConfig()
	dt := &TableTemplate{
		PackageBasePath: "gorm-gen/internal",
		GeneratePath:    "D:\\workspace-go\\src\\gorm-gen\\internal",
		Schema:          "testdb",
		DBCode:          "mysql1",

		TableName:  "sys_office",
		ObjectName: "Office",
		Remark:     "机构信息",
	}

	err := GenerateFile(dt, NewInitMetaDataMysql(&utils.DBConfig{
		DBType:   "mysql",
		DBName:   "testdb",
		Host:     "192.168.19.141",
		Port:     3306,
		User:     "root",
		Password: "root123",
	}))
	if err != nil {
		t.Errorf("generate err :%v", err)
	}
}

func TestMapTableFieldType(t *testing.T) {
	utils.InitGlobalConfig()
	qm := &initMetaData_Mysql{
		db: _db,
	}
	f, err := qm.QueryMetaData("testdb", "sys_office")
	if err != nil {
		t.Errorf("err:%v", err)
	}
	tt := &TableTemplate{
		Schema:    "testdb",
		TableName: "sys_office",
		DBCode:    "mysql1",
		Fields:    f,
	}
	err = mapTableFieldType(tt)
	if err != nil {
		t.Errorf("err:%+v", err)
		return
	}

	t.Logf("fields:%+v", tt.Fields)

}

func TestChangeSliceValue(t *testing.T) {
	v := []*Field{{FieldName: "f1", FieldType: "string"}, {FieldName: "f2", FieldType: "int"}}
	t.Log("1: ", v[0])
	for _, itm := range v {
		itm.FieldType = "string-1"
	}
	t.Log("1: ", v[0])

	v1 := []Field{{FieldName: "f1", FieldType: "string"}, {FieldName: "f2", FieldType: "int"}}
	t.Log("2: ", v1[0])
	for _, itm := range v1 {
		itm.FieldType = "string-1"
	}
	t.Log("2: ", v1[0])

	v2 := []Field{{FieldName: "f1", FieldType: "string"}, {FieldName: "f2", FieldType: "int"}}
	t.Log("3: ", v2[0])
	for _, itm := range v2 {
		changeValue(&itm)
	}
	t.Log("3: ", v2[0])
}
func changeValue(f *Field) {
	f.FieldType = "string-1"
}
