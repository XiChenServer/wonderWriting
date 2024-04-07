package logic

import (
	groupModel "calligraphy/apps/group/model"
	"context"

	"calligraphy/apps/group/rpc/internal/svc"
	"calligraphy/apps/group/rpc/types/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookRecordByUserIdLogic {
	return &LookRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看某人的书法记录
func (l *LookRecordByUserIdLogic) LookRecordByUserId(in *group.LookRecordByUserIdRequest) (*group.LookRecordByUserIdResponse, error) {
	// todo: add your logic here and delete this line
	res, err := (&groupModel.RecordContent{}).LookAllRecordByOwn(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		return nil, err
	}
	var recordInfo []*group.RecordSimpleInfo
	for _, v := range res {
		newRecord := &group.RecordSimpleInfo{
			RecordId:   uint32(v.ID),
			UserId:     uint32(v.UserID),
			Content:    v.Content,
			Image:      v.Image,
			Score:      float32(v.Score),
			CreateTime: int32(v.CreatedAt.Unix()),
		}
		recordInfo = append(recordInfo, newRecord)
	}
	return &group.LookRecordByUserIdResponse{RecordInfo: recordInfo}, nil
}
