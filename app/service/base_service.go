package service

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

func DbHandleSelectField(field any) map[string]interface{} {
	fields := reflect.TypeOf(field)
	result := make(map[string]interface{})
	for i := 0; i < fields.NumField(); i++ {
		// Get the field
		field := fields.Field(i)

		// Get the json tag value
		jsonTag := field.Tag.Get("json")

		// Print the json tag value
		log.Infof("Field %d: %s\n", i+1, jsonTag)
		result[jsonTag] = ""
	}

	return result
}
