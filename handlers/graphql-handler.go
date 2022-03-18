package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MinorvaFalk/go-graphql-practice/core/middleware"
	"github.com/MinorvaFalk/go-graphql-practice/graph"
	"github.com/MinorvaFalk/go-graphql-practice/services"
	"github.com/gin-gonic/gin"
)

type graphqlHandler struct {
	bs services.BookService
	as services.AuthorService
}

func NewGraphqlHandler(r *gin.Engine, bs services.BookService, as services.AuthorService) {
	h := &graphqlHandler{bs, as}

	r.POST("/query", middleware.GinContextToContextMiddleware(), h.graphqlHandler())
	r.GET("/", h.playgroundHandler())
}

// Defining the Graphql handler
func (gh *graphqlHandler) graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		gh.bs,
		gh.as,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func (gh *graphqlHandler) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
