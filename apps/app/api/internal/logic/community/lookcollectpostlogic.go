package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"
	"fmt"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookCollectPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookCollectPostLogic {
	return &LookCollectPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookCollectPostLogic) LookCollectPost() (resp *types.LookCollectPostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	//调用rpc进行查找数据
	res, err := l.svcCtx.CommunityRpc.LookCollectPost(l.ctx, &community.LookCollectPostRequest{
		UserId: uint32(uid),
	})
	if err != nil {
		return &types.LookCollectPostResponse{}, err
	}

	//进行转换数据
	var postData []*types.PostInfo
	for _, v := range res.PostData {
		var userInfo = types.UserSimpleInfo{
			Id:          uint(v.UserInfo.Id),
			NickName:    v.UserInfo.NickName,
			Account:     v.UserInfo.Account,
			AvatarImage: v.UserInfo.AvatarImage,
		}
		newPostData := &types.PostInfo{
			Id:           uint(v.Id),
			UserId:       uint(v.UserId),
			LikeCount:    uint(v.LikeCount),
			Content:      v.Content,
			ImageUrls:    v.ImageUrls,
			CreateTime:   int32(v.CreateTime),
			CollectCount: uint(v.CollectCount),
			ContentCount: uint(v.ContentCount),
			UserInfo:     userInfo,
		}
		fmt.Println(v.LikeCount)
		fmt.Println(newPostData.LikeCount)
		postData = append(postData, newPostData)
	}
	return &types.LookCollectPostResponse{PostInfo: postData}, nil
}
