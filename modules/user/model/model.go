package user_model

import "github.com/google/uuid"

type CreateUserSchema struct {
	FullName *string `json:"full_name,omitempty" validate:"required,min=4"`
	Age      *int    `json:"age,omitempty" validate:"number"`
	Username *string `json:"username,omitempty" validate:"required,min=4"`
	Password *string `json:"password,omitempty" validate:"required,min=4"`
}

type UpdateUserSchema struct {
	FullName *string    `json:"full_name,omitempty" validate:"omitempty,min=4"`
	Age      *int       `json:"age,omitempty" validate:"omitempty,number"`
	Username *string    `json:"username,omitempty" validate:"omitempty,min=4"`
	Password *string    `json:"password,omitempty" validate:"omitempty,min=4"`
}

type QueryUserSchema struct {
	Id       *uuid.UUID `json:"id,omitempty" validate:"omitempty,uuid4"`
	FullName *string    `json:"full_name,omitempty" validate:"omitempty"`
	Age      *int       `json:"age,omitempty" validate:"omitempty,number"`
	Username *string    `json:"username,omitempty" validate:"omitempty"`
}

type OrderUserSchema struct {
	FullName  *string `json:"full_name,omitempty" validate:"omitempty" enums:"asc,desc"`
	Age       *string `json:"age,omitempty" validate:"omitempty" enums:"asc,desc"`
	Username  *string `json:"username,omitempty" validate:"omitempty" enums:"asc,desc"`
	CreatedAt *string `json:"created_at,omitempty" validate:"omitempty" enums:"asc,desc"`
}
