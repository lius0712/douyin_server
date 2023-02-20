package db

import (
	"github.com/lius0712/douyin_server/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	//m := DB.Migrator()
	//if m.HasTable(&dbUser.User{}) {
	//	return
	//}
	//if m.HasTable(&dbVideo.Video{}) {
	//	return
	//}
	//if err = m.CreateTable(&dbUser.User{}); err != nil {
	//	panic(err)
	//}
	//if err = m.CreateTable(&dbVideo.Video{}); err != nil {
	//	panic(err)
	//}
}
