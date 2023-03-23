package logic

import (
	"context"
	"greet/user-api/internal/svc"
	"greet/user-api/internal/types"
	"greet/user-rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.IdReq) (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line

	UserInfoReply, err := l.svcCtx.UserRpcClient.GetUser(l.ctx, &userclient.IdReq{
		Id: req.Id,
	})
	logx.Info(UserInfoReply)

	return &types.UserInfoReply{
		Id:     UserInfoReply.Id,
		Name:   UserInfoReply.Name,
		Number: UserInfoReply.Number,
		Gender: UserInfoReply.Gender,
	}, nil
}
