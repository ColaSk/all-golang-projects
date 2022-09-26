package utils

import (
	"reflect"
	"strconv"
)

func Default(s interface{}) {
	// 设置结构体默认值

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	} else {
		panic("s must be ptr to struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		defaultTag := field.Tag.Get("default")

		switch value.Kind() {
		case reflect.String:
			value.SetString(defaultTag)
		case reflect.Int:
			if intV, err := strconv.ParseInt(defaultTag, 10, 64); err == nil {
				value.SetInt(intV)
			}
		}
	}
}
