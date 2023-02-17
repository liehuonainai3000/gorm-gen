
## 使用说明

### 本工具用于golang数据库访问的service生成
- 支持gorm和sqlx两种数据库访问方式。

### 基于http方式进行交互

- 发起代码文件生成
```shell
curl -X POST -H "Content-Type:application/json" http://localhost:7000/gen -d @req.json
```

- req.json  参数内容:
```json
{
    "db_code": "pg1",
    "table_name": "users",
    "object_name": "User",
    "package_base_path": "gorm-gen/internal_sqlx",
    "generate_path": "D:\\workspace-go\\src\\gorm-gen\\internal_sqlx",
    "generate_type":"gorm", 
    "create_file_types": ["model","service"],
    "remark": "用户信息"
}
```


### 基本配置文件app.json
```json
{
    "debug":true,
    "server":{
        "port":7000
    },
    "DBConfigs":{
        "pg1": {
            "dbType":"pg",
            "host":"192.168.19.161",
            "port":9999,
            "dbName":"testdb",
            "user":"postgres",
            "password":"postgres123"
        },
        "mysql1": {
            "dbType":"mysql",
            "host":"192.168.19.141",
            "port":3306,
            "dbName":"testdb",
            "user":"root",
            "password":"root123"
        }
    },
    "fieldTypeMap":{
        "pg":{
            "integer": "int",
            "bigint": "int64",
            "serial": "int64",
            "character varying": "string",
            "character": "string",
            "text":  "string",
            "boolean": "bool",
            "decimal": "float64",
            "numeric": "float64",
            "timestamp without time zone":  "time.Time",
            "timestamp with time zone": "time.Time"
        },
        "mysql":{
            "int": "int",
            "tinyint": "int8",
            "bigint": "int64",
            "serial": "int64",
            "char": "string",
            "varchar": "string",
            "text": "string",
            "longtext": "string",
            "bool": "bool",
            "decimal": "float64",
            "numeric": "float64",
            "float":"float32",
            "double":"float64",
            "timestamp": "time.Time",
            "datetime": "time.Time"
        }
    }
}
```