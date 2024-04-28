// Code generated by goctl. DO NOT EDIT.
// Source: community.proto

package communityclient

import (
	"context"

	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancelLikeCommentRequest           = community.CancelLikeCommentRequest
	CancelLikeCommentResponse          = community.CancelLikeCommentResponse
	CommentInfo                        = community.CommentInfo
	CommunityCancelCollectPostRequest  = community.CommunityCancelCollectPostRequest
	CommunityCancelCollectPostResponse = community.CommunityCancelCollectPostResponse
	CommunityCancelContentPostRequest  = community.CommunityCancelContentPostRequest
	CommunityCancelContentPostResponse = community.CommunityCancelContentPostResponse
	CommunityCancelLikePostRequest     = community.CommunityCancelLikePostRequest
	CommunityCancelLikePostResponse    = community.CommunityCancelLikePostResponse
	CommunityCollectPostRequest        = community.CommunityCollectPostRequest
	CommunityCollectPostResponse       = community.CommunityCollectPostResponse
	CommunityContentPostRequest        = community.CommunityContentPostRequest
	CommunityContentPostResponse       = community.CommunityContentPostResponse
	CommunityCreatePostRequest         = community.CommunityCreatePostRequest
	CommunityCreatePostResponse        = community.CommunityCreatePostResponse
	CommunityDelPostRequest            = community.CommunityDelPostRequest
	CommunityDelPostResponse           = community.CommunityDelPostResponse
	CommunityLikePostRequest           = community.CommunityLikePostRequest
	CommunityLikePostResponse          = community.CommunityLikePostResponse
	CommunityLookAllPostsRequest       = community.CommunityLookAllPostsRequest
	CommunityLookAllPostsResponse      = community.CommunityLookAllPostsResponse
	CommunityLookPostByOwnRequest      = community.CommunityLookPostByOwnRequest
	CommunityLookPostByOwnResponses    = community.CommunityLookPostByOwnResponses
	LikeCommentRequest                 = community.LikeCommentRequest
	LikeCommentResponse                = community.LikeCommentResponse
	LookCollectPostRequest             = community.LookCollectPostRequest
	LookCollectPostResponse            = community.LookCollectPostResponse
	LookCommentRequest                 = community.LookCommentRequest
	LookCommentResponse                = community.LookCommentResponse
	LookReplyCommentRequest            = community.LookReplyCommentRequest
	LookReplyCommentResponse           = community.LookReplyCommentResponse
	PostInfo                           = community.PostInfo
	ReplyCommentInfo                   = community.ReplyCommentInfo
	ReplyCommunityRequest              = community.ReplyCommunityRequest
	ReplyCommunityResponse             = community.ReplyCommunityResponse
	StatusWithPost                     = community.StatusWithPost
	UserSimpleInfo                     = community.UserSimpleInfo
	ViewPostDetailsRequest             = community.ViewPostDetailsRequest
	ViewPostDetailsResponse            = community.ViewPostDetailsResponse
	ViewTheLatestPostRequest           = community.ViewTheLatestPostRequest
	ViewTheLatestPostResponse          = community.ViewTheLatestPostResponse
	ViewUnreadCommentsCountRequest     = community.ViewUnreadCommentsCountRequest
	ViewUnreadCommentsCountResponse    = community.ViewUnreadCommentsCountResponse
	ViewUnreadCommentsRequest          = community.ViewUnreadCommentsRequest
	ViewUnreadCommentsResponse         = community.ViewUnreadCommentsResponse
	WhetherCollectPostRequest          = community.WhetherCollectPostRequest
	WhetherCollectPostResponse         = community.WhetherCollectPostResponse
	WhetherLikePostRequest             = community.WhetherLikePostRequest
	WhetherLikePostResponse            = community.WhetherLikePostResponse

	Community interface {
		// 查询最新的帖子
		ViewTheLatestPost(ctx context.Context, in *ViewTheLatestPostRequest, opts ...grpc.CallOption) (*ViewTheLatestPostResponse, error)
		// 查看用户有多少未读的信息
		ViewUnreadCommentsCount(ctx context.Context, in *ViewUnreadCommentsCountRequest, opts ...grpc.CallOption) (*ViewUnreadCommentsCountResponse, error)
		// 查看收藏的帖子
		LookCollectPost(ctx context.Context, in *LookCollectPostRequest, opts ...grpc.CallOption) (*LookCollectPostResponse, error)
		// 查看回复
		LookReplyComment(ctx context.Context, in *LookReplyCommentRequest, opts ...grpc.CallOption) (*LookReplyCommentResponse, error)
		// 回复评论
		ReplyComment(ctx context.Context, in *ReplyCommunityRequest, opts ...grpc.CallOption) (*ReplyCommunityResponse, error)
		// 对评论进行点赞
		LikeComment(ctx context.Context, in *LikeCommentRequest, opts ...grpc.CallOption) (*LikeCommentResponse, error)
		// 对评论点赞的取消
		CancelLikeComment(ctx context.Context, in *CancelLikeCommentRequest, opts ...grpc.CallOption) (*CancelLikeCommentResponse, error)
		// 定义点赞服务
		LikePost(ctx context.Context, in *CommunityLikePostRequest, opts ...grpc.CallOption) (*CommunityLikePostResponse, error)
		CancelLikePost(ctx context.Context, in *CommunityCancelLikePostRequest, opts ...grpc.CallOption) (*CommunityCancelLikePostResponse, error)
		CollectPost(ctx context.Context, in *CommunityCollectPostRequest, opts ...grpc.CallOption) (*CommunityCollectPostResponse, error)
		CancelCollectPost(ctx context.Context, in *CommunityCancelCollectPostRequest, opts ...grpc.CallOption) (*CommunityCancelCollectPostResponse, error)
		// 定义评论服务
		CommentPost(ctx context.Context, in *CommunityContentPostRequest, opts ...grpc.CallOption) (*CommunityContentPostResponse, error)
		CancelCommentPost(ctx context.Context, in *CommunityCancelContentPostRequest, opts ...grpc.CallOption) (*CommunityCancelContentPostResponse, error)
		// 定义帖子服务
		CommunityCreatePost(ctx context.Context, in *CommunityCreatePostRequest, opts ...grpc.CallOption) (*CommunityCreatePostResponse, error)
		CommunityDelPost(ctx context.Context, in *CommunityDelPostRequest, opts ...grpc.CallOption) (*CommunityDelPostResponse, error)
		CommunityLookPostByOwn(ctx context.Context, in *CommunityLookPostByOwnRequest, opts ...grpc.CallOption) (*CommunityLookPostByOwnResponses, error)
		CommunityLookAllPosts(ctx context.Context, in *CommunityLookAllPostsRequest, opts ...grpc.CallOption) (*CommunityLookAllPostsResponse, error)
		// 查看帖子的评论
		LookComment(ctx context.Context, in *LookCommentRequest, opts ...grpc.CallOption) (*LookCommentResponse, error)
		// 用户是否点赞帖子
		WhetherLikePost(ctx context.Context, in *WhetherLikePostRequest, opts ...grpc.CallOption) (*WhetherLikePostResponse, error)
		// 用户是否收藏帖子
		WhetherCollectPost(ctx context.Context, in *WhetherCollectPostRequest, opts ...grpc.CallOption) (*WhetherCollectPostResponse, error)
		// 查看帖子详情
		ViewPostDetails(ctx context.Context, in *ViewPostDetailsRequest, opts ...grpc.CallOption) (*ViewPostDetailsResponse, error)
		// 查看未读的评论
		ViewUnreadComments(ctx context.Context, in *ViewUnreadCommentsRequest, opts ...grpc.CallOption) (*ViewUnreadCommentsResponse, error)
	}

	defaultCommunity struct {
		cli zrpc.Client
	}
)

func NewCommunity(cli zrpc.Client) Community {
	return &defaultCommunity{
		cli: cli,
	}
}

// 查询最新的帖子
func (m *defaultCommunity) ViewTheLatestPost(ctx context.Context, in *ViewTheLatestPostRequest, opts ...grpc.CallOption) (*ViewTheLatestPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.ViewTheLatestPost(ctx, in, opts...)
}

// 查看用户有多少未读的信息
func (m *defaultCommunity) ViewUnreadCommentsCount(ctx context.Context, in *ViewUnreadCommentsCountRequest, opts ...grpc.CallOption) (*ViewUnreadCommentsCountResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.ViewUnreadCommentsCount(ctx, in, opts...)
}

// 查看收藏的帖子
func (m *defaultCommunity) LookCollectPost(ctx context.Context, in *LookCollectPostRequest, opts ...grpc.CallOption) (*LookCollectPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.LookCollectPost(ctx, in, opts...)
}

// 查看回复
func (m *defaultCommunity) LookReplyComment(ctx context.Context, in *LookReplyCommentRequest, opts ...grpc.CallOption) (*LookReplyCommentResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.LookReplyComment(ctx, in, opts...)
}

// 回复评论
func (m *defaultCommunity) ReplyComment(ctx context.Context, in *ReplyCommunityRequest, opts ...grpc.CallOption) (*ReplyCommunityResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.ReplyComment(ctx, in, opts...)
}

// 对评论进行点赞
func (m *defaultCommunity) LikeComment(ctx context.Context, in *LikeCommentRequest, opts ...grpc.CallOption) (*LikeCommentResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.LikeComment(ctx, in, opts...)
}

// 对评论点赞的取消
func (m *defaultCommunity) CancelLikeComment(ctx context.Context, in *CancelLikeCommentRequest, opts ...grpc.CallOption) (*CancelLikeCommentResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CancelLikeComment(ctx, in, opts...)
}

// 定义点赞服务
func (m *defaultCommunity) LikePost(ctx context.Context, in *CommunityLikePostRequest, opts ...grpc.CallOption) (*CommunityLikePostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.LikePost(ctx, in, opts...)
}

func (m *defaultCommunity) CancelLikePost(ctx context.Context, in *CommunityCancelLikePostRequest, opts ...grpc.CallOption) (*CommunityCancelLikePostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CancelLikePost(ctx, in, opts...)
}

func (m *defaultCommunity) CollectPost(ctx context.Context, in *CommunityCollectPostRequest, opts ...grpc.CallOption) (*CommunityCollectPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CollectPost(ctx, in, opts...)
}

func (m *defaultCommunity) CancelCollectPost(ctx context.Context, in *CommunityCancelCollectPostRequest, opts ...grpc.CallOption) (*CommunityCancelCollectPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CancelCollectPost(ctx, in, opts...)
}

// 定义评论服务
func (m *defaultCommunity) CommentPost(ctx context.Context, in *CommunityContentPostRequest, opts ...grpc.CallOption) (*CommunityContentPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CommentPost(ctx, in, opts...)
}

func (m *defaultCommunity) CancelCommentPost(ctx context.Context, in *CommunityCancelContentPostRequest, opts ...grpc.CallOption) (*CommunityCancelContentPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CancelCommentPost(ctx, in, opts...)
}

// 定义帖子服务
func (m *defaultCommunity) CommunityCreatePost(ctx context.Context, in *CommunityCreatePostRequest, opts ...grpc.CallOption) (*CommunityCreatePostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CommunityCreatePost(ctx, in, opts...)
}

func (m *defaultCommunity) CommunityDelPost(ctx context.Context, in *CommunityDelPostRequest, opts ...grpc.CallOption) (*CommunityDelPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CommunityDelPost(ctx, in, opts...)
}

func (m *defaultCommunity) CommunityLookPostByOwn(ctx context.Context, in *CommunityLookPostByOwnRequest, opts ...grpc.CallOption) (*CommunityLookPostByOwnResponses, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CommunityLookPostByOwn(ctx, in, opts...)
}

func (m *defaultCommunity) CommunityLookAllPosts(ctx context.Context, in *CommunityLookAllPostsRequest, opts ...grpc.CallOption) (*CommunityLookAllPostsResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.CommunityLookAllPosts(ctx, in, opts...)
}

// 查看帖子的评论
func (m *defaultCommunity) LookComment(ctx context.Context, in *LookCommentRequest, opts ...grpc.CallOption) (*LookCommentResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.LookComment(ctx, in, opts...)
}

// 用户是否点赞帖子
func (m *defaultCommunity) WhetherLikePost(ctx context.Context, in *WhetherLikePostRequest, opts ...grpc.CallOption) (*WhetherLikePostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.WhetherLikePost(ctx, in, opts...)
}

// 用户是否收藏帖子
func (m *defaultCommunity) WhetherCollectPost(ctx context.Context, in *WhetherCollectPostRequest, opts ...grpc.CallOption) (*WhetherCollectPostResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.WhetherCollectPost(ctx, in, opts...)
}

// 查看帖子详情
func (m *defaultCommunity) ViewPostDetails(ctx context.Context, in *ViewPostDetailsRequest, opts ...grpc.CallOption) (*ViewPostDetailsResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.ViewPostDetails(ctx, in, opts...)
}

// 查看未读的评论
func (m *defaultCommunity) ViewUnreadComments(ctx context.Context, in *ViewUnreadCommentsRequest, opts ...grpc.CallOption) (*ViewUnreadCommentsResponse, error) {
	client := community.NewCommunityClient(m.cli.Conn())
	return client.ViewUnreadComments(ctx, in, opts...)
}
