package user

import (
	userService "practice/auth/modules/user/service"

	"gorm.io/gorm"
)

type UserModule struct {
	Controller *UserController
	Service    *userService.UserService
}

func InitUserModule(SqlDB *gorm.DB) *UserModule {
	userService := userService.NewUserService(SqlDB)
	userController := NewUserController(userService)

	return &UserModule{Controller: userController, Service: userService}
}
