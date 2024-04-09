package activity

import (
	"calligraphy/apps/activity/rpc/types/activity"
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type GrabPointsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrabPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrabPointsLogic {
	return &GrabPointsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 是否已经更新过 PointByOneDay 的标志
var pointUpdated bool

func (l *GrabPointsLogic) GrabPoints(req *types.GrabPointsRequest) (resp *types.GrabPointsResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	//// 每天10点开始活动
	//startTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Local)
	//// 每天10点10分结束活动
	//endTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 10, 0, 0, time.Local)
	//
	//// 计算距离活动开始的时间
	//durationUntilStart := startTime.Sub(time.Now())
	//
	//// 设置定时器，在活动开始时执行更新操作
	//timer := time.NewTimer(durationUntilStart)
	//<-timer.C // 等待定时器到期
	//
	//// 活动结束后关闭定时器
	//defer timer.Stop()
	//
	//// 如果当前时间已经超过了活动结束时间，则不执行后续操作
	//if time.Now().After(endTime) {
	//	return nil, nil // 活动已结束，不执行后续操作
	//}

	// todo: add your logic here

	// 调用 GrabPointsRequest 方法，进行积分抢夺
	_, err = l.svcCtx.Activity.GrabPoints(l.ctx, &activity.GrabPointsRequest{UserId: uint32(uid)})
	if err != nil {
		return nil, err
	}

	return
}
