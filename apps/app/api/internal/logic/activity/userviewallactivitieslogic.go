package activity

import (
	"calligraphy/apps/activity/rpc/types/activity"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserViewAllActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserViewAllActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserViewAllActivitiesLogic {
	return &UserViewAllActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserViewAllActivitiesLogic) UserViewAllActivities(req *types.UserViewAllActivitiesRequest) (resp *types.UserViewAllActivitiesResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.Activity.UserViewAllActivities(l.ctx, &activity.UserViewAllActivitiesRequest{
		UserId:   uint32(uid),
		Page:     req.Page,
		PageSize: req.Page_size,
	})

	if err != nil {
		return &types.UserViewAllActivitiesResponse{}, err
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
	return &types.UserViewAllActivitiesResponse{
		Activities:   activityInfo,
		Current_page: res.CurrentPage,
		Page_size:    res.PageSize,
		Offset:       res.Offset,
		Overflow:     res.Overflow,
		Total_pages:  res.TotalPages,
		Total_count:  res.TotalCount,
	}, nil
}
