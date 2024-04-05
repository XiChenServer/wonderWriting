package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsercretePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsercretePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsercretePostLogic {
	return &UsercretePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsercretePostLogic) UsercretePost(r *http.Request) (resp *types.PostCreateResponse, err error) {
	// todo: add your logic here and delete this line
	content := r.FormValue("content")
	if content == "" {
		err = errors.New("Field 'content' is required")
		return &types.PostCreateResponse{}, err
	}
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	var urls []string
	// 解析表单，获取多个文件字段
	if err := r.ParseMultipartForm(32 << 20); err != nil { // 设置最大内存为 32MB
		fmt.Println(err)
		return &types.PostCreateResponse{}, err
	}
	// 获取表单中的文件字段
	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		return &types.PostCreateResponse{}, errors.New("no images uploaded")
	}

	// 依次处理每个文件
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println(err)
			return &types.PostCreateResponse{}, err
		}
		defer file.Close()

		// 上传到七牛云
		url, err := qiniu.UploadToQiNiu(file, fileHeader.Size, "AvatarBackground/", fileHeader.Filename)
		if err != nil {
			fmt.Println(err.Error())
			return &types.PostCreateResponse{}, err
		}
		urls = append(urls, url)
		// 处理上传成功的情况，比如保存 URL 或其他操作
		fmt.Println("Uploaded file:", url)
	}
	//调用rpc的接口
	res, err := l.svcCtx.CommunityRpc.CommunityCreatePost(l.ctx, &community.CommunityCreatePostRequest{
		UserId:    uint32(uid),
		Content:   content,
		ImageUrls: urls,
	})
	if err != nil {
		return &types.PostCreateResponse{}, err
	}
	return &types.PostCreateResponse{PostId: uint(res.PostId)}, nil
}
