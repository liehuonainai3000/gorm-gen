
## 使用说明

### 本工具用于golang的数据库访问代码生成
- 支持gorm和sqlx两种数据库访问方式。
- 支持的数据库类型包括：postgresql，mysql
- 可以同时配置多个数据源

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

参数说明：

|参数名|说明|
|-|-|
|db_code|用于指定不同给的数据库配置，对应app.json中DBConfigs的key值|
|table_name|要生成代码的数据库表名|
|object_name|生成模型的Struct名|
|package_base_path|生成代码的包基础路径
|generate_path|生成代码的存放路径|
|generate_type|生成数据库访问类型，参数值：gorm、sqlx|
|create_file_types|用于指定代码生成类型，gorm包括：model/service，sqlx包括：dao/entity/do/service。如果不指定，则全部生成。|
|remark|生成对象描述|


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
参数说明：
 - dbType 代表所用的数据库类型，目前支持postgres（pg）和mysql
 - fieldTypeMap用于指定数据库字段类型与struct字段类型的对应关系，可根据需要自行修改或添加。

### 扩展
可通过实现MetaQueryer接口，并通过gen.RegisteMetaQuery(dbCode string, metaQuery MetaQueryer)注册方法,实现对其他数据库的支持。