package main

import (
	"log"
	"os"

	"github.com/MinorvaFalk/go-graphql-practice/core/utils"
	handler "github.com/MinorvaFalk/go-graphql-practice/handlers"
	"github.com/MinorvaFalk/go-graphql-practice/repository"
	"github.com/MinorvaFalk/go-graphql-practice/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	l := log.New(os.Stdout, "[Book API] ", log.LstdFlags)
	env := utils.NewEnv(l)
	db := utils.GetDatabaseGORM(l, env)

	br := repository.NewBookPostgresRepository(l, db.DB)
	bs := services.NewBookService(l, br)

	ar := repository.NewAuthorPostgresRepository(l, db.DB)
	as := services.NewAuthorService(l, ar)

	r := gin.Default()
	r.Use(cors.Default())

	handler.NewGraphqlHandler(r, bs, as)

	// Handle REST Request
	// handler.NewRestHandler(r)

	r.Run(env.Port)
}
