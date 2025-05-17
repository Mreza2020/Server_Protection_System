package Server_Protection_System

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// # Run_Server_Sql
//
// It is for starting the mysql server.
//
// # Returns:
//
// - ch chan string
//
// - database, err := Run_Server_Sql(ch)
func Run_Server_Sql(ch chan string) (*sql.DB, error) {
	if Env_password_Loaded_string("dbPassword") != "" {
		var err error

		db, err := sql.Open("mysql", Env_password_Loaded_string("dbPassword"))
		if err != nil {
			log.Printf("not open mysql server : %v ", err)
			return nil, err
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("not ping :%v", err)
			return nil, err

		}
		ch <- "Server started successfully."
		return db, nil

	}

	return nil, nil
}
