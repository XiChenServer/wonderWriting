package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModPwdLogic {
	return &UserModPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModPwdLogic) UserModPwd(req *types.UserModPwdRequset) (resp *types.UserModPwdResponse, err error) {
	// todo: add your logic here and delete this line
	//从jwt中获取id
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	//调用rpc
	_, err = l.svcCtx.UserRpc.UserModPwd(l.ctx, &user.UserModPwdRequest{
		Id:       uid,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserModPwdResponse{}, nil
}
