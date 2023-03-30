package logic

import (
	"context"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.UserReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserModel2.Edit(l.ctx, req.Id, req.Name, req.Phone, req.Password)
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
