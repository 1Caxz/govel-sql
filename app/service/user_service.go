package service

import "govel/app/model"

type UserService interface {
	RefreshToken(request model.RefreshTokenUserRequest) (response model.RefreshTokenUserResponse)

	Login(request model.LoginUserRequest) (response model.LoginUserResponse)

	Register(request model.RegisterUserRequest) (response model.RegisterUserResponse)

	Single(request model.GetUserRequest) (response model.GetUserResponse)

	List(request model.GetUserRequest) (responses []model.GetUserResponse, isNextPage bool)

	SearchList(request model.GetUserRequest) (responses []model.GetUserResponse, isNextPage bool)

	Update(request model.UpdateUserRequest) (response model.UpdateUserResponse)

	Delete(request model.DeleteUserRequest) (response model.DeleteUserResponse)
}
