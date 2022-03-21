package handler

type Handler interface {
	SetupHandler()
}

type Handlers []Handler

func NewHandler(gh GraphqlHandler, rh RestHandler) Handlers {
	return Handlers{
		gh,
		rh,
	}
}

func (h Handlers) SetupHandler() {
	for _, handler := range h {
		handler.SetupHandler()
	}
}
