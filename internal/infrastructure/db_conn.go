package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/masato-MacbookPro/go-chat/internal/app/httpserver/config"
)

const driverName = "mysql"

func NewMySQLConnector(mysqlInfo *config.MySQLInfo) (*sql.DB, error) {
	dsn := mysqlConnInfo(mysqlInfo)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func mysqlConnInfo(mysqlInfo *config.MySQLInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true&loc=Local",
		mysqlInfo.MySqlUser,
		mysqlInfo.MySqlPassWrord,
		mysqlInfo.MysqlAddr,
		mysqlInfo.MysqlDBName,
	)

	return dataSourceName
}
