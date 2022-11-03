package go_property_mapper

import (
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"
)

func Transform(source interface{}, dest interface{}) (obj interface{}, err error) {
	defer func() {
		r := recover()
		if r != nil {
			obj = nil
			err = errors.New(fmt.Sprintf("%v \n %s", r, string(debug.Stack())))
		}
	}()
	obj = setDestValues(source, dest)
	return
}

func setDestValues(source interface{}, dest interface{}) interface{} {
	destFields := getFieldNames(dest)
	sv := reflect.ValueOf(source).Elem()
	dv := reflect.ValueOf(dest).Elem()
	dvt := reflect.TypeOf(dest).Elem()
	for i := 0; i < sv.NumField(); i++ {
		field := sv.Type().Field(i)
		if id, exists := destFields[field.Name]; exists {
			//fmt.Println(fmt.Sprintf("%v, %v, %v, %v", dv.Field(id).Type(), dv.Field(id), reflect.ValueOf(sv.FieldByName(field.Name)), dv.Field(id).CanSet()))
			if field.Type == dv.Type().Field(id).Type {
				if dv.Field(id).CanSet() {
					fieldVal := sv.FieldByName(field.Name).Interface()
					dv.Field(id).Set(reflect.ValueOf(fieldVal))
				}
			} else if dv.Field(id).Kind() == reflect.Ptr && sv.Field(i).Kind() == reflect.Ptr {
				s, _ := dvt.FieldByName(field.Name)
				obj := setDestValues(sv.FieldByName(field.Name).Interface(), reflect.New(s.Type.Elem()).Interface())
				dv.Field(id).Set(reflect.ValueOf(obj))
			}
		}
	}
	return dest
}

func getFieldNames(obj interface{}) map[string]int {
	fields := make(map[string]int, 0)
	e := reflect.ValueOf(obj).Elem()
	for i := 0; i < e.NumField(); i++ {
		fields[e.Type().Field(i).Name] = i
	}
	return fields
}
