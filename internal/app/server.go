package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"repository-hotel-booking/internal/app/model"
	"repository-hotel-booking/internal/app/repository"
	"repository-hotel-booking/internal/app/router/routes"
	"repository-hotel-booking/internal/app/service"
	"repository-hotel-booking/internal/app/util"
)

func Init() {

	//auto.Load()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	//engine.Use(gin.Recovery(), cors.New(cfg.CORS))
	//engine.MaxMultipartMemory = app.CFG.Service.RequestSize
	db := util.InitConnection(CFG.DB)
	repo := repository.New(db)
	s := service.NewService(repo)
	routes.SetupRoutes(engine, s)

	fmt.Printf("Running in port %v...\n", CFG.Server.Port)
	listen(CFG.Server.Port, engine)
}
func listen(port string, e *gin.Engine) {

	log.Fatal(http.ListenAndServe(port, e))
}

var CFG = &model.Config{}