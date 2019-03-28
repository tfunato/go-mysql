package main

import (
    //"database/sql"
    "github.com/tfunato/go-mysql/pkg/service/mysql"
    //"fmt"
)

func main() {
    mysql.Select("select id, sub, user_id, first_name, created_at from account where user_id = ?", []string{"2100000003"})
    /*
    db :=  mysql.DbConn()
    rows, err := db.Query("select id, sub, user_id, first_name, created_at from account where user_id = ?", "2100000003")
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
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Println(columns[i], ": ", value)
        }
        fmt.Println("-----------------------------------")
    }

    defer db.Close()
    */
}
