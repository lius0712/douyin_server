package constants

import (
	"fmt"
	"github.com/lius0712/douyin_server/pkg/viper"
)

var (
	mysqlConfig    = viper.ConfigInit("mysqlConfig").Viper
	apiConfig      = viper.ConfigInit("apiConfig").Viper
	tencentConfig  = viper.ConfigInit("tencentCos").Viper
	userConfig     = viper.ConfigInit("userConfig").Viper
	publishConfig  = viper.ConfigInit("publishConfig").Viper
	feedConfig     = viper.ConfigInit("feedConfig").Viper
	relationConfig = viper.ConfigInit("relationConfig").Viper
	jwtConfig      = viper.ConfigInit("jwtConfig").Viper
	logConfig      = viper.ConfigInit("logConfig").Viper
)

var (
	Mode            = apiConfig.GetString("Mode.Name")
	MySQLDefaultDSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.GetString("Mysql.User"),
		mysqlConfig.GetString("Mysql.Password"),
		mysqlConfig.GetString("Mysql.Host"),
		mysqlConfig.GetInt("Mysql.Port"),
		mysqlConfig.GetString("Mysql.DBName"),
		mysqlConfig.GetString("Mysql.CharSet"),
		mysqlConfig.GetBool("Mysql.parseTime"),
		mysqlConfig.GetString("Mysql.loc"),
	)
	EtcdAddress           = fmt.Sprintf("%s:%s", apiConfig.GetString("Etcd.Host"), apiConfig.GetString("Etcd.Port"))
	ClientAddress         = fmt.Sprintf("%s:%s", apiConfig.GetString("Server.Address"), apiConfig.GetString("Server.Port"))
	ClientName            = apiConfig.GetString("Server.Name")
	UserServerAddress     = fmt.Sprintf("%s:%s", userConfig.GetString("Server.Address"), userConfig.GetString("Server.Port"))
	UserServerName        = userConfig.GetString("Server.Name")
	PublishServerAddress  = fmt.Sprintf("%s:%s", publishConfig.GetString("Server.Address"), publishConfig.GetString("Server.Port"))
	PublishServerName     = publishConfig.GetString("Server.Name")
	FeedServerAddress     = fmt.Sprintf("%s:%s", feedConfig.GetString("Server.Address"), feedConfig.GetString("Server.Port"))
	FeedServerName        = feedConfig.GetString("Server.Name")
	RelationServerAddress = fmt.Sprintf("%s:%s", relationConfig.GetString("Server.Address"), relationConfig.GetString("Server.Port"))
	RelationServerName    = relationConfig.GetString("Server.Name")
	JwtKey                = jwtConfig.GetString("JWT.signingKey")
	CosUrl                = tencentConfig.GetString("tencent.Url")
	SecretID              = tencentConfig.GetString("tencent.SecretID")
	SecretKey             = tencentConfig.GetString("tencent.SecretKey")
	LogFileName           = logConfig.GetString("log.filename")
	LogLevel              = logConfig.GetString("log.level")
	LogMaxSize            = logConfig.GetInt("log.max_size")
	LogMaxAge             = logConfig.GetInt("log.max_age")
	LogMaxBackups         = logConfig.GetInt("log.max_backups")
)

const (
	UserTableName     = "t_user"
	VideoTableName    = "t_video"
	RelationTableName = "t_relation"
)
