package gen

import (
	"fmt"
	"gorm-gen/utils"
	"testing"
)

func TestConnMysql(t *testing.T) {
	cfg := &utils.DBConfig{
		DBType:   "mysql",
		DBName:   "testdb",
		Host:     "192.168.19.141",
		Port:     3306,
		User:     "root",
		Password: "root123",
	}
	db := InitMysqlDB(cfg)
	fmt.Println(db)
}
