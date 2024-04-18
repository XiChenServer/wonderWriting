package logic

import (
	groupModel "calligraphy/apps/grow/model"
	"calligraphy/apps/grow/rpc/internal/svc"
	"calligraphy/apps/grow/rpc/types/grow"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type CheckInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckInLogic {
	return &CheckInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckIn 执行打卡操作
func (l *CheckInLogic) CheckIn(in *grow.CheckInRequest) (*grow.CheckInResponse, error) {
	// 获取用户ID
	userID := uint(in.UserId)

	// 检查用户当天是否已经打卡
	hasCheckedIn, err := l.hasCheckedInToday(userID)
	if err != nil {
		return nil, err
	}

	// 如果当天已经打卡过，则返回错误信息
	if hasCheckedIn {
		return nil, errors.New("already checked in today")
	}

	// 更新打卡信息
	err = (&groupModel.CheckIn{}).UpdateCheckInInfo(l.svcCtx.DB, userID)
	if err != nil {
		return nil, err
	}
	// 创建书法记录
	recordContent, err := (&groupModel.RecordContent{}).CreateRecordContent(l.svcCtx.DB, uint(in.UserId), in.Content, in.Image, float64(in.Score))
	if err != nil {
		return nil, err
	}
	// 构造响应
	recordInfo := &grow.RecordSimpleInfo{
		RecordId:   uint32(recordContent.ID),
		UserId:     uint32(recordContent.UserID),
		Content:    recordContent.Content,
		Image:      recordContent.Image,
		Score:      float32(recordContent.Score),
		CreateTime: int32(recordContent.CreatedAt.Unix()),
	}
	// 构造响应
	return &grow.CheckInResponse{RecordInfo: recordInfo}, nil
}

// hasCheckedInToday 检查用户当天是否已经打卡
func (l *CheckInLogic) hasCheckedInToday(userID uint) (bool, error) {
	// 获取今天的起始时间和结束时间
	todayStart := time.Now().Truncate(24 * time.Hour)
	todayEnd := todayStart.Add(24 * time.Hour).Add(-time.Second)

	// 查询用户是否在今天已经打卡过
	var count int
	if err := l.svcCtx.DB.Model(&groupModel.CheckIn{}).Where("user_id = ? AND created_at >= ? AND created_at <= ?", userID, todayStart, todayEnd).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
