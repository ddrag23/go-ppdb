package user

type IService interface{
	GetAll() ([]User,error)
	Create() (string,error)
}

type service struct{
	model Model
}

func NewService(model Model) IService {
	return &service{model}
}

func (service *service) GetAll() ([]User,error) {
	user,err := service.model.FindAll()
	if err !=nil {
		return nil,err
	}
	return user,nil
}
func (service *service)Create() (string,error)  {
	msg,err := service.model.Create()
	if err !=nil {
		return "",err
	}
	return msg,nil
}