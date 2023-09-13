package user_service

import (
	"errors"
	"net/http"
	. "practice/auth/core/base"
	"practice/auth/core/constants"
	. "practice/auth/core/interfaces"
	"practice/auth/core/utils"
	. "practice/auth/modules/user/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	BaseSqlService
	DB    *gorm.DB
	model User
}

func NewUserService(DB *gorm.DB) *UserService {
	BaseSqlService := NewBaseSqlService(DB)
	userService := &UserService{DB: DB, BaseSqlService: BaseSqlService, model: User{}}

	return userService
}

func (us *UserService) CreateUser(payload *CreateUserSchema) (err error, statusCode int) {
	newUser := &User{
		FullName: payload.FullName,
		Age:      payload.Age,
		Username: payload.Username,
		Password: payload.Password,
	}

	return us.Create(newUser)
}

func (us *UserService) GetListUser(
	filter QueryRequest[QueryUserSchema, OrderUserSchema],
	paginate PaginateRequest,
) (responseData PaginateResult, err error, statusCode int) {
	var records []User
	queryFilter := utils.FormatQueryRequest[QueryUserSchema, OrderUserSchema](filter)

	return us.GetList(us.model, records, queryFilter, paginate)
}

func (us *UserService) UpdateUser(id string, payload *UpdateUserSchema) (err error, statusCode int) {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err, http.StatusBadRequest
	}

	filter := &User{BaseModel: BaseModel{Id: uid}}
	mutationPayload := utils.StructToMapStringInterface(payload)

	return us.Update(us.model, filter, mutationPayload)
}

func (us *UserService) GetDetailUser(id string) (record interface{}, err error, statusCode int) {
	filter := utils.StructToMapStringInterface(&FilterById{Id: id})
	result, err, statusCode := us.GetOne(us.model, filter)

	result, ok := result.(*User)

	if !ok {
		return User{}, errors.New(constants.MessageNotFound), statusCode
	}

	return result.(*User), err, statusCode
}

func (us *UserService) DeleteUser(id string) (err error, statusCode int) {
	filter := utils.StructToMapStringInterface(&FilterById{Id: id})
	return us.Delete(us.model, filter)
}
