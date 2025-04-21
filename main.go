package main

import (
	"fmt"

	"agnos-assignment/app/config"
	"agnos-assignment/app/router"
	"agnos-assignment/app/utils"
)

func main() {
	fmt.Println("Main Start")

	config.EnvConfig()
	port := config.GetEnv("PORT", "3000")

	baseModule := router.NewBaseModule()
	app := router.RouterInit(baseModule)

	appInfo := fmt.Sprintf("0.0.0.0:%s", port)

	db := utils.InitDbClient()
	initDataClient := utils.InitDataClientInit(db)
	initDataClient.InitHospital()
	initDataClient.InitPatient()

	for _, item := range app.Routes() {
		println("method:", item.Method, "path:", item.Path)
	}

	app.Run(appInfo) // listen and serve on 0.0.0.0:8080
}
