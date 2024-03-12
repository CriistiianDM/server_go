/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	database *sql.DB
	once sync.Once
)

/**
* Connect to the database
*/
func Connect(wg *sync.WaitGroup) {
	// Realiza el llamado una unica vez.
	once.Do(func() {
		// Initialize the environment variables
		getVarEnv()
		// Get the environment variables
		var (
			host = os.Getenv("DB_HOST")
			port = os.Getenv("DB_PORT")
			user = os.Getenv("DB_USER")
			password = os.Getenv("DB_PASSWORD")
			dbname = os.Getenv("DB_NAME")
		)

		if isNotEmptyEnv(host,port,user,password,dbname) {
			var (
				connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)
				maxAttempts = 5
				db *sql.DB
				err error
			)

			for i := 1; i <= maxAttempts; i++ {
				db, err = sql.Open("postgres", connStr)
				if err == nil {
					err = db.Ping()
					if err == nil {
						break
					}
				}
				time.Sleep(3000 * time.Millisecond)
			}

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Conexión exitosa a la base de datos PostgreSQL")
			database = db
		}
	})
	defer wg.Done()
}

/**
* Get the environment variables
*/
func getVarEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env file, using environment variables")
	}
}

/**
* Get the database connection
*/
func GetConnection() *sql.DB {
	return database
}

/**
  * Validate if propierties of file .env have info
*/
func isNotEmptyEnv(args ...string) bool {
	isEmpty := false

	if len(args) > 0 {
		for _, env := range args {
			if env == "" {
				isEmpty = false
				break
			} else {isEmpty = true}
		}
	}
	
	return isEmpty
}