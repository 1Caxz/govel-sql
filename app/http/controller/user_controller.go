package controller

import (
	"fmt"
	"govel/app/exception"
	"govel/app/helper"
	"govel/app/http/middleware"
	"govel/app/model"
	"govel/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service *service.UserService) UserController {
	return UserController{service: *service}
}

func (controller *UserController) Route(route fiber.Router) {
	group := route.Group("/v1/users")
	group.Get("/", controller.Index)
	group.Get("/search/:query", controller.Search)

	group.Post("/login", controller.Login)
	group.Post("/refresh-token", controller.RefreshToken).Name("refresh-token")
	group.Post("/register", controller.Register)
	group.Post("/update/:id", middleware.Authenticate, controller.Update)
	group.Post("/delete/:id", middleware.Authenticate, controller.Delete)

	// Add this endpoint at the bottom to avoid the path conflict
	group.Get("/:id", controller.Show)
}

func (ctx *UserController) Index(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	exception.PanicIfNeeded(err)

	data, isNextPage := ctx.service.List(model.GetUserRequest{
		Page:  page,
		Limit: 10,
	})

	// Return pagination response if next page is exist
	if isNextPage {
		return c.Status(200).JSON(model.PaginateResponse{
			Code:    200,
			Message: "OK",
			Next:    fmt.Sprintf(c.BaseURL()+c.Path()+"?page=%d", page+1),
			Data:    data,
		})
	}

	// Return non pagination response
	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func (ctx *UserController) Search(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	exception.PanicIfNeeded(err)

	data, isNextPage := ctx.service.SearchList(model.GetUserRequest{
		Query: c.Params("query"),
		Page:  page,
		Limit: 10,
	})

	// Return pagination response if next page is exist
	if isNextPage {
		return c.Status(200).JSON(model.PaginateResponse{
			Code:    200,
			Message: "OK",
			Next:    fmt.Sprintf(c.BaseURL()+c.Path()+"?page=%d", page+1),
			Data:    data,
		})
	}

	// Return non pagination response
	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func (ctx *UserController) Login(c *fiber.Ctx) error {
	data := ctx.service.Login(model.LoginUserRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	})

	token := helper.MakeECDSAToken(&data, jwt.SigningMethodES256)

	refreshTokenURL, err := c.GetRouteURL("refresh-token", nil)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data: model.TokenResponse{
			Type:            "bearer",
			Alg:             "es256",
			RefreshTokenURL: c.BaseURL() + refreshTokenURL,
			Token:           token,
			Claims:          data,
		},
	})
}

func (ctx *UserController) RefreshToken(c *fiber.Ctx) error {
	data := ctx.service.RefreshToken(model.RefreshTokenUserRequest{
		Token: c.FormValue("token"),
	})

	token := helper.MakeECDSAToken(&data, jwt.SigningMethodES256)

	refreshTokenURL, err := c.GetRouteURL("refresh-token", nil)
	exception.PanicIfNeeded(err)

	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data: model.TokenResponse{
			Type:            "bearer",
			Alg:             "es256",
			RefreshTokenURL: c.BaseURL() + refreshTokenURL,
			Token:           token,
			Claims:          data,
		},
	})
}

func (ctx *UserController) Register(c *fiber.Ctx) error {
	data := ctx.service.Register(model.RegisterUserRequest{
		SocialId:   c.FormValue("social_id"),
		Email:      c.FormValue("email"),
		Name:       c.FormValue("name"),
		Password:   c.FormValue("password"),
		Repassword: c.FormValue("repassword"),
	})
	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func (ctx *UserController) Show(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	exception.PanicIfNeeded(err)

	data := ctx.service.Single(model.GetUserRequest{
		Id: id,
	})

	// Return non pagination response
	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func (ctx *UserController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	exception.PanicIfNeeded(err)

	data := ctx.service.Update(model.UpdateUserRequest{
		Token:    c.FormValue("token"),
		Id:       id,
		Name:     c.FormValue("name"),
		Location: c.FormValue("location"),
		Desc:     c.FormValue("desc"),
	})

	// Return non pagination response
	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func (ctx *UserController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	exception.PanicIfNeeded(err)

	data := ctx.service.Delete(model.DeleteUserRequest{
		Token: c.FormValue("token"),
		Id:    id,
	})

	// Return non pagination response
	return c.Status(200).JSON(model.WebResponse{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}
