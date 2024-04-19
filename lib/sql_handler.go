package lib

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SQLHandler ...
type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

func GenerateDsn() string {
	apiRevision := os.Getenv("API_REVISION")
	var dsn string

	if apiRevision == "release" {
		dsn = os.Getenv("DSN")
	} else {
		// user := os.Getenv("DB_USERNAME")
		// pass := os.Getenv("DB_PASSWORD")
		// host := os.Getenv("DB_HOST")
		// port := os.Getenv("DB_PORT")
		// dbName := os.Getenv("DB_DATABASE")

		// dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)
		dsn = "postgres://postgres:postgres@db:5432/shisha?sslmode=disable"
	}

	return dsn
}

// DBOpen は DB connectionを張る。
func DBOpen() {
	dbConn = NewSQLHandler()
}

// DBClose は DB connectionを張る。
func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	var db *gorm.DB
	var err error

	dsn := GenerateDsn()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	//コネクションプールの最大接続数を設定。
	sqlDB.SetMaxIdleConns(100)
	//接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
	sqlDB.SetMaxOpenConns(100)
	//接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	sqlHandler := new(SQLHandler)
	db.Logger.LogMode(4)
	sqlHandler.DB = db

	return sqlHandler
}

// GetDBConn ...
func GetDBConn() *SQLHandler {
	return dbConn
}

// BeginTransaction ...
func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

// Rollback ...
func RollBack() {
	dbConn.DB.Rollback()
}
