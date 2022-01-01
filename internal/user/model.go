package user

import (
	"database/sql"
	"ddrag23/ppdb/db"
	"ddrag23/ppdb/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)


type User struct{
	ID string `json:"id"`
	Photo sql.NullString `json:"photo"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email string `json:"email"`
	RoleId string `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Model interface{
	FindAll() ([]User, error)
	Create() (string,error)
}

type model struct{

}

func NewModel() Model {
	return &model{}
}

func (model *model) FindAll() ([]User,error) {
	var result []User
	var user User
	conn := db.GetConnection()
	rows,err := conn.Query("select id, name, username from users")
	utils.PanicIfError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID,&user.Name,&user.Username)
		utils.PanicIfError(err)
		result = append(result, user)
	}
	return result,nil
}

func (model *model) Create() (string,error)  {
	var id string
	conn := db.GetConnection()
	err := conn.QueryRow("insert into users (id,name,username,email) values ($1,$2,$3,$4) returning id",uuid.New().String(),"coba","coba","coba@mail.co").Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "",err
	}
	return "success create",nil
	
}
