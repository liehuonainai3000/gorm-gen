package utils

import (
	"log"

	"github.com/spf13/viper"
)

var Global = new(GlobalConf)

func InitGlobalConfig() {

	v := viper.New()
	v.AutomaticEnv()       //绑定环境变量
	v.SetConfigName("app") // 指定配置文件
	v.SetConfigType("json")
	v.AddConfigPath(".")                         // 指定查找配置文件的路径
	v.AddConfigPath("./config")                  // 指定查找配置文件的路径
	v.AddConfigPath("../")                       // 指定查找配置文件的路径
	v.AddConfigPath(v.GetString("TRANS_CONFIG")) // 指定查找配置文件的路径

	// fmt.Println("可以设置环境变量 TRANS_CONFIG 来指定配置文件目录", v.GetString("TRANS_CONFIG"))

	err := v.ReadInConfig() // 读取配置信息
	if err != nil {         // 读取配置信息失败
		log.Fatalf("加载交易配置文件[%s]失败 %v", "app.json", err)
		panic(err)
	}

	v.Unmarshal(Global)

}
