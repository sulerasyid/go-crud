package main

import (
	"github.com/sulerasyid/go-crud/config"
	"github.com/sulerasyid/go-crud/controller"

	"fmt"
	"os"

	"github.com/sulerasyid/go-crud/helper"
	"github.com/sulerasyid/go-crud/infrastructure"
	"github.com/sulerasyid/go-crud/repository"
	"github.com/sulerasyid/go-crud/router"
	"github.com/sulerasyid/go-crud/service"

	"net/http"
	"time"

	"github.com/go-playground/validator"
)

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)
	db, errDB := config.InitDB()
	if errDB != nil {
		logger.LogError("%s", errDB)
	}
	logger.LogAccess("a")
	validate := validator.New()

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagsServiceImpl(tagRepository, validate)

	//Init controller
	tagController := controller.NewTagController(tagService, logger)

	routes := router.NewRouter(tagController)

	addr := os.Getenv("SERVER_PORT")
	fmt.Println("run on port", addr)
	server := &http.Server{
		Addr:           addr,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
