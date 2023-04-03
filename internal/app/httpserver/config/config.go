package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MySQLInfo
}

type HTTPInfo struct {
	Addr string
}

type MySQLInfo struct {
	MySqlUser      string
	MySqlPassWrord string
	MysqlAddr      string
	MysqlDBName    string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	addr := ":" + os.Getenv("PORT")
	httpInfo := &HTTPInfo{
		Addr: addr,
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassWord := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	mysqlInfo := &MySQLInfo{
		MySqlUser:      mysqlUser,
		MySqlPassWrord: mysqlPassWord,
		MysqlAddr:      mysqlAddr,
		MysqlDBName:    mysqlDBName,
	}

	conf := &AppConfig{
		HTTPInfo:  httpInfo,
		MySQLInfo: mysqlInfo,
	}

	return conf
}
