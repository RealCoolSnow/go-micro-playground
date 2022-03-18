package svc

import (
	"go-zero-demo/mall/order/api/internal/config"
	"go-zero-demo/mall/order/api/internal/middleware"
	"go-zero-demo/mall/user/rpc/user"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
