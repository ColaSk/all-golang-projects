// package svc

// import (
// 	"go-zero-teemplate/order/api/internal/config"
// )

// type ServiceContext struct {
// 	Config config.Config
// }

// func NewServiceContext(c config.Config) *ServiceContext {
// 	return &ServiceContext{
// 		Config: c,
// 	}
// }

package svc

import (
	"go-zero-teemplate/mall/order/api/internal/config"
	userclient "go-zero-teemplate/mall/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
