package group

import (
	"calligraphy/apps/group/rpc/types/group"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookRecordByUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookRecordByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookRecordByUserLogic {
	return &LookRecordByUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookRecordByUserLogic) LookRecordByUser(req *types.LookRecordByUserIdRequest) (resp *types.LookRecordByUserIdResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.GroupRpc.LookRecordByUserId(l.ctx, &group.LookRecordByUserIdRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	var recordInfo []*types.RecordSimpleInfo
	for _, v := range res.RecordInfo {
		newRecordInfo := &types.RecordSimpleInfo{
			RecordId:   v.RecordId,
			UserId:     v.UserId,
			Content:    v.Content,
			Image:      v.Image,
			Score:      v.Score,
			CreateTime: v.CreateTime,
		}
		recordInfo = append(recordInfo, newRecordInfo)
	}

	return &types.LookRecordByUserIdResponse{RecordInfo: recordInfo}, nil
}
