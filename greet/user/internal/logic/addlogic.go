package logic

import (
	"context"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.UserReq) (resp *types.UserResp, err error) {

	user, err := l.svcCtx.UserModel2.Add(l.ctx, req.Name, req.Phone, req.Password)
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
}
