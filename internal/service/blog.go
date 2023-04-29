package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/cache"
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"net/http"
	"os"
	"strings"
	"time"
)

//type IBlogService interface {
//	FindBlogById(ctx *gin.Context, req *dto.FindBlogByIdReq) (*dto.FindBlogByIdRsp, error)                 // 根据id获取博客
//	ListHotBlogs(ctx *gin.Context, req *dto.BlogHotReq) ([]*dto.BlogHotRsp, error)                         // 获取热门博客
//	ListBlogByUserId(ctx *gin.Context, req *dto.ListBlogsByUserIdReq) ([]*dto.ListBlogsByUserIdRsp, error) // 根据用户id获取博客
//	UploadBlogImg(c *gin.Context, req *dto.UploadBlogImgReq)                                               // 上传博客图片
//	CreateBlog(c *gin.Context, req *dto.CreateBlogReq) (*dto.CreateBlogRsp, error)                         // 创建博客
//	IsBlogLiked(c *gin.Context, req *dto.IsBlogLikeReq) (*dto.IsBlogLikeRsp, error)                        // 判断博客是否被点赞
//	LikeBlogByUserId(c *gin.Context, req *dto.LikeBlogReq) (*dto.LikeBlogRsp, error)                       // 点赞&取消点赞
//	ListBlogLikes(c *gin.Context, req *dto.ListBlogLikesReq) (*dto.ListLikedUsersByBlogIdRsp, error)       // 获取博客点赞列表
//}

type BlogService struct {
	DB  *gorm.DB
	Req *assembler.BlogReq
	Rsp *assembler.BlogRsp
}

func NewBlogService(db *gorm.DB) *BlogService {
	return &BlogService{
		DB:  db,
		Req: &assembler.BlogReq{},
		Rsp: &assembler.BlogRsp{},
	}
}

func (svc *BlogService) FindBlogById(ctx *gin.Context, req *dto.FindBlogByIdReq) (*dto.FindBlogByIdRsp, error) {
	var blog model.Blog
	err := svc.DB.Preload("User").First(&blog, req.Id).Error
	if err != nil {
		return nil, err
	}
	user, ok := ctx.Get("user")
	if !ok {
		return svc.Rsp.E2DFindBlogById(&blog, nil), err
	}
	return svc.Rsp.E2DFindBlogById(&blog, user.(*model.User)), err
}

// ListHotBlogs 获取热门博客
func (svc *BlogService) ListHotBlogs(ctx *gin.Context, req *dto.BlogHotReq) ([]*dto.BlogHotRsp, error) {
	var blogs []*model.Blog
	page, pageSize := req.Current-1, 5
	err := svc.DB.Preload("User").Offset(page * pageSize).Limit(pageSize).Find(&blogs).Error
	if u, ok := ctx.Get("user"); ok {
		return svc.Rsp.E2DHot(blogs, u.(*model.User)), err
	} else {
		return svc.Rsp.E2DHot(blogs, nil), err
	}
}

func (svc *BlogService) ListBlogByUserId(ctx *gin.Context, req *dto.ListBlogsByUserIdReq) ([]*dto.ListBlogsByUserIdRsp, error) {
	var blogs []*model.Blog
	//var page, pageSize int
	err := svc.DB.Where("user_id = ?", req.UserId).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return svc.Rsp.E2DListBlogsByUserId(blogs), nil
}

func (svc *BlogService) UploadBlogImg(c *gin.Context, req *dto.UploadBlogImgReq) (*dto.UploadBlogImgRsp, error) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("获取上传文件失败：%s", err.Error()))
		return nil, err
	}
	url := fmt.Sprintf("imgs/blogs/%v/%v", req.UserId, file.Filename)
	folder := viper.Get("img_folder")
	path := fmt.Sprintf("%v/%v", folder, url)
	// 将文件保存到服务器上
	//fmt.Println(path)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("保存上传文件失败：%s", err.Error()))
		return nil, err
	}
	rsp := &dto.UploadBlogImgRsp{Url: url}
	return rsp, nil
}

// IsBlogLiked 判断博客是否被点赞
func (svc *BlogService) IsBlogLiked(c *gin.Context, req *dto.IsBlogLikeReq) (*dto.IsBlogLikeRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *BlogService) LikeBlogByUserId(c *gin.Context, req *dto.LikeBlogByUserIdReq) (*dto.LikeBlogByUserIdRsp, error) {
	key := model.BlogLikeKey(req.BlogId)
	// 判断当前登录用户是否已经点赞
	// 从redis的set集合中获取当前用户是否已经点赞
	_, err := cache.RedisStore.ZScore(c, key, req.UserId).Result()
	now := time.Now().Unix()
	if err == redis.Nil { // 未点赞
		// 如果返回的错误是 redis.Nil，说明指定的字段不存在于 sorted set 中
		svc.DB.Model(&model.Blog{}).Where("id = ?", req.BlogId).Update("liked", gorm.Expr("liked + ?", 1))
		cache.RedisStore.ZAdd(c, key, redis.Z{
			Score:  float64(now),
			Member: req.UserId,
		})
		return &dto.LikeBlogByUserIdRsp{}, nil
	} else if err != nil {
		// 如果返回的错误不是 redis.Nil，说明发生了其他错误
		return nil, err
	} else { // 如果没有发生错误，result 的值即为指定字段的分数
		cache.RedisStore.ZRem(c, key, req.UserId)
		svc.DB.Model(&model.Blog{}).Where("id = ?", req.BlogId).Update("liked", gorm.Expr("liked - ?", 1))
		return &dto.LikeBlogByUserIdRsp{}, nil
	}
}

// ListLikedUsersByBlogId  获取博客点赞列表
func (svc *BlogService) ListLikedUsersByBlogId(c *gin.Context, req *dto.ListLikedUsersByBlogIdReq) ([]*dto.ListLikedUsersByBlogIdRsp, error) {
	key := model.BlogLikeKey(req.BlogId)
	result, err := cache.RedisStore.ZRangeWithScores(c, key, 0, 9).Result()
	//fmt.Println(len(result))
	if err == redis.Nil { // 无点赞记录
		return nil, err
	}
	userIdxs := make([]string, 0)
	for _, z := range result {
		member := z.Member.(string) // 成员名称
		userIdxs = append(userIdxs, member)
	}

	order := fmt.Sprintf("FIELD(id,%s)", strings.Join(userIdxs, ","))
	var users []*model.User
	err = svc.DB.Select("id,nick_name,icon").Where("id IN ?", userIdxs).Order(order).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return svc.Rsp.E2DListBlogLikes(users), nil
}

func (svc *BlogService) DeleteBlogImg(ctx *gin.Context, req *dto.DeleteBlogImgReq) (*dto.DeleteBlogImgRsp, error) {
	folder := viper.Get("img_folder")
	path := fmt.Sprintf("%v/%v", folder, req.Url)
	err := os.Remove(path)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteBlogImgRsp{}, nil
}

func (svc *BlogService) CreateBlog(ctx *gin.Context, req *dto.CreateBlogReq) (*dto.CreateBlogRsp, error) {
	// TODO 将博客信息推送给粉丝
	blog := &model.Blog{
		ShopId:  req.ShopId,
		UserId:  req.UserId,
		Title:   req.Title,
		Images:  req.Images,
		Content: req.Content,
	}
	err := svc.DB.Create(&blog).Error
	if err != nil {
		return nil, err
	}
	// 获取粉丝列表
	//fans, err := svc.DB.ListFansByUserId(ctx, &dto.ListFansByUserIdReq{UserId: req.UserId})
	fans := make([]*model.Follow, 0)
	svc.DB.Select("user_id").Where("follow_id = ?", req.UserId).Find(&fans)
	for _, fan := range fans {
		key := model.SubscriptionBlogsKey(fan.UserId)
		cache.RedisStore.ZAdd(ctx, key, redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: blog.ID,
		})
	}
	return &dto.CreateBlogRsp{}, nil
}

func (svc *BlogService) ListBlogsBySubscription(ctx *gin.Context, req *dto.ListBlogsBySubscriptionReq) (*dto.ListBlogsBySubscriptionRsp, error) {
	// 从redis中拿到订阅的博客列表
	key := model.SubscriptionBlogsKey(req.UserId)
	results, err := cache.RedisStore.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    req.LastTime,
		Offset: req.Offset,
		Count:  10,
	}).Result()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	var blogIds []string
	for _, z := range results {
		member := z.Member.(string) // 成员名称
		blogIds = append(blogIds, member)
	}
	lastTime := results[len(results)-1].Score
	offset := 0
	for i := len(results) - 1; i >= 0; i-- {
		if results[i].Score == lastTime {
			offset++
		} else {
			break
		}
	}
	var blogs []*model.Blog

	order := fmt.Sprintf("FIELD(id,%s)", strings.Join(blogIds, ","))
	err = svc.DB.Where("id IN ?", blogIds).Order(order).Find(&blogs).Error
	return svc.Rsp.E2DListBlogsBySubscription(blogs, lastTime, 1), nil
}
