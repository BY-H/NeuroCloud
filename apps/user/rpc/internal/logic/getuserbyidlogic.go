package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过用户ID查询
func (l *GetUserByIdLogic) GetUserById(in *user.IdReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResp{}, nil
}
