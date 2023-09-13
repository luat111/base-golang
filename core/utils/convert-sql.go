package utils

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func ConvertOrderQuery(orderFields interface{}) []string {
	var orderBy []string

	for key, value := range StructToMapStringInterface(orderFields) {
		orderStr := key + " " + *value.(*string)
		orderBy = append(orderBy, orderStr)
	}

	return orderBy
}

func ConvertQueryString(baseQuery *gorm.DB, queryFields interface{}) {
	restMap := make(map[string]interface{})

	for key, value := range StructToMapStringInterface(queryFields) {
		fmt.Println(key, reflect.TypeOf(value).Kind())
		if key == "id" {
			restMap[key] = value
			continue
		}

		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			if reflect.ValueOf(value).Elem().Type().Kind() == reflect.String {
				queryStr := key + " ILIKE '" + *value.(*string) + "%'"
				baseQuery = baseQuery.Where(queryStr)
			} else {
				restMap[key] = value
			}
		}
	}

	baseQuery = baseQuery.Where(restMap)
}
