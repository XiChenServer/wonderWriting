package logic

import (
	activityModel "calligraphy/apps/activity/model"
	"context"
	"time"

	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignUpActiityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSignUpActiityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignUpActiityLogic {
	return &UserSignUpActiityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserSignUpActiity 用户进行报名
func (l *UserSignUpActiityLogic) UserSignUpActiity(in *activity.UserSignUpActivityRequest) (*activity.UserSignUpActivityResponse, error) {
	// 获取请求中的用户ID和活动ID
	userID := in.UserId
	activityID := in.ActivityId
	// 查询活动信息，获取活动开始时间和结束时间
	activityInfo, err := (&activityModel.Activity{}).GetActivityInfo(l.svcCtx.DB, uint(activityID))
	if err != nil {
		return nil, err
	}

	endDateTimeStr := activityInfo.EndDateTime
	endDateTime, err := time.Parse("2006-01-02 15:04:05", endDateTimeStr)
	if err != nil {
		// 处理解析时间出错的情况
		return nil, err
	}

	// 现在 endDateTime 是时间类型，可以与当前时间进行比较
	currentTime := time.Now()
	if currentTime.After(endDateTime) {
		// 报名时间已经过了，返回报名失败的响应
		response := &activity.UserSignUpActivityResponse{
			Success: false,
			Message: "报名时间已过",
		}
		return response, nil
	}

	// 检查用户是否已经报名该活动
	isSignedUp, err := (&activityModel.UserSignUpActivity{}).CheckUserSignUp(l.svcCtx.DB, uint(userID), uint(activityID))
	if err != nil {
		return nil, err
	}

	// 如果用户已经报名，则返回报名失败的响应
	if isSignedUp {
		response := &activity.UserSignUpActivityResponse{
			Success: false,
			Message: "用户已经报名该活动",
		}
		return response, nil
	}

	// 创建报名记录
	err = (&activityModel.UserSignUpActivity{}).CreateSignUpRecord(l.svcCtx.DB, uint(userID), uint(activityID))
	if err != nil {
		return nil, err
	}

	// 返回报名成功的响应
	response := &activity.UserSignUpActivityResponse{
		Success: true,
		Message: "报名成功",
	}
	return response, nil
}
