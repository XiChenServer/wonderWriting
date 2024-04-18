package grow

import (
	"calligraphy/apps/grow/rpc/types/grow"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckInLogic {
	return &CheckInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckInLogic) CheckIn(r *http.Request) (resp *types.CheckInResponse, err error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	content := r.FormValue("content")
	if content == "" {
		err = errors.New("Field 'content' is required")
		return &types.CheckInResponse{}, err
	}

	score := r.FormValue("score")
	if score == "" {
		err = errors.New("Field 'score' is required")
		return &types.CheckInResponse{}, err
	}
	scoreFloat, err := strconv.ParseFloat(score, 64)
	if err != nil {
		return nil, err
	}
	// 图片上传到七牛云
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	url, err := qiniu.UploadToQiNiu(file, handler.Size, "BackgroundImage/", handler.Filename)
	if err != nil {
		return nil, err
	}

	//调用rpc获取进行创建
	_, err = l.svcCtx.GrowRpc.CheckIn(l.ctx, &grow.CheckInRequest{
		UserId:  uint32(uid),
		Content: content,
		Image:   url,
		Score:   float32(scoreFloat),
	})
	if err != nil {
		return nil, err
	}
	////创建响应体
	//record := types.RecordSimpleInfo{
	//	RecordId:   res.RecordInfo.RecordId,
	//	UserId:     res.RecordInfo.UserId,
	//	Content:    res.RecordInfo.Content,
	//	Image:      res.RecordInfo.Image,
	//	Score:      res.RecordInfo.Score,
	//	CreateTime: res.RecordInfo.CreateTime,
	//}
	return &types.CheckInResponse{}, nil
}
