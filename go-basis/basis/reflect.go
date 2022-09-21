package basis

import (
	"fmt"
	"reflect"
)

// 反射

func ReflectTest() {
	//定义三个变量
	i := 1
	str := "test"
	l := make([]string, 1)

	l = append(l, "test")

	// TypeOf 返回接口中保存的值得类型，Typeof(nil)会返回nil

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(l))

	// ValueOf 返回一个初始化为interface接口保管的具体值得Value，ValueOf(nil)返回Value零值

	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(str))
	fmt.Println(reflect.ValueOf(l))

	setValueByReflect := func(x interface{}) {
		fmt.Println("x的type是: ", reflect.TypeOf(x))
		fmt.Println("x的value是: ", reflect.ValueOf(x).Elem())
		value := reflect.ValueOf(x)
		// 反射中使用Elem()方法获取指针所指向的值
		if value.Elem().Kind() == reflect.String {
			value.Elem().SetString("set test")
		} else if value.Elem().Kind() == reflect.Array {
			value.Elem().SetString("set test")
		}
		fmt.Println("x的value是: ", value.Elem())
	}
	// 反射修改值必须通过传递变量地址来修改。若函数传递的参数是值拷贝，则会发生下述错误
	setValueByReflect(&str)
	setValueByReflect(&l)
}
