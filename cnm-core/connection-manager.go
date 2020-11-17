package core

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DatabaseManager interface {
	openConnection() ( *gorm.DB, error )
	getDB() ( *gorm.DB, error )
}


type Connection struct {
 db *gorm.DB
 error
}

var Conn= Connection{};

func (s *Connection ) openConnection()  ( *gorm.DB, error ) {
	//TODO get this from docker environment variables
	dataSource:="host=192.168.1.111 user=postgres password=root@123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := s.InitDB(dataSource)

	if err!=nil {
		Conn.error=err;
	}
	Conn.db=db;

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, err
}

func (s *Connection ) InitDB(dataSource string) ( *gorm.DB, error ){
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,         // Disable color
		},
	)
	var err error
	fmt.Println("Initializing database connection . . .")
   var db *gorm.DB
	for i := 0; i < 60; i++ {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dataSource,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage

		}), &gorm.Config{Logger: newLogger,})

		if err != nil {
			fmt.Printf("Unable to Open DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 2)
		}  else {
			err = nil
			break
		}

	}

	return db, err
}

func (s *Connection ) getDB() ( *gorm.DB, error ) {
	return Conn.db,Conn.error;
}

func   GetDB() ( *gorm.DB, error ) {
	return Conn.db,Conn.error;
}

func NewDatabaseManager() ( *gorm.DB, error ) {

	if Conn.db==nil {
		db,err:=Conn.openConnection();
		return db,err
	}
	return Conn.db,Conn.error;
}

