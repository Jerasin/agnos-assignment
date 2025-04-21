package router

import (
	"agnos-assignment/app/dto"
	"agnos-assignment/app/module"
	"agnos-assignment/app/request"
	"agnos-assignment/app/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/most-develop/tonic"
)

type BaseModuleInit struct {
	PatientModule  *module.PatientModule
	HospitalModule *module.HospitalModule
	StaffModule    *module.StaffModule
}

func NewBaseModule() BaseModuleInit {
	patientInit := module.PatientModuleInit()
	hospitalInit := module.HospitalModuleInit()
	staffInit := module.StaffModuleInit()

	return BaseModuleInit{
		PatientModule:  patientInit,
		HospitalModule: hospitalInit,
		StaffModule:    staffInit,
	}
}

func RouterInit(init BaseModuleInit) *gin.Engine {
	tonic.Init(&tonic.Config{
		OpenAPIVersion: "3.0.0",
		Info: map[string]interface{}{
			"title":       "Hello World API",
			"description": "A simple CRUD example using Go and PostgreSQL",
			"version":     "1.0.0",
		},
	})

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")

	patientRouter := api.Group("patient")
	{

		tonic.CreateRoutes(patientRouter.BasePath(), []tonic.Route{
			{

				Tags:   []string{"Patient"},
				Method: tonic.Get,
				Url:    "",
				HandlerRegister: func(path string) {
					patientRouter.GET(path, init.PatientModule.PatientCtrl.GetList)

				},
				Schema: &tonic.RouteSchema{
					// Params: request.ParamID{},
					// Body:   request.PatientRequest{},
					Response: map[int]any{
						200: dto.ApiPaginationResponse[[]response.PatientSearchModel]{},
					},
				},
			},
			{

				Tags:   []string{"Patient"},
				Method: tonic.Get,
				Url:    "/search/:ID",
				HandlerRegister: func(path string) {
					patientRouter.GET(path, init.PatientModule.PatientCtrl.Search)

				},
				Schema: &tonic.RouteSchema{
					Params: request.ParamID{},
					// Body:   request.PatientRequest{},
					Response: map[int]any{
						200: response.PatientSearchModel{},
					},
				},
			},
		})
	}

	// patientRouter.Use(middleware.AuthorizeJwt())
	// patientRouter.GET("/search", init.PatientModule.PatientCtrl.SearchDetail)

	hospitalRouter := api.Group("hospital")
	{
		tonic.CreateRoutes(hospitalRouter.BasePath(), []tonic.Route{
			{
				Tags:   []string{"Hospital"},
				Method: tonic.Post,
				Url:    "",
				HandlerRegister: func(path string) {
					hospitalRouter.POST(path, init.HospitalModule.HospitalCtrl.Create)
				},
				Schema: &tonic.RouteSchema{
					Body: request.Hospital{},
					Response: map[int]any{
						200: dto.ApiCreateResponse{},
					},
				},
			},
			{
				Tags:   []string{"Hospital"},
				Method: tonic.Get,
				Url:    "",
				HandlerRegister: func(path string) {
					hospitalRouter.GET(path, init.HospitalModule.HospitalCtrl.GetList)
				},
				Schema: &tonic.RouteSchema{
					Querystring: request.BasePaginationModel{},
					Response: map[int]any{
						200: dto.ApiPaginationResponse[[]response.PatientSearchModel]{},
					},
				},
			},
			{
				Tags:   []string{"Hospital"},
				Method: tonic.Get,
				Url:    "/:ID",
				HandlerRegister: func(path string) {
					hospitalRouter.GET(path, init.HospitalModule.HospitalCtrl.GetDetail)
				},
				Schema: &tonic.RouteSchema{
					Params: request.ParamID{},
					Response: map[int]any{
						200: dto.ApiResponse[response.HospitalModel]{},
					},
				},
			},
			{
				Tags:   []string{"Hospital"},
				Method: tonic.Put,
				Url:    "/:ID",
				HandlerRegister: func(path string) {
					hospitalRouter.PUT(path, init.HospitalModule.HospitalCtrl.Update)
				},
				Schema: &tonic.RouteSchema{
					Params: request.ParamID{},
					Body:   request.Hospital{},
					Response: map[int]any{
						200: dto.ApiCreateResponse{},
					},
				},
			},
			{
				Tags:   []string{"Hospital"},
				Method: tonic.Delete,
				Url:    "/:ID",
				HandlerRegister: func(path string) {
					hospitalRouter.DELETE(path, init.HospitalModule.HospitalCtrl.Delete)
				},
				Schema: &tonic.RouteSchema{
					Params: request.ParamID{},
					// Body:   request.Hospital{},
					Response: map[int]any{
						200: dto.ApiCreateResponse{},
					},
				},
			},
		})
	}

	// hospitalRouter.GET("", init.HospitalModule.HospitalCtrl.GetList)

	// staffRouter := api.Group("staff")
	// staffRouter.POST("", init.StaffModule.StaffCtrl.CreateStaff)
	// staffRouter.POST("/login", init.StaffModule.StaffCtrl.LoginStaff)

	api.GET("/docs/*w", gin.WrapH(http.StripPrefix("/api/docs", tonic.GetHandler())))
	return router
}
