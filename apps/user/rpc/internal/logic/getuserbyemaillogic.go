package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过邮箱查询
func (l *GetUserByEmailLogic) GetUserByEmail(in *user.EmailReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResp{}, nil
}
