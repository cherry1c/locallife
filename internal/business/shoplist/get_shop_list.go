package shoplist

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/model"
)

func GetShopListImpl(context *gin.Context, request *GetShopListRequest, response *GetShopListResponse) error {
	shopList, err := model.GetShopListFromDb(request.TypeId, int(request.Offset), int(request.Limit))
	if err != nil {
		return err
	}

	response.Total = int32(len(shopList))
	for _, v := range shopList {
		response.ShopList = append(response.ShopList, Shop{
			Name:          v.Name,
			Image:         v.URL,
			Phone:         v.Phone,
			Address:       v.Address,
			BusinessHours: v.BusinessHours,
		})
	}
	return nil
}
