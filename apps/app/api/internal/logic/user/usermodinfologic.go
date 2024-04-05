package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModInfoLogic {
	return &UserModInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModInfoLogic) UserModInfo(req *types.UserModInfoRequest) (resp *types.UserModInfoResponse, err error) {
	// todo: add your logic here and delete this line

	//从jwt获取id
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	//调用rpc进行修改
	_, err = l.svcCtx.UserRpc.UserModInfo(l.ctx, &user.UserModInfoRequest{
		Id:       uid,
		NickName: req.NickName,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserModInfoResponse{}, nil
}
