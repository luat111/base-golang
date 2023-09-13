package base

import (
	"context"
	"errors"
	"net/http"

	"practice/auth/core/constants"
	. "practice/auth/core/interfaces"
	"practice/auth/core/utils"
	. "practice/auth/core/utils"

	"gorm.io/gorm"
)

type BaseSqlService interface {
	GetList(model, records any, filter QueryRequest[any, any], paginate PaginateRequest) (responseData PaginateResult, err error, statusCode int)
	GetMany(model, records any, filter QueryRequest[any, any]) (results any, err error, statusCode int)
	GetOne(model any, filter map[string]any) (results any, err error, statusCode int)
	Create(payload any) (err error, statusCode int)
	Update(model, filter any, payload map[string]any) (err error, statusCode int)
	Delete(model any, filter map[string]any) (err error, statusCode int)
}

type baseSqlService struct {
	DB *gorm.DB
}

func NewBaseSqlService(DB *gorm.DB) BaseSqlService {
	return &baseSqlService{DB: DB}
}

func (base *baseSqlService) GetList(
	model, records any,
	filter QueryRequest[any, any],
	paginate PaginateRequest,
) (responseData PaginateResult, err error, statusCode int) {
	iPage, iSize := PageSize(paginate.Page, paginate.Size)
	db := base.DB.WithContext(context.Background())

	var queryFields = filter.QueryFields
	var orderFields = filter.OrderFields

	if model == nil {
		return PaginateResult{}, errors.New("Not found model"), http.StatusInternalServerError
	}

	query := db.Model(&model)
	if queryFields != nil {
		ConvertQueryString(query, queryFields)
	}

	if orderFields != nil {
		orderFields = utils.ConvertOrderQuery(orderFields)
		orderByArr := orderFields.([]string)

		for i := 0; i < len(orderByArr); i++ {
			query = query.Order(orderByArr[i])
		}
	}

	data := query.Limit(iSize).Offset((iPage - 1) * iSize).Find(&records)
	total, err := data.RowsAffected, data.Error

	if err != nil {
		return PaginateResult{}, err, http.StatusInternalServerError
	}

	return PaginateResult{
		Results: records,
		Pagination: &PaginateTotal{
			Total: total,
		},
	}, nil, http.StatusOK
}

func (base *baseSqlService) GetMany(
	model, records any,
	filter QueryRequest[any, any],
) (results any, err error, statusCode int) {
	db := base.DB.WithContext(context.Background())
	var queryFields = filter.QueryFields
	var orderFields = filter.OrderFields.([]string)

	if model == nil {
		return []any{}, errors.New("Not found model"), http.StatusInternalServerError
	}

	query := db.Model(&model)
	if queryFields != nil {
		ConvertQueryString(query, queryFields)
	}

	if orderFields != nil {
		for i := 0; i < len(orderFields); i++ {
			query = query.Order(orderFields[i])
		}
	}

	err = db.Find(&records).Error

	if err != nil {
		return []any{}, err, http.StatusInternalServerError
	}

	return records, nil, http.StatusOK
}

func (base *baseSqlService) GetOne(model any, filter map[string]any) (record any, err error, statusCode int) {
	db := base.DB.WithContext(context.Background())
	if model == nil {
		return nil, errors.New("Not found model"), http.StatusInternalServerError
	}

	query := db.Model(&model)
	if filter != nil {
		ConvertQueryString(query, filter)
	}

	err = query.First(&model).Error

	if err != nil {
		return nil, errors.New(constants.MessageNotFound), http.StatusNotFound
	} else {
		return model, nil, http.StatusOK
	}
}

func (base *baseSqlService) Create(model any) (err error, statusCode int) {
	db := base.DB.WithContext(context.Background())

	err = db.Create(model).Error
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (base *baseSqlService) Update(model, filter any, payload map[string]any) (err error, statusCode int) {
	db := base.DB.WithContext(context.Background())

	err = db.Where(filter).Take(&filter).Error
	if err != nil {
		return err, http.StatusNotFound
	}

	err = db.Model(filter).Updates(payload).Error
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (base *baseSqlService) Delete(model any, filter map[string]any) (err error, statusCode int) {
	db := base.DB.WithContext(context.Background())

	record, err, statusCode := base.GetOne(model, filter)

	if record == nil {
		return err, statusCode
	} else {
		err = db.Delete(record).Error
		if err != nil {
			return err, http.StatusInternalServerError
		}

		return nil, http.StatusOK
	}
}
