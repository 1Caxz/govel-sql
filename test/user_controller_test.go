package test

import (
	"encoding/json"
	"govel/app/helper"
	"govel/app/model"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mintance/go-uniqid"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Users(t *testing.T) {
	// Setup request
	request := httptest.NewRequest("GET", "/api/v1/users", nil)

	// Test the request
	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	list := webResponse.Data.([]interface{})
	assert.Equal(t, 10, len(list))
}
func TestUserController_Search(t *testing.T) {
	// Setup request
	request := httptest.NewRequest("GET", "/api/v1/users/search/a", nil)

	// Test the request
	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	list := webResponse.Data.([]interface{})
	assert.LessOrEqual(t, len(list), 10)
}

func TestUserController_Register(t *testing.T) {
	// Setup dynamic email
	email := uniqid.New(uniqid.Params{Prefix: "govel", MoreEntropy: false})
	// Setup form data
	data := strings.NewReader("email=" + email + "@gmail.com&name=Saiful Wicaksana&password=rahasia&repassword=rahasia")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/register", data)

	// Setup header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/x-www-form-urlencoded")

	// Test the request
	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	jsonData, _ := json.Marshal(webResponse.Data)
	registerUserResponse := model.RegisterUserResponse{}
	json.Unmarshal(jsonData, &registerUserResponse)
	assert.Empty(t, registerUserResponse.SocialId)
	assert.Equal(t, request.FormValue("name"), registerUserResponse.Name)
	assert.Equal(t, request.FormValue("email"), registerUserResponse.Email)
}

func TestUserController_Login(t *testing.T) {
	// Setup form data
	data := strings.NewReader("email=consequatur@gmail.com&password=rahasia")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/login", data)

	// Setup header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/x-www-form-urlencoded")

	// Test the request
	response, _ := app.Test(request)
	// assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	jsonData, _ := json.Marshal(webResponse.Data)
	tokenResponse := model.TokenResponse{}
	json.Unmarshal(jsonData, &tokenResponse)
	assert.Equal(t, "bearer", tokenResponse.Type)
	assert.Equal(t, "es256", tokenResponse.Alg)

	// Check token is valid
	token := helper.ParseECDSAToken(tokenResponse.Token, jwt.SigningMethodES256)
	assert.True(t, token.Valid)

	jsonClaims, _ := json.Marshal(tokenResponse.Claims)
	loginUserResponse := model.LoginUserResponse{}
	json.Unmarshal(jsonClaims, &loginUserResponse)
	assert.NotEmpty(t, loginUserResponse.Name)
	assert.Equal(t, request.FormValue("email"), loginUserResponse.Email)
	assert.Equal(t, 1, loginUserResponse.Role)
}

func TestUserController_RefreshToken(t *testing.T) {
	// Setup form data
	data := strings.NewReader("token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInNvY2lhbF9pZCI6IiIsImVtYWlsIjoiY29uc2VxdWF0dXJAZ21haWwuY29tIiwibmljayI6Im9jY2FlY2F0aSIsIm5hbWUiOiJlaXVzIiwicGljIjoiL2Fzc2V0cy9zdGF0aWMvdXNlci5wbmciLCJsb2NhdGlvbiI6IkluZG9uZXNpYSIsImRlc2MiOiIiLCJyb2xlIjoxfQ.8SoZMH7f_iReJ-YQmks92QAXo_TjRzrO3xtsdnD_6bvFcJiwXiBSmr1D_7oHEtjdkHNJpeWGCy_yhDiTlzZSXQ")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/refresh-token", data)

	// Setup header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/x-www-form-urlencoded")

	// Test the request
	response, _ := app.Test(request)
	// assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	jsonData, _ := json.Marshal(webResponse.Data)
	tokenResponse := model.TokenResponse{}
	json.Unmarshal(jsonData, &tokenResponse)
	assert.Equal(t, "bearer", tokenResponse.Type)
	assert.Equal(t, "es256", tokenResponse.Alg)

	// Check token is valid
	token := helper.ParseECDSAToken(tokenResponse.Token, jwt.SigningMethodES256)
	assert.True(t, token.Valid)

	jsonClaims, _ := json.Marshal(tokenResponse.Claims)
	refreshTokenUserResponse := model.RefreshTokenUserResponse{}
	json.Unmarshal(jsonClaims, &refreshTokenUserResponse)
	assert.NotEmpty(t, refreshTokenUserResponse.Name)
	assert.NotEmpty(t, refreshTokenUserResponse.Email)
	assert.Equal(t, 1, refreshTokenUserResponse.Role)
}

func TestUserController_Show(t *testing.T) {
	// Setup request
	request := httptest.NewRequest("GET", "/api/v1/users/3", nil)

	// Test the request
	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	jsonData, _ := json.Marshal(webResponse.Data)
	getUserResponse := model.GetUserResponse{}
	json.Unmarshal(jsonData, &getUserResponse)
	assert.NotEmpty(t, getUserResponse.Name)
	assert.NotEmpty(t, getUserResponse.Email)
	assert.NotEmpty(t, getUserResponse.Nick)
}

func TestUserController_Update(t *testing.T) {
	// Setup form data
	token := "token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInNvY2lhbF9pZCI6IiIsImVtYWlsIjoiY29uc2VxdWF0dXJAZ21haWwuY29tIiwibmljayI6Im9jY2FlY2F0aSIsIm5hbWUiOiJlaXVzIiwicGljIjoiL2Fzc2V0cy9zdGF0aWMvdXNlci5wbmciLCJsb2NhdGlvbiI6IkluZG9uZXNpYSIsImRlc2MiOiIiLCJyb2xlIjoxfQ.8SoZMH7f_iReJ-YQmks92QAXo_TjRzrO3xtsdnD_6bvFcJiwXiBSmr1D_7oHEtjdkHNJpeWGCy_yhDiTlzZSXQ"
	data := strings.NewReader(token + "&name=Saiful Wicaksana&location=Jakarta&desc=Engineer")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/update/11", data)

	// Setup header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/x-www-form-urlencoded")

	// Test the request
	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	jsonData, _ := json.Marshal(webResponse.Data)
	updateUserResponse := model.UpdateUserResponse{}
	json.Unmarshal(jsonData, &updateUserResponse)
	assert.Equal(t, request.FormValue("name"), updateUserResponse.Name)
	assert.Equal(t, request.FormValue("location"), updateUserResponse.Location)
	assert.Equal(t, request.FormValue("desc"), updateUserResponse.Desc)
	assert.Equal(t, request.FormValue("desc"), updateUserResponse.Desc)
}

func TestUserController_Delete(t *testing.T) {
	// Setup form data
	data := strings.NewReader("token=eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwic29jaWFsX2lkIjoiIiwiZW1haWwiOiJxdWlAZ21haWwuY29tIiwibmljayI6InNpbnQiLCJuYW1lIjoibWluaW1hIiwicGljIjoiL2Fzc2V0cy9zdGF0aWMvdXNlci5wbmciLCJsb2NhdGlvbiI6IkluZG9uZXNpYSIsImRlc2MiOiIiLCJyb2xlIjoxfQ.fAT_hDCcUq7eL7CfI-Z6UIOhxTChJeeScR9-BcaoA86nxw_7z5RuSfZa59DFRJApabm-s1TeFSE92aw4w2pwng")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/delete/1", data)

	// Setup header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/x-www-form-urlencoded")

	// Test the request
	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test default json result
	responseBody, _ := io.ReadAll(response.Body)
	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Message)

	// Test response data
	jsonData, _ := json.Marshal(webResponse.Data)
	deleteUserResponse := model.DeleteUserResponse{}
	json.Unmarshal(jsonData, &deleteUserResponse)
	assert.Equal(t, uint(1), deleteUserResponse.Id)
}
