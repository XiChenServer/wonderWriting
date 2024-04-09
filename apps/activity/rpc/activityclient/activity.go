// Code generated by goctl. DO NOT EDIT.
// Source: activity.proto

package activityclient

import (
	"context"

	"calligraphy/apps/activity/rpc/types/activity"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GrabPointsRequest  = activity.GrabPointsRequest
	GrabPointsResponse = activity.GrabPointsResponse

	Activity interface {
		GrabPoints(ctx context.Context, in *GrabPointsRequest, opts ...grpc.CallOption) (*GrabPointsResponse, error)
	}

	defaultActivity struct {
		cli zrpc.Client
	}
)

func NewActivity(cli zrpc.Client) Activity {
	return &defaultActivity{
		cli: cli,
	}
}

func (m *defaultActivity) GrabPoints(ctx context.Context, in *GrabPointsRequest, opts ...grpc.CallOption) (*GrabPointsResponse, error) {
	client := activity.NewActivityClient(m.cli.Conn())
	return client.GrabPoints(ctx, in, opts...)
}
