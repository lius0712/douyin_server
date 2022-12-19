package viper

import (
	"fmt"
	"testing"
)

func TestViperConfig(t *testing.T) {
	v := ConfigInit("mysqlConfig").Viper
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		v.GetString("Mysql.User"),
		v.GetString("Mysql.Password"),
		v.GetString("Mysql.Host"),
		v.GetInt("Mysql.Port"),
		v.GetString("Mysql.DBName"),
		v.GetString("Mysql.CharSet"),
		v.GetBool("Mysql.parseTime"),
		v.GetString("Mysql.loc"),
	)

	e := ConfigInit("apiConfig").Viper
	etcd := fmt.Sprintf("%s:%s", e.GetString("Etcd.Host"), e.GetString("Etcd.Port"))
	fmt.Println(dsn)
	fmt.Println(etcd)

	tencentConfig := ConfigInit("tencentCos").Viper
	CosUrl := tencentConfig.GetString("tencent.Url")
	SecretID := tencentConfig.GetString("tencent.SecretID")
	SecretKey := tencentConfig.GetString("tencent.SecretKey")
	fmt.Println(CosUrl)
	fmt.Println(SecretID)
	fmt.Println(SecretKey)

}
