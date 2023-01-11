package service

import (
	"govel/app/entity"
	"govel/app/exception"
	"govel/app/model"
	"govel/app/repository"
	"govel/app/validation"

	"github.com/mintance/go-uniqid"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: *userRepository,
	}
}

func (service *userServiceImpl) Login(request model.LoginUserRequest) (response model.LoginUserResponse) {
	// Validate the user request data
	validation.UserLoginValidate(request)

	// Check user is exist
	user := service.UserRepository.FetchByEmail(request.Email)
	if user == nil {
		exception.PanicResponse("Email not registered.")
	}

	// Check the hash password is correct
	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if errPassword != nil {
		exception.PanicResponse("Email or password doesn't match.")
	}

	// Response the data
	response = model.LoginUserResponse{
		Id:       user.ID,
		SocialId: user.SocialId,
		Email:    user.Email,
		Nick:     user.Nick,
		Name:     user.Name,
		Pic:      user.Pic,
		Location: user.Location,
		Desc:     user.Desc,
	}
	return response
}

func (service *userServiceImpl) Register(request model.RegisterUserRequest) (response model.RegisterUserResponse) {
	// Validate the user request data
	validation.UserRegisterValidate(request)

	// Check the user email is exist
	result := service.UserRepository.FetchByEmail(request.Email)
	if result != nil {
		exception.PanicResponse("Email already register, please login.")
	}

	// Hasing the password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfNeeded(err)

	// Insert the data
	data := entity.User{
		SocialId: request.SocialId,
		Email:    request.Email,
		Nick:     uniqid.New(uniqid.Params{Prefix: "govel", MoreEntropy: false}),
		Name:     request.Name,
		Password: string(password),
	}
	user := service.UserRepository.Insert(data)

	// Response the data
	response = model.RegisterUserResponse{
		Id:       user.ID,
		SocialId: user.SocialId,
		Email:    user.Email,
		Nick:     user.Nick,
		Name:     user.Name,
		Pic:      user.Pic,
		Location: user.Location,
		Desc:     user.Desc,
	}
	return response
}

func (service *userServiceImpl) Single(request model.GetUserRequest) (response model.GetUserResponse) {
	// Validate the user request data
	validation.UserShowValidate(request)

	// Get the data
	user := service.UserRepository.Fetch(uint(request.Id))

	// Response the data
	response = model.GetUserResponse{
		Id:       user.ID,
		SocialId: user.SocialId,
		Email:    user.Email,
		Nick:     user.Nick,
		Name:     user.Name,
		Pic:      user.Pic,
		Location: user.Location,
		Desc:     user.Desc,
	}
	return response
}

func (service *userServiceImpl) List(request model.GetUserRequest) (responses []model.GetUserResponse, isNextPage bool) {
	// Validate the user request data
	validation.UserListhValidate(request)

	// Get the pagination data, +1 limit to check the next page is exist
	offset := 0
	limit := request.Limit + 1
	if request.Page > 1 {
		offset = (request.Page * limit) - limit
	}
	users := service.UserRepository.FetchAll(limit, offset)

	// Response the data
	isNextPage = false
	for i := 0; i < len(users); i++ {
		if i == request.Limit {
			isNextPage = true
			break
		}
		user := users[i]
		responses = append(responses, model.GetUserResponse{
			Id:       user.ID,
			SocialId: user.SocialId,
			Email:    user.Email,
			Nick:     user.Nick,
			Name:     user.Name,
			Pic:      user.Pic,
			Location: user.Location,
			Desc:     user.Desc,
		})
	}

	return responses, isNextPage
}

func (service *userServiceImpl) SearchList(request model.GetUserRequest) (responses []model.GetUserResponse, isNextPage bool) {
	// Validate the user request data
	validation.UserSearchValidate(request)

	// Get the data
	offset := 0
	limit := request.Limit + 1
	if request.Page > 1 {
		offset = (request.Page * limit) - limit
	}
	users := service.UserRepository.FindAll(request.Query, limit, offset)

	// Response the data
	isNextPage = false
	for i := 0; i < len(users); i++ {
		if i == request.Limit {
			isNextPage = true
			break
		}
		user := users[i]
		responses = append(responses, model.GetUserResponse{
			Id:       user.ID,
			SocialId: user.SocialId,
			Email:    user.Email,
			Nick:     user.Nick,
			Name:     user.Name,
			Pic:      user.Pic,
			Location: user.Location,
			Desc:     user.Desc,
		})
	}
	return responses, isNextPage
}

func (service *userServiceImpl) Update(request model.UpdateUserRequest) (response model.UpdateUserResponse) {
	// Validate the user request data
	validation.UserUpdateValidate(request)

	// Update the data
	data := entity.User{
		ID:       uint(request.Id),
		Name:     request.Name,
		Location: request.Location,
		Desc:     request.Desc,
	}
	user := service.UserRepository.Update(data)

	// Response the new data
	response = model.UpdateUserResponse{
		Id:       user.ID,
		SocialId: user.SocialId,
		Email:    user.Email,
		Nick:     user.Nick,
		Name:     user.Name,
		Pic:      user.Pic,
		Location: user.Location,
		Desc:     user.Desc,
	}
	return response
}

func (service *userServiceImpl) Delete(request model.DeleteUserRequest) (response model.DeleteUserResponse) {
	// Validate the user request data
	validation.UserDeleteValidate(request)

	// Delete the data
	service.UserRepository.Delete(uint(request.Id))

	// Response
	response = model.DeleteUserResponse{
		Id:      uint(request.Id),
		Message: "Data deleted.",
	}
	return response
}
