// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yyh-gl/hobigon-golang-api-server/app"
	"github.com/yyh-gl/hobigon-golang-api-server/cmd/api/di"
	"github.com/yyh-gl/hobigon-golang-api-server/handler/rest"
	"github.com/yyh-gl/hobigon-golang-api-server/infra"
	"github.com/yyh-gl/hobigon-golang-api-server/infra/db"
	"github.com/yyh-gl/hobigon-golang-api-server/infra/igateway"
	"github.com/yyh-gl/hobigon-golang-api-server/infra/irepository"
	"github.com/yyh-gl/hobigon-golang-api-server/infra/iservice"
	"github.com/yyh-gl/hobigon-golang-api-server/usecase"
)

// Injectors from wire.go:

func initApp() *di.Container {
	gormDB := db.NewDB()
	blogRepository := irepository.NewBlogRepository(gormDB)
	slackGateway := igateway.NewSlackGateway()
	blogUseCase := usecase.NewBlogUseCase(blogRepository, slackGateway)
	blogHandler := rest.NewBlogHandler(blogUseCase)
	birthdayRepository := irepository.NewBirthdayRepository(gormDB)
	birthdayUseCase := usecase.NewBirthdayUseCase(birthdayRepository)
	birthdayHandler := rest.NewBirthdayHandler(birthdayUseCase)
	taskGateway := igateway.NewTaskGateway()
	notificationService := iservice.NewNotificationService(slackGateway)
	rankingService := iservice.NewRankingService()
	notificationUseCase := usecase.NewNotificationUseCase(taskGateway, slackGateway, birthdayRepository, notificationService, rankingService)
	notificationHandler := rest.NewNotificationHandler(notificationUseCase)
	logger := app.NewCLILogger()
	container := &di.Container{
		HandlerBlog:         blogHandler,
		HandlerBirthday:     birthdayHandler,
		HandlerNotification: notificationHandler,
		DB:                  gormDB,
		Logger:              logger,
	}
	return container
}

// wire.go:

var appSet = wire.NewSet(app.CLISet, infra.WireSet, usecase.WireSet, rest.WireSet)
