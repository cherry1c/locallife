package homepage

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/util"
)

type AddImageResRequest struct {
	Name  string `json:"name"`
	Scene string `json:"scene"`
	Url   string `json:"url"`
}

type AddImageResResponse struct {
}

func AddImageRes(context *gin.Context) {
	request := AddImageResRequest{}
	response := AddImageResResponse{}
	util.GetRequest(context.Request.Body, &request)
	if err := AddImageResImpl(context, &request, &response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
}

type GetSwiperListRequest struct {
	Scene string `json:"scene"`
}

type GetSwiperListResponse struct {
	Image []util.Image `json:"image"`
}

func GetSwiperList(context *gin.Context) {
	request := GetSwiperListRequest{}
	response := GetSwiperListResponse{}
	request.Scene = context.Query("scene")
	//util.GetRequest(context.Request.Body, &request)
	if err := GetSwiperListImpl(context, &request, &response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
}

func GetGridList(context *gin.Context) {

}
