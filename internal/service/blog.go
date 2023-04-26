package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/cache"
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"net/http"
)

type IBlogService interface {
	FindBlogById(ctx *gin.Context, req *dto.FindBlogByIdReq) (*dto.FindBlogByIdRsp, error)              // 根据id获取博客
	ListHotBlogs(ctx *gin.Context, req *dto.BlogHotReq) ([]*dto.BlogHotRsp, error)                      // 获取热门博客
	ListBlogByUserId(ctx *gin.Context, req *dto.ListBlogsByUserIdReq) ([]*dto.BlogGetByUseIdRsp, error) // 根据用户id获取博客
	UploadBlogImg(c *gin.Context, req *dto.UploadBlogImgReq)                                            // 上传博客图片
	CreateBlog(c *gin.Context, req *dto.CreateBlogReq) (*dto.CreateBlogRsp, error)                      // 创建博客
	IsBlogLiked(c *gin.Context, req *dto.IsBlogLikeReq) (*dto.IsBlogLikeRsp, error)                     // 判断博客是否被点赞
	LikeBlogByUserId(c *gin.Context, req *dto.LikeBlogReq) (*dto.LikeBlogRsp, error)                    // 点赞&取消点赞
	ListBlogLikes(c *gin.Context, req *dto.ListBlogLikesReq) (*dto.ListLikedUsersByBlogIdRsp, error)    // 获取博客点赞列表
}

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
	user := ctx.MustGet("user").(*model.User)
	return svc.Rsp.E2DFindBlogById(&blog, user), err
}

// ListHotBlogs 获取热门博客
func (svc *BlogService) ListHotBlogs(ctx *gin.Context, req *dto.BlogHotReq) ([]*dto.BlogHotRsp, error) {
	var blogs []*model.Blog
	user := ctx.MustGet("user").(*model.User)
	//page, pageSize := req.Page, req.PageSize
	// TODO
	page, pageSize := req.Current, 5
	//err := svc.DB.Model(model.Blog{}).Preload("User").Offset(page * pageSize).Limit(pageSize).Find(&blogs).Error
	err := svc.DB.Preload("User").Offset(page * pageSize).Limit(pageSize).Find(&blogs).Error

	return svc.Rsp.E2DHot(blogs, user), err
}

func (svc *BlogService) ListBlogByUserId(ctx *gin.Context, req *dto.ListBlogsByUserIdReq) ([]*dto.BlogGetByUseIdRsp, error) {
	var blogs []*model.Blog
	//var page, pageSize int
	err := svc.DB.Where("user_id = ?", req.UserId).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return svc.Rsp.E2DListBlogsByUserId(blogs), nil
}

func (svc *BlogService) UploadBlogImg(c *gin.Context, req *dto.UploadBlogImgReq) {
	folder := viper.Get("img_folder")
	// 图像的保存位置

	// 实现博客图像的上传
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("获取上传文件失败：%s", err.Error()))
		return
	}
	path := fmt.Sprintf("%v/%v/%v", folder, req.UserId, file.Filename)
	// 将文件保存到服务器上
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("保存上传文件失败：%s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' 上传成功！", file.Filename))
}

// CreateBlog 创建博客
func (svc *BlogService) CreateBlog(c *gin.Context, req *dto.CreateBlogReq) (*dto.CreateBlogRsp, error) {
	// 获取登录用户
	// 获取博客信息
	// 返回博客id
	return nil, nil
}

// IsBlogLiked 判断博客是否被点赞
func (svc *BlogService) IsBlogLiked(c *gin.Context, req *dto.IsBlogLikeReq) (*dto.IsBlogLikeRsp, error) {
	//TODO implement me
	panic("implement me")
}

// LikeBlogByUserId 点赞&取消点赞
func (svc *BlogService) LikeBlogByUserId(c *gin.Context, req *dto.LikeBlogByUserIdReq) (*dto.LikeBlogByUserIdRsp, error) {
	key := fmt.Sprintf("like:%v", req.BlogId)
	// 判断当前登录用户是否已经点赞
	// 从redis的set集合中获取当前用户是否已经点赞
	eixst, _ := cache.RedisStore.SIsMember(c, key, req.UserId).Result()
	if !eixst { // 未点赞
		svc.DB.Model(&model.Blog{}).Where("id = ?", req.BlogId).Update("liked", gorm.Expr("liked + ?", 1))
		cache.RedisStore.SAdd(c, key, req.UserId)
		return &dto.LikeBlogByUserIdRsp{}, nil
	}

	// 4.已经点赞逻辑   => 数据库点赞数-1 把用户从redis的set集合中删除
	svc.DB.Model(&model.Blog{}).Where("id = ?", req.BlogId).Update("liked", gorm.Expr("liked - ?", 1))
	cache.RedisStore.SRem(c, key, req.UserId)
	return &dto.LikeBlogByUserIdRsp{}, nil
}

// ListLikedUsersByBlogId  获取博客点赞列表
func (svc *BlogService) ListLikedUsersByBlogId(c *gin.Context, req *dto.ListLikedUsersByBlogIdReq) ([]*dto.ListLikedUsersByBlogIdRsp, error) {
	userIdxs, err := cache.RedisStore.SMembers(c, fmt.Sprintf("like:%v", req.BlogId)).Result()
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for _, userIdx := range userIdxs {
		var user model.User
		svc.DB.First(&user, userIdx)
		users = append(users, &user)
	}
	return svc.Rsp.E2DListBlogLikes(users), nil
}
