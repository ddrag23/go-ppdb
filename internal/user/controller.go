package user

import (
	"ddrag23/ppdb/utils"
	"encoding/json"
	"net/http"
)

type IController interface{
	index(w http.ResponseWriter, r *http.Request)
	store(w http.ResponseWriter, r *http.Request)
}
type controller struct{
	service IService
}

func NewController(service IService) IController {
	return &controller{service}
}

func (controller *controller) index(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")
	result, err := controller.service.GetAll()
	utils.PanicIfError(err)
	marshal,err := json.Marshal(result)
	utils.PanicIfError(err)
	w.Write(marshal)
}

func (c *controller) store(w http.ResponseWriter, r *http.Request)  {
	result,err := c.service.Create()
	utils.PanicIfError(err)
	marshal,err := json.Marshal(result)
	utils.PanicIfError(err)
	w.Write(marshal)
}
