package user

import "apipost/model/global"

// User - пользователь сервиса с его параметрами
type User struct {
	Id      global.UserId //
	Name    string        //
	Surname string        //
}

// NewUser -
func NewUser(id global.UserId, name string, surName string) User {
	return User{
		Id:      id,
		Name:    name,
		Surname: surName,
	}
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

func GetAllUsers() (users []User, err error) {
	users = []User{
		{1, "Джон", "До"},
		{2, "Говард", "Рорк"},
		{3, "Джек", "Доусон"},
		{4, "Лизель", "Мемингер"},
		{5, "Джейн", "Эйр"},
		{6, "Мартин", "Иден"},
		{7, "Джон", "Голт"},
		{8, "Сэмвелл", "Тарли"},
		{9, "Гермиона", "Грейнджер"},
	}
	return
}
