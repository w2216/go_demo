package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.UserReq) (resp *types.UserResps, err error) {
	users, err := l.svcCtx.UserModel2.List(l.ctx, req.Id, req.Name, req.Phone, req.Password)
	if err != nil {
		logx.Info(err)
		return nil, err
	}

	var a []types.UserResp
	for _, user := range *users {
		var b types.UserResp
		_ = copier.Copy(&b, user)
		a = append(a, b)
	}

	return &types.UserResps{
		List: a,
	}, nil

}
