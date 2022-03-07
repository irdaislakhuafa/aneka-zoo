package helpers

import (
	"reflect"
)

func GetFields(data interface{}) map[string]interface{} {
	dataType := reflect.TypeOf(data)
	// sliceFields := make([]string, dataType.NumField())
	mapFields := make(map[string]interface{})
	fieldName := ""

	for i := 0; i < dataType.NumField(); i++ {
		fieldName = dataType.Field(i).Tag.Get("json")
		if fieldName != "id" {
			mapFields[fieldName] = dataType.Field(i).Type.Name()
		}
	}
	return mapFields
}
