package test

import (
	"encoding/json"
	"govel/app/model"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

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
	data := strings.NewReader("email=govel63bf26a363c2a@gmail.com&password=rahasia")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/login", data)

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
	loginUserResponse := model.LoginUserResponse{}
	json.Unmarshal(jsonData, &loginUserResponse)
	assert.NotEmpty(t, loginUserResponse.Name)
	assert.Equal(t, request.FormValue("email"), loginUserResponse.Email)
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
	data := strings.NewReader("name=Saiful Wicaksana&location=Jakarta&desc=Engineer")

	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/update/3", data)

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
	// Setup request
	request := httptest.NewRequest("POST", "/api/v1/users/delete/1", nil)

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
