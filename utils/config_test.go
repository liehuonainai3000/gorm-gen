package utils

import "testing"

func init() {
	InitGlobalConfig()
}

func TestConfig(t *testing.T) {

	t.Logf("%+v", Global.FieldTypeMap)
	t.Logf("%+v", Global.DBConfigs)
	t.Logf("%+v", Global.Server)
}
