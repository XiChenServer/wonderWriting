package activity

import (
	"calligraphy/apps/activity/rpc/types/activity"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignUpActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSignUpActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignUpActivityLogic {
	return &UserSignUpActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignUpActivityLogic) UserSignUpActivity(req *types.UserSignUpActivityRequest) (resp *types.UserSignUpActivityResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.Activity.UserSignUpActiity(l.ctx, &activity.UserSignUpActivityRequest{
		UserId:     req.User_id,
		ActivityId: req.Activity_id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserSignUpActivityResponse{}, nil
}
