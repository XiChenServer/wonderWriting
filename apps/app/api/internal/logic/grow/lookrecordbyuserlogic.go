package grow

import (
	"calligraphy/apps/grow/rpc/types/grow"
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
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	res, err := l.svcCtx.GrowRpc.LookRecordByUserId(l.ctx, &grow.LookRecordByUserIdRequest{
		UserId:   req.UserId,
		Page:     req.Page,
		PageSize: pageSize,
	})
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

	return &types.LookRecordByUserIdResponse{
		RecordInfo:  recordInfo,
		CurrentPage: res.CurrentPage,
		PageSize:    res.PageSize,
		Offset:      res.Offset,
		Overflow:    res.Overflow,
		TotalPage:   res.TotalPages,
		TotalCount:  res.TotalCount,
	}, nil
}
