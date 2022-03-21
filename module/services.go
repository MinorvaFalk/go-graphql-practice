package module

import (
	"github.com/MinorvaFalk/go-graphql-practice/services"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	fx.Provide(services.NewAuthorService),
	fx.Provide(services.NewBookService),
)
