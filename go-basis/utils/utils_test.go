package utils_test

import (
	"fmt"
	"go-basis/utils"
	"testing"
)

type User struct {
	username string
	password string
	info     string
}

func SetUserInfo(info string) utils.DefaultExtInterface[User] {
	return utils.NewDefaultExts(func(user *User) {
		user.info = info
	})
}

func NewUser(username string, password string, defaultF ...utils.DefaultExtInterface[User]) *User {

	user := &User{
		username: username,
		password: password,
	}

	for _, f := range defaultF {
		f.Apply(user)
	}

	return user
}

func Test(t *testing.T) {
	user := NewUser("sun", "12343", SetUserInfo("test"))
	fmt.Println(*user)
}
