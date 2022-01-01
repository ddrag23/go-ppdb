package user

import "github.com/go-chi/chi/v5"

func RunRouter(r chi.Router)  {
	model := NewModel()
	service := NewService(model)
	controller := NewController(service)
	r.Get("/user",controller.index)
	r.Get("/user/store",controller.store)
}