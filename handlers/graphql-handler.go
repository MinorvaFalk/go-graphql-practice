package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MinorvaFalk/go-graphql-practice/core/middleware"
	"github.com/MinorvaFalk/go-graphql-practice/graph"
	"github.com/MinorvaFalk/go-graphql-practice/services"
	"github.com/gin-gonic/gin"
)

type GraphqlHandler struct {
	r  *gin.Engine
	bs services.BookService
	as services.AuthorService
}

func NewGraphqlHandler(r *gin.Engine, bs services.BookService, as services.AuthorService) GraphqlHandler {
	h := GraphqlHandler{r, bs, as}
	return h
}

func (gh GraphqlHandler) SetupHandler() {
	gh.r.POST("/query", middleware.GinContextToContextMiddleware(), gh.graphqlHandler())
	gh.r.GET("/", gh.playgroundHandler())
}

// Defining the Graphql handler
func (gh *GraphqlHandler) graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Bs: gh.bs,
		As: gh.as,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func (gh *GraphqlHandler) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
