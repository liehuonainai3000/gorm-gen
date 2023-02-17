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
	_db := InitMysqlDB(&utils.DBConfig{
		Host:     "192.168.19.141",
		Port:     3306,
		DBName:   "testdb",
		User:     "root",
		Password: "root123",
	})
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
