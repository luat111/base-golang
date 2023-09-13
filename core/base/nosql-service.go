package base

import (
	"context"
	"net/http"
	. "practice/auth/core/interfaces"
	. "practice/auth/core/utils"
	"reflect"

	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseSchemaLessService interface {
	GetListSchemaLess(model string, records any, filter QueryRequest[any, any], paginate PaginateRequest) (responseData PaginateResult, err error, statusCode int)
	GetManySchemaLess(model string, records any, filter QueryRequest[any, any]) (results any, err error, statusCode int)
	GetOneSchemaLess(model string, filter any) (results map[string]any, err error, statusCode int)
	CreateSchemaLess(model string, payload any) (err error, statusCode int)
	UpdateSchemaLess(model string, filter, payload any) (err error, statusCode int)
	DeleteSchemaLess(model string, filter any) (err error, statusCode int)
}

type baseSchemaLessService struct {
	DB *mongo.Database
}

func NewBaseSchemaLessService(DB *mongo.Database) BaseSchemaLessService {
	return &baseSchemaLessService{DB: DB}
}

func (base *baseSchemaLessService) GetListSchemaLess(
	model string,
	records any,
	filter QueryRequest[any, any],
	paginate PaginateRequest,
) (responseData PaginateResult, err error, statusCode int) {
	iPage, iSize := PageSize(paginate.Page, paginate.Size)
	db := base.DB.Collection(model)

	var queryFields = filter.QueryFields
	var orderFields = filter.OrderFields
	var otps *options.FindOptions = options.Find()

	if queryFields != nil {
		queryFields, err = ToBson(queryFields)

		if err != nil {
			return PaginateResult{}, err, http.StatusInternalServerError
		}
	}

	if orderFields != nil {
		orderFields, err = ToBson(orderFields)
		otps = otps.SetSort(orderFields)

		if err != nil {
			return PaginateResult{}, err, http.StatusInternalServerError
		}
	}

	otps.SetLimit(int64(iSize)).SetSkip(int64((iPage - 1) * iSize))
	cntOpts := options.Count().SetHint("_id_")

	wg := sync.WaitGroup{}
	wg.Add(2)

	totalCh := make(chan int64, 1)
	recordsCh := make(chan any, 1)
	errCh := make(chan error, 2)

	defer close(totalCh)
	defer close(recordsCh)
	defer close(errCh)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		cusor, err := db.Find(context.Background(), queryFields, otps)

		if err != nil {
			errCh <- err
			recordsCh <- nil
		} else {
			records, err = GetDataFromCusor(cusor)

			if err != nil {
				errCh <- err
				recordsCh <- nil
			}

			errCh <- nil
			recordsCh <- records
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		count, _ := db.CountDocuments(context.Background(), queryFields, cntOpts)
		totalCh <- count

	}(&wg)

	err = <-errCh
	total := <-totalCh
	records = <-recordsCh

	wg.Wait()

	if err != nil {
		return PaginateResult{}, err, http.StatusInternalServerError
	}

	return PaginateResult{
		Pagination: &PaginateTotal{
			Total: total,
		},
		Results: records,
	}, nil, http.StatusOK
}

func (base *baseSchemaLessService) GetManySchemaLess(
	model string,
	records any,
	filter QueryRequest[any, any],
) (results any, err error, statusCode int) {
	db := base.DB.Collection(model)

	var queryFields = filter.QueryFields
	var orderFields = filter.OrderFields
	var otps *options.FindOptions = options.Find()

	if queryFields != nil {
		queryFields, err = ToBson(queryFields)

		if err != nil {
			return []any{}, err, http.StatusInternalServerError
		}
	}

	if orderFields != nil {
		orderFields, err = ToBson(orderFields)
		otps = otps.SetSort(orderFields)

		if err != nil {
			return []any{}, err, http.StatusInternalServerError
		}
	}

	cusor, err := db.Find(context.Background(), queryFields, otps)
	err = cusor.Decode(&results)

	if err != nil {
		return []any{}, err, http.StatusInternalServerError
	}

	return results, nil, http.StatusOK
}

func (base *baseSchemaLessService) GetOneSchemaLess(
	model string,
	filter any,
) (record map[string]any, err error, statusCode int) {
	db := base.DB.Collection(model)
	var otps *options.FindOneOptions = options.FindOne()

	if filter != nil {
		filter, err = ToBson(filter)

		if err != nil {
			return make(map[string]any), err, http.StatusInternalServerError
		}
	}
	err = db.FindOne(context.Background(), filter, otps).Decode(&record)

	if err != nil {
		return make(map[string]any), err, http.StatusInternalServerError
	}

	return record, nil, http.StatusOK
}

func (base *baseSchemaLessService) CreateSchemaLess(model string, payload any) (err error, statusCode int) {
	db := base.DB.Collection(model)

	if reflect.TypeOf(payload).Kind() == reflect.Slice {
		_, err = db.InsertMany(context.Background(), payload.([]any))
	} else {
		_, err = db.InsertOne(context.Background(), payload)
	}

	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

func (base *baseSchemaLessService) UpdateSchemaLess(
	model string,
	filter, payload any,
) (err error, statusCode int) {
	db := base.DB.Collection(model)

	filter, err = ToBson(filter)

	if err != nil {
		return err, http.StatusInternalServerError
	}

	payload, err = ToBson(payload)
	if err != nil {
		return err, http.StatusInternalServerError
	} else {
		payload = bson.D{{Key: "$set", Value: payload}}
	}

	if reflect.TypeOf(payload).Kind() == reflect.Slice {
		_, err = db.UpdateMany(context.Background(), filter, payload)
	} else {
		_, err = db.UpdateOne(context.Background(), filter, payload)
	}

	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (base *baseSchemaLessService) DeleteSchemaLess(
	model string, filter any,
) (err error, statusCode int) {
	db := base.DB.Collection(model)

	filter, err = ToBson(filter)

	if err != nil {
		return err, http.StatusInternalServerError
	}

	_, err = db.DeleteMany(context.Background(), filter)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
