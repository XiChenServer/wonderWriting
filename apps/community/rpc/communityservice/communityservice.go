// Code generated by goctl. DO NOT EDIT.
// Source: community.proto

package communityservice

import (
	"context"

	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
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
	PostInfo                           = community.PostInfo

	CommunityService interface {
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
	}

	defaultCommunityService struct {
		cli zrpc.Client
	}
)

func NewCommunityService(cli zrpc.Client) CommunityService {
	return &defaultCommunityService{
		cli: cli,
	}
}

// 定义点赞服务
func (m *defaultCommunityService) LikePost(ctx context.Context, in *CommunityLikePostRequest, opts ...grpc.CallOption) (*CommunityLikePostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.LikePost(ctx, in, opts...)
}

func (m *defaultCommunityService) CancelLikePost(ctx context.Context, in *CommunityCancelLikePostRequest, opts ...grpc.CallOption) (*CommunityCancelLikePostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CancelLikePost(ctx, in, opts...)
}

func (m *defaultCommunityService) CollectPost(ctx context.Context, in *CommunityCollectPostRequest, opts ...grpc.CallOption) (*CommunityCollectPostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CollectPost(ctx, in, opts...)
}

func (m *defaultCommunityService) CancelCollectPost(ctx context.Context, in *CommunityCancelCollectPostRequest, opts ...grpc.CallOption) (*CommunityCancelCollectPostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CancelCollectPost(ctx, in, opts...)
}

// 定义评论服务
func (m *defaultCommunityService) CommentPost(ctx context.Context, in *CommunityContentPostRequest, opts ...grpc.CallOption) (*CommunityContentPostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CommentPost(ctx, in, opts...)
}

func (m *defaultCommunityService) CancelCommentPost(ctx context.Context, in *CommunityCancelContentPostRequest, opts ...grpc.CallOption) (*CommunityCancelContentPostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CancelCommentPost(ctx, in, opts...)
}

// 定义帖子服务
func (m *defaultCommunityService) CommunityCreatePost(ctx context.Context, in *CommunityCreatePostRequest, opts ...grpc.CallOption) (*CommunityCreatePostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CommunityCreatePost(ctx, in, opts...)
}

func (m *defaultCommunityService) CommunityDelPost(ctx context.Context, in *CommunityDelPostRequest, opts ...grpc.CallOption) (*CommunityDelPostResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CommunityDelPost(ctx, in, opts...)
}

func (m *defaultCommunityService) CommunityLookPostByOwn(ctx context.Context, in *CommunityLookPostByOwnRequest, opts ...grpc.CallOption) (*CommunityLookPostByOwnResponses, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CommunityLookPostByOwn(ctx, in, opts...)
}

func (m *defaultCommunityService) CommunityLookAllPosts(ctx context.Context, in *CommunityLookAllPostsRequest, opts ...grpc.CallOption) (*CommunityLookAllPostsResponse, error) {
	client := community.NewCommunityServiceClient(m.cli.Conn())
	return client.CommunityLookAllPosts(ctx, in, opts...)
}
