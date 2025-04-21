package utils

import (
	"encoding/json"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

func ReadFile(path string) any {
	var err error
	plan, _ := os.ReadFile(path)
	var data []map[string]interface{}
	err = json.Unmarshal(plan, &data)

	if err != nil {
		panic("ReadFile Error")
	}

	logrus.Infof("ReadFile = %T: %s\n", data, data)

	return data
}

func GetFieldInStruct(model any, fieldName string) reflect.Value {
	val := reflect.ValueOf(model)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	idField := val.FieldByName("ID")
	if !idField.IsValid() {
		panic("GetFieldInStruct Is Not Found Field")
	}

	return idField
}

func CheckIdExist(model any) bool {
	val := reflect.ValueOf(model)

	// ถ้าเป็น pointer → ต้องดึง value ที่ชี้อยู่
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// เช็คว่า model มี field ชื่อ "ID"
	idField := val.FieldByName("ID")
	if !idField.IsValid() {
		return false
	}

	// เช็คว่า field เป็น int/uint และมีค่า == 0
	switch idField.Kind() {
	case reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64, reflect.Uint32:
		return idField.Uint() == 0
	}

	return false
}
