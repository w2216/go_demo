package logic

import (
	"context"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.UserReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserModel2.Del(l.ctx, req.Id)
	if err != nil {
		logx.Info(err)
		return nil, err
	}

	return &types.UserResp{
		Id:       user.Id,
		Name:     user.Name,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil

	return
}
