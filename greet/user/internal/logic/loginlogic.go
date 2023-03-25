package logic

import (
	"context"
	"github.com/pkg/errors"
	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Name)
	if err != nil {
		return nil, errors.Wrapf(errors.New(""), "HomestayOrderModel delete err : %+v", err)
	}
	logx.Info(user)

	return &types.LoginReply{
		Id:       user.Id,
		Name:     user.Name,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil
}
