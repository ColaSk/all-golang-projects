package utils

import (
	"reflect"
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
		tag := field.Tag

		if value.Kind() != reflect.String {
			continue
		}
		if tag != "" {
			d := tag.Get("default")
			if d != "" {
				value.SetString(d)
			}
		}

	}

}
