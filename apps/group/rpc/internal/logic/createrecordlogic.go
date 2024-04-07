package logic

import (
	groupModel "calligraphy/apps/group/model"
	"context"
	"time"

	"calligraphy/apps/group/rpc/internal/svc"
	"calligraphy/apps/group/rpc/types/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRecordLogic {
	return &CreateRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateRecord 上传书法记录
func (l *CreateRecordLogic) CreateRecord(in *group.CreateRecordRequest) (*group.CreateRecordResponse, error) {
	// 调用模型层方法创建记录
	hasCheckedIn, err := l.hasCheckedInToday(uint(in.UserId))
	if err != nil {
		return nil, err
	}

	// 如果当天还未上传过书法记录，则更新打卡信息
	if !hasCheckedIn {
		// 更新打卡信息
		err = (&groupModel.CheckIn{}).UpdateCheckInInfo(l.svcCtx.DB, uint(in.UserId))
		if err != nil {
			return nil, err
		}
	}

	// 创建书法记录
	recordContent, err := (&groupModel.RecordContent{}).CreateRecordContent(l.svcCtx.DB, uint(in.UserId), in.Content, in.Image, float64(in.Score))
	if err != nil {
		return nil, err
	}

	// 构造响应
	recordInfo := &group.RecordSimpleInfo{
		RecordId:   uint32(recordContent.ID),
		UserId:     uint32(recordContent.UserID),
		Content:    recordContent.Content,
		Image:      recordContent.Image,
		Score:      float32(recordContent.Score),
		CreateTime: int32(recordContent.CreatedAt.Unix()),
	}
	return &group.CreateRecordResponse{RecordInfo: recordInfo}, nil
}

// hasCheckedInToday 检查用户当天是否已经上传过书法记录
func (l *CreateRecordLogic) hasCheckedInToday(userId uint) (bool, error) {
	// 获取今天的起始时间和结束时间
	todayStart := time.Now().Truncate(24 * time.Hour)
	todayEnd := todayStart.Add(24 * time.Hour).Add(-time.Second)

	// 查询用户是否在今天已经上传过书法记录
	var count int
	if err := l.svcCtx.DB.Model(&groupModel.RecordContent{}).Where("user_id = ? AND created_at >= ? AND created_at <= ?", userId, todayStart, todayEnd).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
