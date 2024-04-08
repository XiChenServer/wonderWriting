package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	userModel "calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookCommentLogic {
	return &LookCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看帖子的评论
func (l *LookCommentLogic) LookComment(in *community.LookCommentRequest) (*community.LookCommentResponse, error) {
	// todo: add your logic here and delete this line

	res, err := (&model.Comment{}).FindComment(l.svcCtx.DB, uint(in.PostId))
	if err != nil {
		return &community.LookCommentResponse{}, err
	}

	var comment []*community.CommentInfo

	//取出数据，进行转换
	for _, v := range *res {
		// 将 time.Time 转换为 Unix 时间戳 (int64)
		unixTime := v.CreatedAt.Unix()

		// 将 int64 类型的 Unix 时间戳转换为 int32 类型
		int32Time := int32(unixTime)
		newComment := community.CommentInfo{
			Id:         uint32(v.ID),
			CreateTime: int32Time,
			PostId:     uint32(v.PostID),
			Comment:    v.Content,
		}
		var User userModel.User
		err = l.svcCtx.DB.Where("user_id = ?", v.UserID).First(&User).Error
		if err != nil {
			fmt.Println(err, v.UserID)
			return nil, err
		}
		var userInfo = community.UserSimpleInfo{
			Id:          uint32(User.UserID),
			NickName:    User.Nickname,
			Account:     User.Account,
			AvatarImage: qiniu.ImgUrl + User.AvatarBackground,
		}
		newComment.UserInfo = &userInfo

		comment = append(comment, &newComment)
	}

	return &community.LookCommentResponse{
		CommentData: comment,
	}, nil
}
