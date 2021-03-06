// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yyh-gl/hobigon-golang-api-server/app"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/dao"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/db"
	"github.com/yyh-gl/hobigon-golang-api-server/app/presentation/cli"
	"github.com/yyh-gl/hobigon-golang-api-server/app/usecase"
	"github.com/yyh-gl/hobigon-golang-api-server/cmd/rest/di"
)

// Injectors from wire.go:

func initApp() *di.ContainerCLI {
	task := dao.NewTask()
	slack := dao.NewSlack()
	gormDB := db.NewDB()
	birthday := dao.NewBirthday(gormDB)
	notification := usecase.NewNotification(task, slack, birthday)
	cliNotification := cli.NewNotification(notification)
	logger := app.NewCLILogger()
	containerCLI := &di.ContainerCLI{
		HandlerNotification: cliNotification,
		DB:                  gormDB,
		Logger:              logger,
	}
	return containerCLI
}

// wire.go:

var appSet = wire.NewSet(app.CLISet, infra.CLISet, usecase.CLISet, cli.WireSet)
