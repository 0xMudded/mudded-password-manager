package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var vp = viper.New()

func Load() {
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func AddConfig(key, value string) {
	vp.Set(key, value)
	vp.WriteConfig()
}

func GetViper() *viper.Viper {
	return vp
}
