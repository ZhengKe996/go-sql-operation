package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"'`
	MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func main() {
	viper.SetConfigFile("./config.yaml")         // 指定配置文件路径
	if err := viper.ReadInConfig(); err != nil { // 查找并读取配置文件, 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig() // 监控配置文件的变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("viper unmarshal failed, err:%v\n", err)
		return
	}

	fmt.Printf("C:%#v\n", c)

	//r := gin.Default()
	//r.GET("/version", func(context *gin.Context) {
	//	context.String(http.StatusOK, viper.GetString("version"))
	//})
	//r.Run()
}
