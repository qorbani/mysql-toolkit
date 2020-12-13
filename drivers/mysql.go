package drivers

import (
	"database/sql"
	"encoding/json"
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

// Query execute SELECT sql statement
func (sql *MySQL) Query(query string) (string, error) {
	stmt, err := sql.conn.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	dataRaw := make([]map[string]interface{}, 0)

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return "", err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]

			b, ok := v.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}

		dataRaw = append(dataRaw, entry)
	}

	dataJSON, err := json.Marshal(dataRaw)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}

// // Load use LOAD DATA command to load file to specified table.
// func (sql *MySQL) Load(filePath string) error {
// 	mysql.RegisterLocalFile(filePath)
//	// https://dev.mysql.com/doc/refman/8.0/en/load-data.html
// 	err := sql.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE foo")
// 	return err
// }
