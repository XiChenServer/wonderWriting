package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllFansLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookAllFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllFansLogic {
	return &LookAllFansLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookAllFansLogic) LookAllFans(req *types.LookAllFansRequest) (resp *types.LookAllFansResponse, err error) {
	// todo: add your logic here and delete this line
	//调用rpc获取数据
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.LookAllFans(l.ctx, &user.LookAllFansRequest{
		UserId:   uint32(uid),
		Page:     req.Page,
		PageSize: pageSize,
	})
	if err != nil {
		return &types.LookAllFansResponse{}, err
	}
	var userInfo []*types.UserExhibitInfo
	for _, v := range res.UserInfo {
		newUserInfo := types.UserExhibitInfo{
			UserId:           v.UserId,
			AvatarBackground: v.AvatarBackground,
			NickName:         v.NickName,
			FollowCount:      v.FollowCount,
			FansCount:        v.FansCount,
			Email:            v.Email,
		}
		userInfo = append(userInfo, &newUserInfo)
	}
	return &types.LookAllFansResponse{
		UserData:    userInfo,
		CurrentPage: res.CurrentPage,
		PageSize:    res.PageSize,
		Offset:      res.Offset,
		Overflow:    res.Overflow,
		TotalPage:   res.TotalPages,
		TotalCount:  res.TotalCount,
	}, nil
}
