package controller

import "github.com/gin-gonic/gin"

type FollowerController interface {
	Follow(ctx *gin.Context)
	IsFollow(ctx *gin.Context)
	Common(ctx *gin.Context)
}

func NewFollowerController() FollowerController {
	return FollowerControllerImp{}
}

type FollowerControllerImp struct {
}

func (f FollowerControllerImp) Follow(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f FollowerControllerImp) IsFollow(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (f FollowerControllerImp) Common(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
