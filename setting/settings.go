package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*MySqlConfig `mapstructure:"log"`
	*RedisConfig `mapstructure:"redis"`
	*WebConfig   `mapstructure:"web"`
}
type MySqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
type WebConfig struct {
	Name       string `mapstructure:"name"`
	CommendIds string `mapstructure:"commend_ids"`
	TxtUrl     string `mapstructure:"txt_url"`
	ImgUrl     string `mapstructure:"img_url"`
	SearchNum  string `mapstructure:"search_num"`
}

func Init() (err error) {
	viper.SetConfigName("config") //配置指定文件名称（无需后缀）
	viper.SetConfigType("yaml")   //指定配置文件类型
	viper.AddConfigPath(".")      //指定查找配置文件的路径 （相对路径）
	err = viper.ReadInConfig()
	if err != nil {
		//读取配置信息失败
		fmt.Println("读取配置信息失败", err)
		return
	}
	//把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("vconst.Unmarshal failed,err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了。。。")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("vconst.Unmarshal failed,err:%v\n", err)
		}
	})
	return
}
