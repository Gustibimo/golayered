package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
driver package will have only the func to connect to db
*/

type MySQLConfig struct {
	Host     string
	Port     string
	Db       string
	User     string
	Password string
}

// ConnectToMySQL method takes mysql defined config and connect to mysql
func ConnectToMysql(conf MySQLConfig) (*sql.DB, error) {
	connection := fmt.Sprintf("%v:%v@tcp(%v%:v)/%v", conf.User, conf.Password,
		conf.Db, conf.Host, conf.Port)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	return db, nil
}
