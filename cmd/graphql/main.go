package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/dao"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/db"
	"github.com/yyh-gl/hobigon-golang-api-server/app/presentation/graphql"
	"github.com/yyh-gl/hobigon-golang-api-server/app/presentation/graphql/generated"
)

func main() {
	// TODO: wire導入
	br := dao.NewBlog(db.NewDB())
	r := graphql.Resolver{BlogRepository: br}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:8081/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
