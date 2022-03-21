package module

import (
	"github.com/MinorvaFalk/go-graphql-practice/core/utils"
	"go.uber.org/fx"
)

var UtilsModule = fx.Options(
	fx.Provide(utils.NewLogger),
	fx.Provide(utils.NewEnv),
	fx.Provide(utils.GetDatabaseGORM),
	fx.Provide(utils.NewGinEngine),
)
