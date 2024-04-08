package logic

import (
	groupModel "calligraphy/apps/group/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"calligraphy/apps/group/rpc/internal/svc"
	"calligraphy/apps/group/rpc/types/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookRecordByUserIdLogic {
	return &LookRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LookRecordByUserId 查看某人的书法记录（带缓存和分页）
func (l *LookRecordByUserIdLogic) LookRecordByUserId(in *group.LookRecordByUserIdRequest) (*group.LookRecordByUserIdResponse, error) {
	// 构建缓存键
	cacheKey := fmt.Sprintf("user_records:%d:%d:%d", in.UserId, in.Page, in.PageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheForRecords(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	recordInfo, totalCount, err := getRecordsByUser(l.svcCtx.DB, in.UserId, int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := totalCount / int64(in.PageSize)
	if totalCount%int64(in.PageSize) != 0 {
		totalPages++
	}

	// 构建并返回记录信息响应
	resp := &group.LookRecordByUserIdResponse{
		RecordInfo:  recordInfo,
		CurrentPage: in.Page,
		PageSize:    in.PageSize,
		Offset:      uint32((in.Page - 1) * in.PageSize),
		Overflow:    in.Page > uint32(totalPages),
		TotalPages:  uint32(totalPages),
		TotalCount:  uint64(totalCount),
	}

	cacheTime := 60 * 5 // 缓存时间为5分钟
	// 将查询结果存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

func getRecordsByUser(db *gorm.DB, userId uint32, page, pageSize int) ([]*group.RecordSimpleInfo, int64, error) {
	var res []groupModel.RecordContent

	if pageSize <= 0 {
		return nil, 0, errors.New("pageSize should be a positive integer")
	}

	offset := (page - 1) * pageSize
	err := db.Where("user_id = ?", userId).Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取总记录数
	var totalCount int64
	err = db.Model(&groupModel.RecordContent{}).Where("user_id = ?", userId).Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	if totalCount == 0 {
		return []*group.RecordSimpleInfo{}, 0, nil
	}

	// 计算总页数
	totalPage := uint32(totalCount) / uint32(pageSize)
	if uint32(totalCount)%uint32(pageSize) != 0 {
		totalPage++
	}

	// 其他分页和数据获取逻辑...

	// 构建用于返回的记录信息切片
	var recordInfo []*group.RecordSimpleInfo

	// 遍历查询到的记录信息
	for _, v := range res {
		newRecord := &group.RecordSimpleInfo{
			RecordId:   uint32(v.ID),
			UserId:     uint32(v.UserID),
			Content:    v.Content,
			Image:      v.Image,
			Score:      float32(v.Score),
			CreateTime: int32(v.CreatedAt.Unix()),
		}

		recordInfo = append(recordInfo, newRecord)
	}

	return recordInfo, totalCount, nil
}

// 从缓存中获取数据
func getFromCacheForRecords(l *LookRecordByUserIdLogic, key string) (*group.LookRecordByUserIdResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data group.LookRecordByUserIdResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func toJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
