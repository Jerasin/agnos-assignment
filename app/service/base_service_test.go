package service

import (
	"agnos-assignment/app/mocks"
	"agnos-assignment/app/model"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/response"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestBaseService_Create(t *testing.T) {
	mocks.Setup()
	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		c         *gin.Context
		condition map[string]any
		model     any
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		statusCode int
	}{
		// TODO: Add test cases.
		{
			name: "BaseService Create Success",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()

					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectBegin()
					sqlMock.ExpectQuery(`SELECT .* FROM "hospitals".*`).
						WithArgs("โรงพยาบาลพระราม 9", "praram9").
						WillReturnRows(sqlMock.NewRows([]string{"id"}).AddRow(0))
					mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					return c
				}(),

				condition: map[string]any{
					"name_th": "โรงพยาบาลพระราม 9",
					"name_en": "praram9",
				},

				model: &model.Hospital{
					NameTh: "โรงพยาบาลพระราม 9",
					NameEn: "praram9",
				},
			},
			wantErr:    false,
			statusCode: http.StatusOK,
		},
		{
			name: "BaseService Create Failed",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()

					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectBegin()
					sqlMock.ExpectQuery(`SELECT .* FROM "hospitals".*`).
						WithArgs("โรงพยาบาลพระราม 9", "praram9").
						WillReturnRows(sqlMock.NewRows([]string{"id"}).AddRow(1))
					mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					return c
				}(),

				condition: map[string]any{
					"name_th": "โรงพยาบาลพระราม 9",
					"name_en": "praram9",
				},

				model: &model.Hospital{
					NameTh: "โรงพยาบาลพระราม 9",
					NameEn: "praram9",
				},
			},
			wantErr:    true,
			statusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered:", r)
				}
			}()

			b := &BaseService{
				BaseRepository: tt.fields.BaseRepository,
			}

			assert.NotPanics(t, func() {
				b.Create(tt.args.c, tt.args.condition, tt.args.model)
			})

			logrus.Infof("Writer Status = %+v", tt.args.c.Writer.Status())

			if tt.wantErr {

				if tt.statusCode == tt.args.c.Writer.Status() {
					assert.NotNil(t, tt.args.c)
					assert.Equal(t, tt.statusCode, tt.args.c.Writer.Status())
				} else {
					panic("error")
				}

			} else {
				assert.Equal(t, http.StatusOK, tt.args.c.Writer.Status())
			}

		})
	}
}

func TestBaseService_GetDetail(t *testing.T) {
	mocks.Setup()

	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		c        *gin.Context
		model    any
		resModel any
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		statusCode int
	}{
		// TODO: Add test cases.
		{
			name: "BaseService Get Detail Success",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockRepo.On("FindOne", mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Params = gin.Params{
						gin.Param{Key: "ID", Value: "123"},
					}
					return c
				}(),

				resModel: response.HospitalModel{
					Id:     1,
					NameTh: "โรงพยาบาลพระราม 9",
					NameEn: "praram9",
				},

				model: &model.Hospital{
					BaseModel: model.BaseModel{
						ID: 1,
					},
					NameTh: "โรงพยาบาลพระราม 9",
					NameEn: "praram9",
				},
			},
			wantErr:    false,
			statusCode: 200,
		},
		{
			name: "BaseService Get Detail Failed",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockRepo.On("FindOne", mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(errors.New("some error"))

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Params = gin.Params{
						gin.Param{Key: "ID", Value: "123"},
					}
					return c
				}(),

				resModel: response.HospitalModel{
					Id:     1,
					NameTh: "โรงพยาบาลพระราม 9",
					NameEn: "praram9",
				},

				model: &model.Hospital{
					BaseModel: model.BaseModel{
						ID: 1,
					},
					NameTh: "โรงพยาบาลพระราม 9",
					NameEn: "praram9",
				},
			},
			wantErr:    true,
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered:", r)
				}
			}()

			b := &BaseService{
				BaseRepository: tt.fields.BaseRepository,
			}

			assert.NotPanics(t, func() {
				b.GetDetail(tt.args.c, tt.args.model, tt.args.resModel)
			})

			logrus.Infof("Writer Status = %+v", tt.args.c.Writer.Status())

			if tt.wantErr {

				if tt.statusCode == tt.args.c.Writer.Status() {
					assert.NotNil(t, tt.args.c)
					assert.Equal(t, tt.statusCode, tt.args.c.Writer.Status())
				} else {
					panic("error")
				}

			} else {
				assert.Equal(t, http.StatusOK, tt.args.c.Writer.Status())
			}

		})
	}
}

func TestBaseService_Update(t *testing.T) {
	mocks.Setup()

	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		c     *gin.Context
		model any
		body  any
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		statusCode int
	}{
		// TODO: Add test cases.
		{
			name: "BaseService Update Success",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectBegin()

					mockRepo.On("FindOne", mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					mockRepo.On("Updates", mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Params = gin.Params{
						gin.Param{Key: "ID", Value: "123"},
					}
					return c
				}(),
			},
			wantErr:    false,
			statusCode: 200,
		},
		{
			name: "BaseService Update Failed",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectBegin()

					mockRepo.On("FindOne", mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(gorm.ErrRecordNotFound)

					mockRepo.On("Updates", mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Params = gin.Params{
						gin.Param{Key: "ID", Value: "123"},
					}
					return c
				}(),
			},
			wantErr:    true,
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseService{
				BaseRepository: tt.fields.BaseRepository,
			}

			assert.NotPanics(t, func() {
				b.Updates(tt.args.c, tt.args.model, tt.args.body)
			})

			logrus.Infof("Writer Status = %+v", tt.args.c.Writer.Status())

			if tt.wantErr {

				if tt.statusCode == tt.args.c.Writer.Status() {
					assert.NotNil(t, tt.args.c)
					assert.Equal(t, tt.statusCode, tt.args.c.Writer.Status())
				} else {
					panic("error")
				}

			} else {
				assert.Equal(t, http.StatusOK, tt.args.c.Writer.Status())
			}
		})
	}
}

func TestBaseService_Delete(t *testing.T) {
	mocks.Setup()

	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		c     *gin.Context
		model any
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		statusCode int
	}{
		// TODO: Add test cases.
		{
			name: "BaseService Delete Success",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectBegin()

					mockRepo.On("FindOne", mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					mockRepo.On("Delete", mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Params = gin.Params{
						gin.Param{Key: "ID", Value: "123"},
					}
					return c
				}(),
			},
			wantErr:    false,
			statusCode: 200,
		},
		{
			name: "BaseService Delete Failed",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectBegin()

					mockRepo.On("FindOne", mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(gorm.ErrRecordNotFound)

					mockRepo.On("Delete", mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)

					sqlMock.ExpectCommit()

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					c.Params = gin.Params{
						gin.Param{Key: "ID", Value: "123"},
					}
					return c
				}(),
			},
			wantErr:    true,
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseService{
				BaseRepository: tt.fields.BaseRepository,
			}

			assert.NotPanics(t, func() {
				b.Delete(tt.args.c, tt.args.model)
			})

			logrus.Infof("Writer Status = %+v", tt.args.c.Writer.Status())

			if tt.wantErr {

				if tt.statusCode == tt.args.c.Writer.Status() {
					assert.NotNil(t, tt.args.c)
					assert.Equal(t, tt.statusCode, tt.args.c.Writer.Status())
				} else {
					panic("error")
				}

			} else {
				assert.Equal(t, http.StatusOK, tt.args.c.Writer.Status())
			}
		})
	}
}

func TestBaseService_IsExit(t *testing.T) {
	type fields struct {
		BaseRepository repository.BaseRepositoryInterface
	}
	type args struct {
		c         *gin.Context
		condition map[string]any
		model     any
		tx        *gorm.DB
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		statusCode int
	}{
		// TODO: Add test cases.
		{
			name: "BaseService Is Exist Panic",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectQuery(`SELECT \* FROM "base_models" WHERE ID = \$1 AND "base_models"\."deleted_at" IS NULL ORDER BY "base_models"\."id" LIMIT \$2`).
						WithArgs(1, 1).
						WillReturnRows(sqlMock.NewRows([]string{"id"}).AddRow(0))

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					return c
				}(),
				condition: map[string]any{
					"ID": 1,
				},
				model: &model.BaseModel{},
			},
			wantErr:    true,
			statusCode: 400,
		},
		{
			name: "BaseService Is Exist Not Panic",
			fields: fields{
				BaseRepository: func() repository.BaseRepositoryInterface {
					mockRepo := new(mocks.BaseRepositoryInterface)

					mockDb, sqlMock := mocks.NewMockDB()
					mockRepo.On("ClientDb").Return(mockDb)

					sqlMock.ExpectQuery(`SELECT \* FROM "base_models" WHERE ID = \$1 AND "base_models"\."deleted_at" IS NULL ORDER BY "base_models"\."id" LIMIT \$2`).
						WithArgs(1, 1).
						WillReturnRows(sqlMock.NewRows([]string{"id"}))

					return mockRepo
				}(),
			},
			args: args{
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					return c
				}(),
				condition: map[string]any{
					"ID": 1,
				},
				model: &model.BaseModel{},
			},
			wantErr:    false,
			statusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseService{
				BaseRepository: tt.fields.BaseRepository,
			}

			assert.NotPanics(t, func() {
				b.IsExist(tt.args.tx, tt.args.c, tt.args.condition, tt.args.model)
			})

			logrus.Infof("Writer Status = %+v", tt.args.c.Writer.Status())

			if tt.wantErr {

				if tt.statusCode == tt.args.c.Writer.Status() {
					assert.NotNil(t, tt.args.c)
					assert.Equal(t, tt.statusCode, tt.args.c.Writer.Status())
				} else {
					panic("error")
				}

			} else {
				assert.Equal(t, http.StatusOK, tt.args.c.Writer.Status())
			}
		})
	}
}
