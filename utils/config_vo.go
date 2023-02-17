package utils

type GlobalConf struct {
	Debug        bool                         `json:"debug"`
	Server       *Server                      `json:"server"`
	DBConfigs    map[string]DBConfig          `json:"DBConfigs"`
	FieldTypeMap map[string]map[string]string `json:"fieldTypeMap"`
}
type DBConfig struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}
type Server struct {
	Port int `json:"port"`
}
