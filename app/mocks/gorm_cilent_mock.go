package mocks

import (
	"io"
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	gin "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	// สร้าง mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	// ใช้ PostgreSQL driver ของ GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	// คืนค่า GORM DB พร้อมกับ mock
	return gormDB, mock
}

func Setup() {
	gin.SetMode(gin.ReleaseMode)
	env := os.Getenv("RUN_TEST_MODE")
	logrus.Infof("os.Args = %+v \n", env)

	if env == "prod" {
		logrus.Warn("Disable Logging...")
		logrus.SetOutput(io.Discard)
	}
}
