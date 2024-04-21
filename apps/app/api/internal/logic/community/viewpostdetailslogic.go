package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewPostDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewPostDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewPostDetailsLogic {
	return &ViewPostDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewPostDetailsLogic) ViewPostDetails(req *types.ViewPostDetailsRequest) (resp *types.ViewPostDetailsResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommunityRpc.ViewPostDetails(l.ctx, &community.ViewPostDetailsRequest{
		UserId: uint32(uid),
		PostId: req.PostId,
	})
	if err != nil {
		return &types.ViewPostDetailsResponse{}, err
	}
	postData := types.PostInfo{
		UserInfo: types.UserSimpleInfo{
			Id:          uint(res.PostData.UserInfo.Id),
			NickName:    res.PostData.UserInfo.NickName,
			Account:     res.PostData.UserInfo.Account,
			AvatarImage: res.PostData.UserInfo.AvatarImage,
		},
		Id:           uint(res.PostData.Id),
		UserId:       uint(res.PostData.UserId),
		ContentCount: uint(res.PostData.ContentCount),
		LikeCount:    uint(res.PostData.LikeCount),
		CollectCount: uint(res.PostData.CollectCount),
		Content:      res.PostData.Content,
		ImageUrls:    res.PostData.ImageUrls,
		CreateTime:   int32(res.PostData.CreateTime),
		DeleteTime:   int32(res.PostData.DeleteTime),
	}
	statusData := types.StatusWithPost{
		WhetherLike:    res.RelatedStatus.WhetherLike,
		WhetherCollect: res.RelatedStatus.WhetherCollect,
		WhetherFollow:  res.RelatedStatus.WhetherFollow,
	}
	if res.PostData.UserId == uint32(uid) {
		statusData.WhetherBelongOne = true
	}
	return &types.ViewPostDetailsResponse{
		PostData:   postData,
		StatusData: statusData,
	}, nil
}
