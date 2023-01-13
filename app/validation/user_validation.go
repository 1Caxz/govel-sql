package validation

import (
	"govel/app/exception"
	"govel/app/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func UserRefreshTokenValidate(request model.RefreshTokenUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Token, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UserLoginValidate(request model.LoginUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UserRegisterValidate(request model.RegisterUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Repassword, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}

	if request.Password != request.Repassword {
		panic(exception.ValidationError{
			Message: "Password doesn't match.",
		})
	}
}

func UserUpdateValidate(request model.UpdateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Token, validation.Required),
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Location, validation.Required),
		validation.Field(&request.Desc, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UserDeleteValidate(request model.DeleteUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Token, validation.Required),
		validation.Field(&request.Id, validation.Required, validation.Min(1)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UserShowValidate(request model.GetUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required, validation.Min(1)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UserListhValidate(request model.GetUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Page, validation.Required, validation.Min(1)),
		validation.Field(&request.Limit, validation.Required, validation.Min(1)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UserSearchValidate(request model.GetUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Query, validation.Required),
		validation.Field(&request.Page, validation.Required, validation.Min(1)),
		validation.Field(&request.Limit, validation.Required, validation.Min(1)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
