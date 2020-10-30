package utils

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
)

func GetDatabaseConfig(host string, user string, password string, dbname string) (*sql.DB, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&", user, password, host, dbname)
	val := url.Values{}
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	fmt.Println(connection)
	dbConn, err := sql.Open(`postgres`, dsn)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return dbConn, nil
}
