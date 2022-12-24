package shoplist

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/util"
)

type AddShopRequest struct {
	TypeId        uint64 `json:"type_id"`
	Image         string `json:"image"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	BusinessHours string `json:"business_hours"`
}

type AddShopResponse struct {
}

func AddShop(context *gin.Context) {
	request := AddShopRequest{}
	response := AddShopResponse{}

	util.GetRequest(context.Request.Body, &request)
	if err := AddShopImpl(context, &request, &response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err := util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
}

type GetShopListRequest struct {
	TypeId uint64 `json:"id"`
	Offset int32  `json:"offset"`
	Limit  int32  `json:"limit"`
}

type Shop struct {
	Image         string `json:"image"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	BusinessHours string `json:"business_hours"`
}

type GetShopListResponse struct {
	ShopList []Shop `json:"shop_list"`
	Total    int32  `json:"total"`
}

func GetShopList(context *gin.Context) {
	request := GetShopListRequest{}
	response := GetShopListResponse{}
	var err error
	if request.TypeId, err = util.StrToUint64(context.Query("id")); err != nil {
		return
	}
	if request.Offset, err = util.StrToInt32(context.Query("offset")); err != nil {
		return
	}
	if request.Limit, err = util.StrToInt32(context.Query("limit")); err != nil {
		return
	}

	if err = GetShopListImpl(context, &request, &response); err != nil {
		util.PackageFailed(context, err)
		return
	}
	if err = util.PackageSuccess(context, response); err != nil {
		util.PackageFailed(context, err)
	}
}
