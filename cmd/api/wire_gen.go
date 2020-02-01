// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yyh-gl/hobigon-golang-api-server/app"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/dao"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/db"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/iservice"
	"github.com/yyh-gl/hobigon-golang-api-server/app/interface/rest"
	"github.com/yyh-gl/hobigon-golang-api-server/app/usecase"
	"github.com/yyh-gl/hobigon-golang-api-server/cmd/api/di"
)

// Injectors from wire.go:

func initApp() *di.ContainerAPI {
	gormDB := db.NewDB()
	blog := dao.NewBlog(gormDB)
	slack := dao.NewSlack()
	usecaseBlog := usecase.NewBlog(blog, slack)
	restBlog := rest.NewBlog(usecaseBlog)
	birthday := dao.NewBirthday(gormDB)
	usecaseBirthday := usecase.NewBirthday(birthday)
	restBirthday := rest.NewBirthday(usecaseBirthday)
	task := dao.NewTask()
	ranking := iservice.NewRanking()
	notification := usecase.NewNotification(task, slack, birthday, ranking)
	restNotification := rest.NewNotification(notification)
	logger := app.NewAPILogger()
	containerAPI := &di.ContainerAPI{
		HandlerBlog:         restBlog,
		HandlerBirthday:     restBirthday,
		HandlerNotification: restNotification,
		DB:                  gormDB,
		Logger:              logger,
	}
	return containerAPI
}

// wire.go:

var appSet = wire.NewSet(app.APISet, infra.APISet, usecase.APISet, rest.WireSet)
