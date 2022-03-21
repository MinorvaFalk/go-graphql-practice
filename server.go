package main

import (
	"context"

	"github.com/MinorvaFalk/go-graphql-practice/core/utils"
	handler "github.com/MinorvaFalk/go-graphql-practice/handlers"
	"github.com/MinorvaFalk/go-graphql-practice/module"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		module.Module,
		fx.Invoke(startup),
	)

	app.Run()

	// l := log.New(os.Stdout, "[Book API] ", log.LstdFlags)
	// env := utils.NewEnv(l)
	// db := utils.GetDatabaseGORM(l, env)

	// br := repository.NewBookPostgresRepository(l, db.DB)
	// bs := services.NewBookService(l, br)

	// ar := repository.NewAuthorPostgresRepository(l, db.DB)
	// as := services.NewAuthorService(l, ar)

	// r := gin.Default()
	// r.Use(cors.Default())

	// handler.NewGraphqlHandler(r, bs, as)

	// // Handle REST Request
	// // handler.NewRestHandler(r)

	// go func() {

	// }()
}

func startup(lifecycle fx.Lifecycle, db *utils.Database, env *utils.Env, handler handler.Handlers, r *gin.Engine) {
	conn, _ := db.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				handler.SetupHandler()
				r.Run(env.Port)
			}()
			return nil
		},

		OnStop: func(ctx context.Context) error {

			conn.Close()
			return nil
		},
	})

}
