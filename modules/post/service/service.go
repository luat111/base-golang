package post_service

import (
	"net/http"
	. "practice/auth/core/base"
	. "practice/auth/core/interfaces"
	. "practice/auth/core/utils"
	. "practice/auth/modules/post/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostService struct {
	BaseSchemaLessService
	DB    *mongo.Database
	model string
}

func NewPostService(DB *mongo.Database) *PostService {
	BaseSqlService := NewBaseSchemaLessService(DB)
	postService := &PostService{DB: DB, BaseSchemaLessService: BaseSqlService, model: ToLowerCase(GetStructName(Post{}))}

	return postService
}

func (ps *PostService) CreatePost(payload *CreatePostSchema) (err error, statusCode int) {
	now := time.Now()
	newPost := &Post{
		Name:      payload.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return ps.CreateSchemaLess(ps.model, newPost)
}

func (ps *PostService) GetListPost(
	filter QueryRequest[QueryPostSchema, OrderPostSchema],
	paginate PaginateRequest,
) (responseData PaginateResult, err error, statusCode int) {
	queryFilter := FormatQueryRequest[QueryPostSchema, OrderPostSchema](filter)
	var records []Post
	return ps.GetListSchemaLess(ps.model, records, queryFilter, paginate)
}

func (ps *PostService) UpdatePost(id string, payload *UpdatePostSchema) (err error, statusCode int) {
	objId, err := ConvertToObjectId(id)

	if err != nil {
		return err, http.StatusBadRequest
	}

	filter := &Post{Id: objId}
	payload.UpdatedAt = time.Now()

	return ps.UpdateSchemaLess(ps.model, filter, payload)
}

func (ps *PostService) GetDetailPost(id string) (record interface{}, err error, statusCode int) {
	objId, err := ConvertToObjectId(id)

	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	filter := &FilterByObjectId{Id: objId}
	result, err, statusCode := ps.GetOneSchemaLess(ps.model, filter)

	return result, err, statusCode
}

func (ps *PostService) DeletePost(id string) (err error, statusCode int) {
	objId, err := ConvertToObjectId(id)

	if err != nil {
		return err, http.StatusBadRequest
	}

	filter := &FilterByObjectId{Id: objId}

	return ps.DeleteSchemaLess(ps.model, filter)
}
