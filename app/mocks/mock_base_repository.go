package mocks

import (
	"agnos-assignment/app/repository"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

type MockBaseRepository struct {
	mock.Mock
}

func (m *MockBaseRepository) ClientDb() *gorm.DB {
	fmt.Println("ClientDb")
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockBaseRepository) Save(tx *gorm.DB, model interface{}) error {
	args := m.Called(tx, model)
	return args.Error(0)
}

func (m *MockBaseRepository) FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	argsIsExits := m.Called(tx, model, query, args)
	return argsIsExits.Error(0)
}

func (m *MockBaseRepository) Transaction(fn func(tx *gorm.DB) error) error {
	args := m.Called(fn)

	if fn != nil {
		tx := new(gorm.DB)
		return fn(tx)
	}

	return args.Error(0)
}

func (m *MockBaseRepository) Pagination(p repository.PaginationModel, query interface{}, args ...interface{}) (interface{}, error) {
	argsPagination := m.Called(p, query, args)
	return argsPagination.Get(0), argsPagination.Error(1)
}

func (m *MockBaseRepository) TotalPage(model interface{}, pageSize int) (int64, error) {
	argsIsExits := m.Called(model, pageSize)
	return argsIsExits.Get(0).(int64), argsIsExits.Error(1)
}
