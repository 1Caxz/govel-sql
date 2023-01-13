# Govel
Go clean architecture adopted from Laravel.

Main layer: Entity -> Repository -> Service -> Controller

## Features
- [x] Support `mysql`, `postgres`, `sqlite`, `sqlserver` databases
- [x] Support ORM database
- [x] Support database migration
- [x] Support fake data
- [x] Support JWT

## Main Packages
- [x] Gorm: The fantastic ORM library for Golang, aims to be developer friendly. `github.com/go-gorm/gorm`
- [x] Fiber: The fastest HTTP engine for Go. `github.com/gofiber/fiber`

## Builds
Build project based on your system:
- `make linux-build` for linux system
- `make mac-build` for macOs darwin system
- `make migrate-linux-build` build migrate only for linux system
- `make migrate-mac-build` build migrate only for macOs darwin system

## Database Migration
After build, use these commands:
- Use command `./migrate start` to start migration
- Use command `./migrate seed` to create fake data

## Declaring Models
Govel using `gorm` package to manage the database. Please follow this docs for more https://gorm.io/docs/models.html
```go
type User struct {
	ID              uint   `gorm:"primaryKey"`
	SocialId        string `gorm:"type:varchar(255);unique;default:null"`
	Email           string `gorm:"type:varchar(255);unique;not null"`
	Password        string `gorm:"type:varchar(255);default:null"`
	EmailVerifiedAt *time.Time
	Nick            string `gorm:"type:varchar(50);unique;not null"`
	Name            string `gorm:"type:varchar(255);index:,class:FULLTEXT;not null"`
	Pic             string `gorm:"type:varchar(255);not null;default:/assets/static/user.png"`
	Location        string `gorm:"type:varchar(255);default:Indonesia"`
	Desc            string `gorm:"type:varchar(255);default:null"`
	Role            int    `gorm:"type:tinyint(2);default:1"`
	Status          int    `gorm:"type:tinyint(2);default:0"`
	ApiToken        string `gorm:"type:varchar(80);default:null"`
	RememberToken   string `gorm:"type:varchar(100);default:null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
```

## Migration
Register your table entity to the migration `database/migration/migrator.go`.
```go
func Migrator(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.AnotherTable{})
}
```
Build the project and run command `migration start`.

## Seeder With Faker
Setup your fake data `database/seeder/seeder.go`.
```go
func Seeder(db *gorm.DB) {
	for i := 0; i < 30; i++ {
		hashed, err := bcrypt.GenerateFromPassword([]byte(faker.Word()), bcrypt.DefaultCost)
		exception.PanicIfNeeded(err)
		db.Create(&entity.User{
			Email:    faker.Word() + "@gmail.com",
			Password: string(hashed),
			Name:     faker.Word(),
			Nick:     faker.Word(),
			Role:     1,
			Status:   1,
		})
	}
}
```

## Route
Like laravel, you can add your route in `route/api.go` or `route/web.go`.

For more follow this docs `https://docs.gofiber.io/guide/routing`
```go
route.Post("/create", userController.Create)
route.Get("/show/:id", userController.Show)
```

Or you can register your controller directly:
```go
func APIRoute(route fiber.Router, database *gorm.DB) {
	// Setup Repository
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)
	userController.Route(route)
}
```
And define your route from controller:
```go
func (controller *UserController) Route(route fiber.Router) {
	group := route.Group("/users")
	group.Post("/all", controller.All)
	group.Post("/create", controller.Create)
	group.Post("/update/:id", controller.Update)
	group.Post("/delete/:id", controller.Delete)
}
```

## Middleware
There are 3 default middlewares:
- [x] APPMiddleware: Used by all routes including api and web.
- [x] APIMiddleware: Used by api route only.
- [x] WebMiddleware: Used by web route only.

You can create your own middleware and use it:
```go
// Use middleware on route under /users path
route.Use("/users", middleware.Authenticate)

// Use middleware on specific route
route.Post("/update/:id", middleware.Authenticate, controller.Update)
```
For more follow this docs `https://docs.gofiber.io/guide/routing#middleware`

## JWT
There is helper `jwt.go` that you can use to create or parse the token:
```go
// Create token
type LoginUserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	jwt.StandardClaims
}
data :=  model.LoginUserResponse{
	Id:       user.ID,
	Email:    user.Email,
	Name:     user.Name,
}
token := helper.MakeECDSAToken(&data, jwt.SigningMethodES256)

// Parsing the token
token := helper.ParseECDSAToken(request.Token, jwt.SigningMethodES256)
claims := token.Claims.(jwt.MapClaims)
email := claims["email"].(string)
```
For more follow this docs `https://github.com/golang-jwt/jwt`

## Access Database from Controller
You can access the database object from fiber context in the controller directly. But this is not recommended.
```go
func (ctx *UserController) Example(c *fiber.Ctx) error {
	// Access the database object from fiber context
	db := c.Locals("DB").(*gorm.DB)

	result := map[string]interface{}{}
	db.Table("users").Where("id = ?", 1).Take(&result)
	return c.Status(200).JSON(result)
}
```