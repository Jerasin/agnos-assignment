package service

import (
	"agnos-assignment/app/mocks"
	"agnos-assignment/app/repository"
	"agnos-assignment/app/request"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPatientService_SearchDetail(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type fields struct {
		PatientRepository repository.PatientRepositoryInterface
	}
	type args struct {
		c     *gin.Context
		query *request.PatientRequestModel
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
			name: "Success Patient Search Detail",
			fields: fields{
				PatientRepository: func() repository.PatientRepositoryInterface {
					mockRepo := new(mocks.MockPatientRepository)

					mockBaseRepo := new(mocks.MockBaseRepository)
					mockRepo.On("GetBaseRepo").Return(mockBaseRepo)

					mockDb, sqlMock := mocks.NewMockDB()
					mockBaseRepo.On("ClientDb").Return(mockDb)
					row := sqlMock.NewRows([]string{"FirstNameTh", "LastNameTh"}).AddRow("John", "Doe")

					sqlMock.ExpectQuery(`SELECT`).WillReturnRows(row)
					return mockRepo
				}(),
			},
			args: args{
				query: &request.PatientRequestModel{
					PassportId: "P1234567",
				},
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					mockAuthService := new(mocks.MockJWTService)
					mockAuthService.On("GetPayloadInToken", mock.Anything).Return(jwt.MapClaims{
						"hospital_id": "1",
					}).Once()

					c.Set("JWTService", mockAuthService)
					return c
				}(),
			},
			wantErr:    false,
			statusCode: 200,
		},
		{
			name: "Error Patient Search Detail Data Not Found",
			fields: fields{
				PatientRepository: func() repository.PatientRepositoryInterface {
					mockRepo := new(mocks.MockPatientRepository)

					mockBaseRepo := new(mocks.MockBaseRepository)
					mockRepo.On("GetBaseRepo").Return(mockBaseRepo)

					mockDb, sqlMock := mocks.NewMockDB()
					mockBaseRepo.On("ClientDb").Return(mockDb)

					row := sqlMock.NewRows([]string{"FirstNameTh", "LastNameTh"})
					sqlMock.ExpectQuery(`SELECT`).WillReturnRows(row)

					return mockRepo
				}(),
			},
			args: args{
				query: &request.PatientRequestModel{
					PassportId: "P1234567",
				},
				c: func() *gin.Context {
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)
					mockAuthService := new(mocks.MockJWTService)
					mockAuthService.On("GetPayloadInToken", mock.Anything).Return(jwt.MapClaims{
						"hospital_id": "1",
					}).Once()

					c.Set("JWTService", mockAuthService)
					return c
				}(),
			},
			wantErr:    true,
			statusCode: 404,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			s := PatientService{
				PatientRepository: tt.fields.PatientRepository,
			}

			s.SearchDetail(tt.args.c, tt.args.query)

			if tt.wantErr {
				// fmt.Printf("Writer Status = %+v", c.Writer.Status())
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
