package module

import "go.uber.org/fx"

var Module = fx.Options(
	UtilsModule,
	RepositoryModule,
	ServiceModule,
	HandlerModule,
)
