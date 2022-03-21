package module

import (
	"github.com/MinorvaFalk/go-graphql-practice/repository"
	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	fx.Provide(repository.NewAuthorPostgresRepository),
	fx.Provide(repository.NewBookPostgresRepository),
)
