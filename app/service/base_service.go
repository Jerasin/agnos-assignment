package service

import (
	"agnos-assignment/app/constant"
	"agnos-assignment/app/pkg"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/utils"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func DbHandleSelectField(field any) map[string]interface{} {
	fields := reflect.TypeOf(field)
	result := make(map[string]interface{})
	for i := 0; i < fields.NumField(); i++ {
		// Get the field
		field := fields.Field(i)

		// Get the json tag value
		jsonTag := field.Tag.Get("json")

		// Print the json tag value
		logrus.Infof("Field %d: %s\n", i+1, jsonTag)
		result[jsonTag] = ""
	}

	return result
}

func ConvertMapToQuery(conditions map[string]any, db *gorm.DB) *gorm.DB {
	for field, value := range conditions {
		if value != "" {
			db = db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}

	return db
}

type BaseServiceInterface interface {
	Create(c *gin.Context, condition map[string]any, model any)
	Pagination(c *gin.Context, field any, model any, resModel any)
	GetDetail(c *gin.Context, model any, resModel any)
	Updates(c *gin.Context, model any, body any)
	IsExist(tx *gorm.DB, c *gin.Context, condition map[string]any, model any)
	Delete(c *gin.Context, model any)
}

type BaseService struct {
	BaseRepository repository.BaseRepositoryInterface
}

func BaseServiceInit(baseRepo repository.BaseRepositoryInterface) *BaseService {
	return &BaseService{
		BaseRepository: baseRepo,
	}
}

func (b *BaseService) Create(c *gin.Context, condition map[string]any, model any) {
	defer pkg.PanicHandler(c)

	b.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error
		db := ConvertMapToQuery(condition, tx)

		err = db.Find(model).Error
		if err != nil {
			pkg.PanicDatabaseException(err, c, nil)
		}

		logrus.Infof("model = %+v \n", model)
		logrus.Infof("model = %T \n", model)
		fmt.Println("CheckIdIsDefault", utils.CheckIdExist(model))

		if !utils.CheckIdExist(model) {
			pkg.PanicException(constant.DataIsExit)
		}

		err = b.BaseRepository.Save(tx, model)
		if err != nil {
			pkg.PanicDatabaseException(err, c, nil)
		}

		fmt.Println("After Save", model)

		return nil
	})

	c.JSON(http.StatusOK, pkg.CreateResponse())
}

func (b *BaseService) Pagination(c *gin.Context, field any, model any, resModel any) {
	defer pkg.PanicHandler(c)
	params := pkg.CreatePagination(c)
	fields := DbHandleSelectField(field)

	logrus.Infof("fields = %+v \n", fields)

	paginationModel := repository.PaginationModel{
		Limit:     params.Limit,
		Offset:    params.Offset,
		Search:    params.Search,
		SortField: params.SortField,
		SortValue: params.SortValue,
		Field:     fields,
		Dest:      model,
	}

	logrus.Infof("paginationModel = %+v \n", paginationModel)

	data, err := b.BaseRepository.Pagination(paginationModel, nil)
	if err != nil {
		logrus.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	fmt.Println("data", data)

	totalPage, err := b.BaseRepository.TotalPage(&model, params.PageSize)
	if err != nil {
		logrus.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	total, err := b.BaseRepository.Count(&model, nil)
	if err != nil {
		logrus.Error("Count Data Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	pkg.ModelDump(resModel, data)
	c.JSON(http.StatusOK, pkg.BuildPaginationResponse(constant.Success, resModel, total, totalPage, params.Page, params.PageSize))
}

func (b *BaseService) GetDetail(c *gin.Context, model any, resModel any) {
	defer pkg.PanicHandler(c)

	ID := c.Param("ID")

	conditions := map[string]any{
		"id": ID,
	}

	err := b.BaseRepository.FindOne(nil, model, &conditions)

	logrus.Infof("model ZZZ = %+v \n", model)

	if err != nil {
		pkg.PanicException(constant.DataNotFound)
	}

	pkg.ModelDump(resModel, model)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, resModel))
}

func (b *BaseService) Updates(c *gin.Context, model any, body any) {
	defer pkg.PanicHandler(c)

	ID := c.Param("ID")

	logrus.Infof("ID = %+v \n", ID)

	b.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {
		var err error

		conditions := map[string]any{
			"id": ID,
		}

		err = b.BaseRepository.FindOne(tx, model, &conditions)
		if err != nil {
			fmt.Println("error", err)
			pkg.PanicDatabaseException(err, c, nil)
		}

		err = b.BaseRepository.Updates(tx, model, body)
		if err != nil {
			pkg.PanicDatabaseException(err, c, nil)
		}

		return nil
	})

	c.JSON(http.StatusOK, pkg.UpdateResponse())

}

func (b *BaseService) IsExist(tx *gorm.DB, c *gin.Context, condition map[string]any, model any) {
	defer pkg.PanicHandler(c)

	db := b.BaseRepository.ClientDb()

	if tx != nil {
		db = ConvertMapToQuery(condition, tx)
	} else {
		db = ConvertMapToQuery(condition, db)
	}

	result := db.First(model)
	err := result.Error

	if err != nil {
		logrus.Infof("IsExist error = %+v \n", err)
		pkg.PanicDatabaseException(err, c, gorm.ErrRecordNotFound)
	}

	if result.RowsAffected > 0 {
		pkg.PanicException(constant.DataIsExit)
	}

}

func (b *BaseService) Delete(c *gin.Context, model any) {
	defer pkg.PanicHandler(c)

	ID := c.Param("ID")

	b.BaseRepository.ClientDb().Transaction(func(tx *gorm.DB) error {

		conditions := map[string]any{
			"id": ID,
		}

		err := b.BaseRepository.FindOne(tx, model, &conditions)
		if err != nil {
			pkg.PanicDatabaseException(err, c, nil)
		}

		logrus.Infof("model = %+v \n", model)

		b.BaseRepository.Delete(tx, model, map[string]any{"id": ID})

		return nil
	})
}
