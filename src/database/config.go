package database

import (
	"fmt"
	"sync"
	"sync/atomic"

	// _ "github.com/lib/pq" //Postgres Driver
	// _ "github.com/golang-migrate/migrate/database/postgres" //Postgres Driver

	_ "github.com/jinzhu/gorm/dialects/postgres" //Dialect ORM Postgress
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "etl-golang"
)

var dbInstance *gorm.DB
var atomicinz uint64
var lock = &sync.Mutex{}

//DrivePg -
type DrivePg struct {
	conn string
}

//Connect - Configuração de banco de dados
func Connect() *gorm.DB {

	if atomic.LoadUint64(&atomicinz) == 1 {

		return dbInstance
	}

	lock.Lock()
	defer lock.Unlock()

	if atomicinz == 0 {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		dbInstance = db
		atomic.StoreUint64(&atomicinz, 1)

		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully connected!")
	}

	return dbInstance
}
