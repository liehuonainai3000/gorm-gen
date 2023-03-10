package gen

import (
	"gorm-gen/utils"
	"testing"
)

func TestQueryMetaData(t *testing.T) {

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

	t.Logf("rst:%+v", f)
}
