package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"greet/user-api/internal/config"
	"greet/user-rpc/userclient"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
