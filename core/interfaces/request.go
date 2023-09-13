package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaginateRequest struct {
	Page int `json:"page,omitempty" validate:"number,min=1"`
	Size int `json:"size,omitempty" validate:"number,min=1"`
}

type QueryRequest[T, R interface{}] struct {
	QueryFields T `json:"queryFields,omitempty"`
	OrderFields R `json:"orderFields,omitempty"`
}

type FilterByIds struct {
	Ids []string `json:"ids,omitempty"`
}

type FilterById struct {
	Id string `json:"id,omitempty"`
}

type FilterByObjectId struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
}
