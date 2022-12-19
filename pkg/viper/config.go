package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	viper *viper.Viper
}

func ConfigInit(yamlFileName string) Config {
	//v := viper.New()
	config := Config{viper: viper.New()}

	v := config.viper

	v.SetConfigName(yamlFileName)
	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")
	v.AddConfigPath("../../../config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config: %s", err))
	}

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
	//	v.GetString("Mysql.User"),
	//	v.GetString("Mysql.Password"),
	//	v.GetString("Mysql.Host"),
	//	v.GetInt("Mysql.Port"),
	//	v.GetString("Mysql.DBName"),
	//	v.GetString("Mysql.CharSet"),
	//	v.GetBool("Mysql.parseTime"),
	//	v.GetString("Mysql.loc"),
	//)
	//fmt.Println(dsn)
	return config
}
