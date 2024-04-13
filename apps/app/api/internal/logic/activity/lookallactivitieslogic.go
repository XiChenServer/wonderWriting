package activity

import (
	"calligraphy/apps/activity/rpc/types/activity"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookAllActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllActivitiesLogic {
	return &LookAllActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookAllActivitiesLogic) LookAllActivities(req *types.LookAllActivitiesRequest) (resp *types.LookAllActivitiesResponse, err error) {
	// todo: add your logic here and delete this line

	pageSize := req.Page_size
	pageSizeNum := 20 // 默认每页大小为20
	if pageSize != 0 {
		pageSizeNum = int(pageSize)
	}

	res, err := l.svcCtx.Activity.LookAllActivities(l.ctx, &activity.LookAllActivitiesRequest{
		Page:     req.Page,
		PageSize: uint32(pageSizeNum),
	})
	if err != nil {
		return &types.LookAllActivitiesResponse{}, err
	}
	var activityInfo []*types.ActivityInfo
	for _, v := range res.Activities {
		newInfo := types.ActivityInfo{
			Id:          v.Id,
			Name:        v.Name,
			Info:        v.ActivityInfo,
			Location:    v.Location,
			DateTime:    v.DateTime,
			Organizer:   v.Organizer,
			EndDateTime: v.EndDateTime,
			Duration:    v.Duration,
			RewardsInfo: v.RewardsInfo,
		}
		activityInfo = append(activityInfo, &newInfo)
	}
	return &types.LookAllActivitiesResponse{
		Activities:   activityInfo,
		Current_page: res.CurrentPage,
		Page_size:    res.PageSize,
		Offset:       res.Offset,
		Overflow:     res.Overflow,
		Total_pages:  res.TotalPages,
		Total_count:  res.TotalCount,
	}, nil
}
