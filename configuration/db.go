package configuration

import (
	"ams/services"
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB db object.
// We are making this public, so that it can be passed
// to other package.
// This is a global variables, so that we don't have to create
// a new connection everytime GetDB is called.
var DB *gorm.DB

// GetConnectionString takes as input the program's environment and returns
// the connection string for postgres database.
func getDatabaseConnectionString() string {
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")

	// Will check is any variable is not empty, is so it will panic with error message.
	if dbuser == "" || dbpass == "" || dbname == "" || dbhost == "" || dbport == "" {
		err := errors.New("Database environment variable is not good, its empty")
		services.PanicIfError(err)
	}

	connectionString := "user=" + dbuser + " password=" + dbpass + " dbname=" + dbname + " port=" + dbport + " host=" + dbhost + " sslmode=disable"

	return connectionString
}

// openDB will make connection with postgres database.
func openDB() {
	var err error

	// https://github.com/go-gorm/postgres
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  getDatabaseConnectionString(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	services.PanicIfError(err)
}

// GetDB will return orm object for databae connection.
// It will check for open database connection if not found
// will make one.
// We are checking, because we don't want to setup
// multiple connection.
func GetDB() *gorm.DB {
	if DB == nil {
		openDB()
	}
	return DB
}
