package mysql

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func DbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "password"
    dbName := "local"
    dbHost := "127.0.0.1"
    db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName)
    if err != nil {
        panic(err.Error())
    }
    return
}

//func Select(sqlStr string, params []string) []map[string]string {
func Select(sqlStr string, params []string) {
    db := DbConn()
    rows, err := db.Query(sqlStr, params)
    if err != nil {
        panic(err.Error())
    }
    columns, err := rows.Columns() // カラム名を取得
    if err != nil {
        panic(err.Error())
    }
    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error())
        }

        var value string
        for i, col := range values {
            value = string(col)
            fmt.Println(columns[i], ": ", value)
        }
    }

    defer db.Close()
}
