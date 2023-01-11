package repository

import (
	"govel/app/entity"
)

type UserRepository interface {
	Fetch(id uint) (user *entity.User)

	FetchByEmail(email string) (user *entity.User)

	FetchAll(limit int, offset int) (users []entity.User)

	FindAll(query string, limit int, offset int) (users []entity.User)

	Insert(data entity.User) (user entity.User)

	Update(data entity.User) (user entity.User)

	Delete(id uint)
}
