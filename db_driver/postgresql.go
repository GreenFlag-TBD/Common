package db_driver

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"reflect"
)

const scan_tag = "db"

type PostgresSQLConn struct {
	*pgx.Conn
}

func Connect(dsn string) *PostgresSQLConn {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &PostgresSQLConn{conn}
}

func ScanStruct(rows pgx.Row, dest interface{}) error {
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dest must be a pointer to a struct")
	}
	destElem := destVal.Elem()
	destType := destElem.Type()
	var ptrs []interface{}
	for i := 0; i < destType.NumField(); i++ {
		field := destElem.Field(i)
		tag := destType.Field(i).Tag.Get(scan_tag)
		if tag == "" {
			continue
		}
		ptrs = append(ptrs, field.Addr().Interface())
	}

	if err := rows.Scan(ptrs...); err != nil {
		return err
	}
	return nil
}
