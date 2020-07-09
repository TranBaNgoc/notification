package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



func GetUserTokens(id string) []string {
	db, err := sql.Open("mysql", "5EIxJdtsCE:mX0MKS9YYG@tcp(remotemysql.com:3306)/5EIxJdtsCE")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	tokens := []string{}
	results, err := db.Query("select device_token from user where device_token != \"\" and device_token is not null and id = " + id)
	for results.Next() {
		var token string
		err = results.Scan(&token)
		if err != nil {
			log.Println(err)
		}
		tokens = append(tokens, token)
	}
	return tokens
}
