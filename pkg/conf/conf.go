package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var ProjectName string = "local_life"

func InitFromFile(fname, ftype string, fpath ...string) {
	viper.SetConfigName(fname)
	viper.SetConfigType(ftype)
	for _, path := range fpath {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found.")
			return
		}
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func Init() {
	InitFromFile("conf", "yaml", "conf", ".")
}
