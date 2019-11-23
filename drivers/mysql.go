package drivers

import (
	"database/sql"
	"fmt"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// MySQL wraps a mysql db instance
type MySQL struct {
	conn *sql.DB
}

// NewMySQL create MySQL struct
func NewMySQL() (*MySQL, error) {
	server := viper.GetString("server") // old: rootCmd.Flags().GetString("server")
	user := viper.GetString("user")
	pass := viper.GetString("pass")
	db := viper.GetString("db")

	conn, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=UTF8", user, pass, server, db))
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &MySQL{
		conn: conn,
	}, nil
}
