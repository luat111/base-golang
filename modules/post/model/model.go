package post_model

import (
	"time"

	"github.com/google/uuid"
)

type CreatePostSchema struct {
	Name *string `json:"name,omitempty" validate:"required,min=4" bson:"name,omitempty"`
}

type UpdatePostSchema struct {
	Name      *string   `json:"name,omitempty" validate:"omitempty,min=4" bson:"name,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type QueryPostSchema struct {
	Id   *uuid.UUID `json:"id,omitempty" validate:"omitempty,uuid4" bson:"id,omitempty"`
	Name *string    `json:"name,omitempty" validate:"omitempty" bson:"name,omitempty"`
}

type OrderPostSchema struct {
	Name *int `json:"name,omitempty" validate:"omitempty" enums:"0,1" bson:"name,omitempty"`
}
