package module

import (
	handler "github.com/MinorvaFalk/go-graphql-practice/handlers"
	"go.uber.org/fx"
)

var HandlerModule = fx.Provide(
	handler.NewGraphqlHandler,
	handler.NewRestHandler,
	handler.NewHandler,
)
