package basis

import (
	"errors"
	"fmt"
)

// 自定义错误
type MyErr struct {
	err string
}

// 实现错误接口
func (e *MyErr) Error() string {
	return e.err
}

func New(text string) error {
	return &MyErr{text}
}

func Exce() {

	defer func() {
		// 通过recover捕获异常
		if err := recover(); err != nil {
			fmt.Println("recover msg: ", err)
		} else {
			fmt.Println("recover ok")
		}
	}()

	// 默认方式生成一个error
	err := errors.New("err-info")
	fmt.Println(err)
	// 自定义myerr
	myerr := New("my-err-info")
	fmt.Println(myerr)
	// 通过panic抛出异常,并终止程序
	panic(myerr)

}
