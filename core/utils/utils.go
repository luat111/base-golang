package utils

import (
	. "practice/auth/core/constants"
	"practice/auth/core/interfaces"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func PageSize(page, size any) (iPage, iSize int) {
	iPage = ToNumber(page)

	if iPage == 0 {
		iPage = DefaultPage
	}

	iSize = ToNumber(size)
	if iSize == 0 {
		iSize = DefaultPageSize
	}

	return iPage, iSize
}

func StructToMapStringInterface(data interface{}) map[string]interface{} {
	var result map[string]interface{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &result,
	})

	if err != nil {
		panic(err)
	}

	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}

	return result
}

func FormatQueryRequest[T, R any](queryRequest interfaces.QueryRequest[T, R]) interfaces.QueryRequest[any, any] {
	var temp interfaces.QueryRequest[any, any]

	temp.QueryFields = queryRequest.QueryFields
	temp.OrderFields = queryRequest.OrderFields

	return temp
}

func GetStructName(model interface{}) string {
	if t := reflect.TypeOf(model); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ToLowerCase(str string) string {
	return strings.ToLower(str)
}

func ConvertByteToString(bytes []byte) string {
	return string(bytes[:])
}