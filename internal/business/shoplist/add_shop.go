package shoplist

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/errcode"
	"locallife/internal/model"
)

func AddShopImpl(context *gin.Context, request *AddShopRequest, response *AddShopResponse) error {

	if _, err := model.AddShopToDb(request.TypeId, request.Phone, request.Phone, request.Address, request.BusinessHours, request.Image); err != nil {
		return errcode.ErrAddShop
	}
	return nil
}
