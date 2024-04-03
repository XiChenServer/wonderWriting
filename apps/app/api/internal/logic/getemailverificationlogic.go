package logic

import (
	"calligraphy/pkg/app_math"
	"calligraphy/pkg/verification"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailVerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmailVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailVerificationLogic {
	return &GetEmailVerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmailVerificationLogic) GetEmailVerification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	// todo: add your logic here and delete this line
	code := app_math.GenerateRandomNumber(6)
	err = verification.SendEmailVerificationCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return &types.VerificationResponse{
		Code:    200,
		Message: "Success",
	}, nil
}
