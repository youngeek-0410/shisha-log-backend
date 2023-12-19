package lib

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQLHandler ...
type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

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
	// user := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_DATABASE")
	// fmt.Print(user, password, host, port)

	var db *gorm.DB
	var err error

	dsn := "root:@tcp(db:3306)/shisha-log?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
