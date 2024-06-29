package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "repluser:replpassword@tcp(localhost:3307)/fileserver?charset=utf8")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to mysql, err: %s\n", err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		// store data in records
		err := rows.Scan(scanArgs...)
		checkErr(err)
		record := make(map[string]interface{})
		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
