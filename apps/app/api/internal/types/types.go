// Code generated by goctl. DO NOT EDIT.
package types

type PostInfo struct {
	UserInfo     UserSimpleInfo `json:"user_info"`
	Id           uint           `json:"id"`
	UserId       uint           `json:"user_id"`
	ContentCount uint           `json:"content_count"`
	LikeCount    uint           `json:"like_count"`
	CollectCount uint           `json:"collect_count"`
	Content      string         `json:"content"`
	ImageUrls    []string       `json:"image_urls"`
	CreateTime   int32          `json:"create_time"`
	DeleteTime   int32          `json:"delete_time"`
}

type UserSimpleInfo struct {
	Id          uint   `json:"user_id"`
	NickName    string `json:"nick_name"`
	Account     string `json:"account"`
	AvatarImage string `json:"avatar_image"`
}

type PostCreateRequest struct {
}

type PostCreateResponse struct {
	PostId uint `json:"post_id"`
}

type PostDelRequest struct {
	PostId uint `json:"post_id"`
}

type PostDelResponse struct {
}

type LookPostByOwnRequest struct {
	UserId   uint32 `json:"user_id"`
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}

type LookPostByOwnResponses struct {
	PostData    []*PostInfo `json:"post_data"`
	CurrentPage uint32      `json:"current_page"`
	PageSize    uint32      `json:"page_size"`
	Offset      uint32      `json:"offset"`
	Overflow    bool        `json:"overflow"`
	TotalPage   uint32      `json:"total_page"`
	TotalCount  uint64      `json:"total_count"`
}

type LookAllPostsRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}

type LookAllPostsResponse struct {
	PostData    []*PostInfo `json:"post_data"`
	CurrentPage uint32      `json:"current_page"`
	PageSize    uint32      `json:"page_size"`
	Offset      uint32      `json:"offset"`
	Overflow    bool        `json:"overflow"`
	TotalPage   uint32      `json:"total_page"`
	TotalCount  uint64      `json:"total_count"`
}

type LikePostRequest struct {
	PostId uint `json:"post_id"`
}

type LikePostResponse struct {
	LikeId uint `json:"like_id"`
}

type CancelLikePostRequest struct {
	LikeId uint `json:"like_id"`
	PostId uint `json:"post_id"`
}

type CancelLikePostResponse struct {
}

type CommentPostRequest struct {
	PostId  uint   `json:"post_id"`
	Content string `json:"content"`
}

type CommentPostResponse struct {
	CommentId uint `json:"comment_id"`
}

type CancelCommentPostRequest struct {
	PostId    uint `json:"post_id"`
	CommentId uint `json:"comment_id"`
}

type CancelCommentPostResponse struct {
}

type CollectPostRequest struct {
	PostId uint `json:"post_id"`
}

type CollectPostResponse struct {
	CollectId uint `json:"collect_id"`
}

type CancelCollectPostRequest struct {
	PostId uint `json:"post_id"`
}

type CancelCollectPostResponse struct {
}

type CommentInfo struct {
	Id         uint           `json:"id"`
	CreateTime int32          `json:"create_time"`
	PostId     uint           `json:"post_id"`
	Comment    string         `json:"comment"`
	UserInfo   UserSimpleInfo `json:"user_info"`
}

type LookCommentRequest struct {
	PostId   uint   `json:"post_id"`
	Page     uint   `json:"page"`
	PageSize uint32 `json:"page_size"`
}

type LookCommentResponse struct {
	CommentData []*CommentInfo `json:"comment_data"`
	CurrentPage uint32         `json:"current_page"`
	PageSize    uint32         `json:"page_size"`
	Offset      uint32         `json:"offset"`
	Overflow    bool           `json:"overflow"`
	TotalPage   uint32         `json:"total_page"`
	TotalCount  uint64         `json:"total_count"`
}

type WhetherLikePostRequest struct {
	OtherId uint32 `json:"other_id"`
}

type WhetherLikePostResponse struct {
}

type WhetherCollectPostRequest struct {
	OtherId uint32 `json:"other_id"`
}

type WhetherCollectPostResponse struct {
}

type StatusWithPost struct {
	WhetherLike    bool `json:"whether_like"`
	WhetherCollect bool `json:"whether_collect"`
	WhetherFollow  bool `json:"whether_follow"`
}

type ViewPostDetailsRequest struct {
	PostId uint32 `json:"post_id"`
}

type ViewPostDetailsResponse struct {
	PostData   PostInfo       `json:"post_data"`
	StatusData StatusWithPost `json:"status_data"`
}

type VerificationRequest struct {
	Email string `json:"email"`
}

type VerificationResponse struct {
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type JwtInfo struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type UserRegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	EmailCode string `json:"email_code"`
}

type UserRegisterResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UserInfoResponse struct {
	Id               int64  `json:"id"`
	NickName         string `json:"nick_name"`
	Account          string `json:"account"`
	Email            string `json:"email"`
	AvatarBackground string `json:"avatar_background"`
	BackgroundImage  string `json:"background_image"`
	Phone            string `json:"phone"`
	PostCount        int64  `json:"post_count"`
	FollowCount      int64  `json:"follow_count"`
	FansCount        int64  `json:"fans_count"`
	LikeCount        int64  `json:"like_count"`
	PointCount       int64  `json:"point_count"`
}

type UserForgetPwdRequest struct {
	Email     string `json:"email"`
	EmailCode string `json:"email_code"`
}

type UserForgetPwdResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type UserModPwdRequset struct {
	Password string `json:"password"`
}

type UserModPwdResponse struct {
}

type UserModAvatarResponse struct {
}

type UserModBackgroundResponse struct {
}

type UserModInfoRequest struct {
	NickName string `json:"nick_name, optional"`
	Phone    string `json:"phone, optional"`
}

type UserModInfoResponse struct {
}

type UserFollowRequest struct {
	UserId uint32 `json:"user_id"`
}

type UserFollowResponse struct {
}

type UserCancelFollowRequest struct {
	UserId uint32 `json:"user_id"`
}

type UserCancelFollowResponse struct {
}

type UserExhibitInfo struct {
	UserId           uint32 `json:"user_id"`
	AvatarBackground string `json:"avatar_background"`
	NickName         string `json:"nick_name"`
	FollowCount      int64  `json:"follow_count"`
	FansCount        int64  `json:"fans_count"`
	Email            string `json:"email"`
}

type LookAllFollowRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size, optional"`
}

type LookAllFollowResponse struct {
	UserData    []*UserExhibitInfo `json:"post_data"`
	CurrentPage uint32             `json:"current_page"`
	PageSize    uint32             `json:"page_size"`
	Offset      uint32             `json:"offset"`
	Overflow    bool               `json:"overflow"`
	TotalPage   uint32             `json:"total_page"`
	TotalCount  uint64             `json:"total_count"`
}

type LookAllFansRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size, optional"`
}

type LookAllFansResponse struct {
	UserData    []*UserExhibitInfo `json:"post_data"`
	CurrentPage uint32             `json:"current_page"`
	PageSize    uint32             `json:"page_size"`
	Offset      uint32             `json:"offset"`
	Overflow    bool               `json:"overflow"`
	TotalPage   uint32             `json:"total_page"`
	TotalCount  uint64             `json:"total_count"`
}

type WhetherFollowUserRequest struct {
	OtherId uint32 `json:"other_id"`
}

type WhetherFollowUserResponse struct {
}

type UserPopularInfo struct {
	UserId    uint32 `json:"user_id"`
	NickName  string `json:"nick_name"`
	Account   string `json:"account"`
	LikeCount int64  `json:"like_count"`
	Avatar    string `json:"avatar"`
}

type UserPopularityRankingsRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}

type UserPopularityRankingsResponse struct {
	UserPopularData []*UserPopularInfo `json:"user_popular_data"`
	CurrentPage     uint32             `json:"current_page"`
	PageSize        uint32             `json:"page_size"`
	Offset          uint32             `json:"offset"`
	Overflow        bool               `json:"overflow"`
	TotalPage       uint32             `json:"total_page"`
	TotalCount      uint64             `json:"total_count"`
}

type PostPopularityInfo struct {
	PostId          uint32           `json:"post_id"`
	Content         string           `json:"content"`
	LikeCont        int64            `json:"like_count"`
	CollectionCount int64            `json:"collection_count"`
	CommentCount    int64            `json:"comment_count"`
	PopularInfo     *UserPopularInfo `json:"popular_info"`
	PostImage       []string         `json:"post_image"`
}

type PostPopularityRankingsRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}

type PostPopularityRankingsResponse struct {
	PostPopularData []*PostPopularityInfo `json:"post_popular_data"`
	CurrentPage     uint32                `json:"current_page"`
	PageSize        uint32                `json:"page_size"`
	Offset          uint32                `json:"offset"`
	Overflow        bool                  `json:"overflow"`
	TotalPage       uint32                `json:"total_page"`
	TotalCount      uint64                `json:"total_count"`
}

type StartCheckRequest struct {
}

type StartCheckResponse struct {
	CheckId         uint32 `json:"check_id"`
	UserId          uint32 `json:"user_id"`
	ContinuousDays  int32  `json:"continuous_days"`
	CreateTime      int32  `json:"create_time"`
	LastCheckInTime int32  `json:"last_check_in_time"`
}

type RecordSimpleInfo struct {
	RecordId   uint32  `json:"record_id"`
	UserId     uint32  `json:"user_id"`
	Content    string  `json:"content"`
	Image      string  `json:"image"`
	Score      float32 `json:"score"`
	CreateTime int32   `json:"create_time"`
}

type CreateRecordRequest struct {
	UserId  uint32  `json:"user_id"`
	Content string  `json:"content"`
	Image   string  `json:"image"`
	Score   float32 `json:"score"`
}

type CreateRecordResponse struct {
	RecordInfo RecordSimpleInfo `json:"record_info"`
}

type LookRecordByUserIdRequest struct {
	UserId   uint32 `json:"user_id"`
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}

type LookRecordByUserIdResponse struct {
	RecordInfo  []*RecordSimpleInfo `json:"record_info"`
	CurrentPage uint32              `json:"current_page"`
	PageSize    uint32              `json:"page_size"`
	Offset      uint32              `json:"offset"`
	Overflow    bool                `json:"overflow"`
	TotalPage   uint32              `json:"total_page"`
	TotalCount  uint64              `json:"total_count"`
}

type CheckPunchCardModelRequest struct {
}

type CheckPunchCardModelResponse struct {
	Data bool `json:"data"`
}

type CheckInResponse struct {
}

type GrabPointsRequest struct {
}

type GrabPointsResponse struct {
}

type ActivityInfo struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name"`
	Info        string `json:"info"`
	Location    string `json:"location"`
	DateTime    string `json:"date_time"`
	Organizer   string `json:"organizer"`
	EndDateTime string `json:"end_date_time"`
	Duration    string `json:"duration"`
	RewardsInfo string `json:"rewards_info"`
}

type LookAllActivitiesRequest struct {
	Page      uint32 `json:"page"`
	Page_size uint32 `json:"page_size"`
}

type LookAllActivitiesResponse struct {
	Activities   []*ActivityInfo `json:"activities"`
	Current_page uint32          `json:"current_page"`
	Page_size    uint32          `json:"page_size"`
	Offset       uint32          `json:"offset"`
	Overflow     bool            `json:"overflow"`
	Total_pages  uint32          `json:"total_pages"`
	Total_count  uint64          `json:"total_count"`
}

type UserSignUpActivityRequest struct {
	User_id     uint32 `json:"user_id"`
	Activity_id uint32 `json:"activity_id"`
}

type UserSignUpActivityResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UserViewAllActivitiesRequest struct {
	User_id   uint32 `json:"user_id"`
	Page      uint32 `json:"page"`
	Page_size uint32 `json:"page_size"`
}

type UserViewAllActivitiesResponse struct {
	Activities   []*ActivityInfo `json:"activities"`
	Current_page uint32          `json:"current_page"`
	Page_size    uint32          `json:"page_size"`
	Offset       uint32          `json:"offset"`
	Overflow     bool            `json:"overflow"`
	Total_pages  uint32          `json:"total_pages"`
	Total_count  uint64          `json:"total_count"`
}
