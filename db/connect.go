/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2023
 * @license: GPL-3.0
 */
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

/**
  * Connect to the database
*/
func Connect() {
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

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexión exitosa a la base de datos PostgreSQL")
}

/**
  * Get the environment variables
*/
func getVarEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}

/**
  * Get the database connection
*/
func GetConnection() *sql.DB {
	return db
}