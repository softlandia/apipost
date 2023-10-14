package user

// Data - пользователь сервиса с его параметрами
type Data struct {
	Id      int    `json:"id"`       //
	Name    string `json:"name"`     //
	Surname string `json:"sur_name"` //
}

type List []Data

// NewUser -
func NewUser(id int, name string, surName string) Data {
	return Data{
		Id:      id,
		Name:    name,
		Surname: surName,
	}
}

// IsEmpty - пустой ли объект
func (u Data) IsEmpty() bool {
	return u.Id == 0 //&& u.Name == "" && len(u.Surname) == 0
}
