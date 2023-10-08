package global

import "fmt"

type UserId int

func (id UserId) String() string {
	return fmt.Sprintf("идентификатор пользователя: %d", id)
}
