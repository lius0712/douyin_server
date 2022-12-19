package viper

import (
	"fmt"
	"testing"
)

func TestViperConfig(t *testing.T) {
	v := ConfigInit("mysqlConfig").viper
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
	fmt.Println(dsn)
}
