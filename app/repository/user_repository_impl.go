package repository

import (
	"errors"
	"govel/app/entity"
	"govel/app/exception"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: database,
	}
}

func (repository *userRepositoryImpl) Fetch(id uint) (user *entity.User) {
	var data entity.User
	result := repository.DB.First(&data, id)
	exception.PanicIfNeeded(result.Error)
	return &data
}

func (repository *userRepositoryImpl) FetchByEmail(email string) (user *entity.User) {
	var data entity.User
	result := repository.DB.Where("email = ?", email).First(&data)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	exception.PanicIfNeeded(result.Error)
	return &data
}

func (repository *userRepositoryImpl) FetchAll(limit int, offset int) (users []entity.User) {
	var data []entity.User
	result := repository.DB.Limit(limit).Offset(offset).Find(&data)
	exception.PanicIfNeeded(result.Error)
	return data
}

func (repository *userRepositoryImpl) FindAll(query string, limit int, offset int) (users []entity.User) {
	var data []entity.User
	result := repository.DB.Where("MATCH (name) AGAINST ('*" + query + "*' IN BOOLEAN MODE)").Limit(limit).Offset(offset).Find(&data)
	exception.PanicIfNeeded(result.Error)
	return data
}

func (repository *userRepositoryImpl) Insert(data entity.User) (user entity.User) {
	result := repository.DB.Create(&data)
	exception.PanicIfNeeded(result.Error)
	return data
}

func (repository *userRepositoryImpl) Update(data entity.User) (user entity.User) {
	mData := entity.User{}
	result := repository.DB.Model(&mData).Where("id = ?", data.ID).Updates(data)
	exception.PanicIfNeeded(result.Error)
	return mData
}

func (repository *userRepositoryImpl) Delete(id uint) {
	result := repository.DB.Delete(&entity.User{}, id)
	exception.PanicIfNeeded(result.Error)
}
