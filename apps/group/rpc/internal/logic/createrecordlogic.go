package logic

import (
	groupModel "calligraphy/apps/group/model"
	"context"

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

// 上传书法记录
func (l *CreateRecordLogic) CreateRecord(in *group.CreateRecordRequest) (*group.CreateRecordResponse, error) {
	// 调用模型层方法创建记录
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
