package repository

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/pkg"
	"bytes"
	"fmt"
	"math"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type PaginationModel struct {
	Limit     int
	Offset    int
	Search    string
	SortField string
	SortValue string
	Field     map[string]any
	Dest      any
}

func getField(field map[string]interface{}) string {
	b := new(bytes.Buffer)
	index := 0
	for key := range field {
		// fmt.Println("key", key)
		if index > 0 {
			fmt.Fprintf(b, ",%s", strings.ToLower(key))
		} else {
			fmt.Fprintf(b, "%s", strings.ToLower(key))
		}

		index += 1
	}
	return b.String()

}

type BaseRepositoryInterface interface {
	ClientDb() *gorm.DB
	FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error
	Pagination(p PaginationModel, query interface{}, args ...interface{}) (result interface{}, Error error)
	TotalPage(model interface{}, pageSize int) (int64, error)
	Save(tx *gorm.DB, model interface{}) error
	Updates(tx *gorm.DB, model any, values any) error
	Delete(tx *gorm.DB, model any, body any) error
	Count(model any, conditions *map[string]any) (int64, error)
}

type BaseRepository struct {
	db *gorm.DB
}

func BaseRepositoryInit(db *gorm.DB) *BaseRepository {

	return &BaseRepository{
		db: db,
	}
}

func (b BaseRepository) ClientDb() *gorm.DB {
	return b.db
}

func (b BaseRepository) FindOne(tx *gorm.DB, model interface{}, query interface{}, args ...interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}
	var err error
	if query == nil || args == nil {
		log.Error("Got an error when findOne required query")
		pkg.PanicException(constant.RequiredQuery)
	}

	err = db.Where(query, args...).First(model).Error

	if err != nil {
		log.Error("Got an error when findOne Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) Pagination(p PaginationModel, query interface{}, args ...interface{}) (result interface{}, Error error) {
	var err error
	order := fmt.Sprintf("%s %s", p.SortField, strings.ToUpper(p.SortValue))
	fields := getField(p.Field)
	var db *gorm.DB
	db = b.db

	if fields != "" {
		db = b.db.Select(fields)
	}

	if query != nil {
		db = b.db.Where(query, args...)
	}

	db = db.Order(order).Offset(p.Offset).Limit(p.Limit).Find(p.Dest)

	if db.Error != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return p.Dest, nil
}

func (b BaseRepository) TotalPage(model interface{}, pageSize int) (int64, error) {
	var count int64
	err := b.db.Model(model).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(pageSize)))
	return totalPage, err
}

func (b BaseRepository) Save(tx *gorm.DB, model interface{}) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	var err = db.Save(model).Error
	if err != nil {
		log.Error("Got an error when save Error: ", err)
		return err
	}
	return nil
}

func (b BaseRepository) Updates(tx *gorm.DB, model any, values any) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	err := db.Model(model).Updates(values).Error

	if err != nil {
		log.Error("Got an error when save Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) Delete(tx *gorm.DB, model any, body any) error {
	db := b.db

	if tx != nil {
		db = tx
	}

	err := db.Model(model).Delete(body).Error

	if err != nil {
		log.Error("Got an error when save Error: ", err)
		return err
	}

	return nil
}

func (b BaseRepository) Count(model any, conditions *map[string]any) (int64, error) {
	var count int64
	err := b.db.Model(model).Count(&count).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return count, err
	}

	return count, err
}
