package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
}

func ConfigInit(yamlFileName string) Config {
	//v := viper.New()
	config := Config{Viper: viper.New()}

	v := config.Viper

	v.SetConfigName(yamlFileName)
	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")
	v.AddConfigPath("../../../config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}
	return config
}
