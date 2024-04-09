package mqs

import (
	"calligraphy/apps/app/api/internal/svc"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentSuccess) Consume(key, val string) error {
	fmt.Println("123")
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	return nil
}
