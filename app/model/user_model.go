package model

import "github.com/golang-jwt/jwt/v4"

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Id       uint   `json:"id"`
	SocialId string `json:"social_id"`
	Email    string `json:"email"`
	Nick     string `json:"nick"`
	Name     string `json:"name"`
	Pic      string `json:"pic"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
	jwt.StandardClaims
}

type RegisterUserRequest struct {
	SocialId   string `json:"social_id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

type RegisterUserResponse struct {
	Id       uint   `json:"id"`
	SocialId string `json:"social_id"`
	Email    string `json:"email"`
	Nick     string `json:"nick"`
	Name     string `json:"name"`
	Pic      string `json:"pic"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

type UpdateUserRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

type UpdateUserResponse struct {
	Id       uint   `json:"id"`
	SocialId string `json:"social_id"`
	Email    string `json:"email"`
	Nick     string `json:"nick"`
	Name     string `json:"name"`
	Pic      string `json:"pic"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type DeleteUserResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
}

type GetUserRequest struct {
	Id    int    `json:"id"`
	Query string `json:"q"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

type GetUserResponse struct {
	Id       uint   `json:"id"`
	SocialId string `json:"social_id"`
	Email    string `json:"email"`
	Nick     string `json:"nick"`
	Name     string `json:"name"`
	Pic      string `json:"pic"`
	Location string `json:"location"`
	Desc     string `json:"desc"`
}
