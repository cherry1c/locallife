package homepage

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/util"
	"locallife/pkg/log"
)

type AddImageResRequest struct {
	Name  string `json:"name"`
	Scene string `json:"scene"`
	Url   string `json:"url"`
}

type AddImageResResponse struct {
}

func AddImageRes(context *gin.Context) {
	request := &AddImageResRequest{}
	response := &AddImageResResponse{}
	util.GetRequest(context.Request.Body, request)
	log.Info("input param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
	if err := AddImageResImpl(context, request, response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
	log.Info("output param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
}

type GetSwiperListRequest struct {
	Scene string `json:"scene"`
}

type GetSwiperListResponse struct {
	Image []util.Image `json:"image"`
}

func GetSwiperList(context *gin.Context) {
	request := &GetSwiperListRequest{}
	response := &GetSwiperListResponse{}
	request.Scene = context.Query("scene")
	//util.GetRequest(context.Request.Body, &request)
	log.Info("input param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
	if err := GetSwiperListImpl(context, request, response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
	log.Info("output param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
}

type GridOption struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type GetGridOptionListRequest struct {
}

type GetGridOptionListResponse struct {
	GridOptionList []GridOption `json:"grid_option_list"`
}

func GetGridOptionList(context *gin.Context) {
	request := &GetGridOptionListRequest{}
	response := &GetGridOptionListResponse{}
	log.Info("input param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
	if err := GetGridOptionListImpl(context, request, response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
	log.Info("output param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
}

type AddGridOptionRequest struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type AddGridOptionResponse struct {
}

func AddGridOption(context *gin.Context) {
	request := &AddGridOptionRequest{}
	response := &AddGridOptionResponse{}
	util.GetRequest(context.Request.Body, request)
	log.Info("input param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
	if err := AddGridOptionImpl(context, request, response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
	log.Info("output param",
		log.String("request", util.TypeToString(request)),
		log.String("response", util.TypeToString(response)))
}
