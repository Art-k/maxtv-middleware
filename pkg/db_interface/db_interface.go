package db_interface

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	. "maxtv_middleware/pkg/common"
	"os"
	"time"
)

func InitDatabase() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			//LogLevel:      logger.Info, // Log level
			LogLevel: logger.Error, // Log level
			Colorful: true,         // Disable color
		},
	)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	"developer", "12345678", "sql.maxtvmedia.com", "3306", "maxtv_cms_live")

	//DATABASE_NAME = maxtv_cms_live
	//DATABASE_USER = hmuser
	//DATABASE_PASS = xyKjg9SlxpwKgHK447VcgCt9A9e+kAM6
	//DATABASE_HOST = sql.maxtvmedia.com
	//DATABASE_PORT = 3306

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASS"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

	DB, DBErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if DBErr != nil {
		log.Panic("[DATABASE INIT] ", DBErr)
		panic("ERROR failed to connect database ")
	}

}
