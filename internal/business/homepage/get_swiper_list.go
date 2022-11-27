package homepage

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/model"
	"locallife/internal/util"
)

func GetSwiperListImpl(context *gin.Context, request *GetSwiperListRequest, response *GetSwiperListResponse) error {
	if len(request.Scene) == 0 {
		return nil
	}
	images, err := model.GetImageResFromDb(request.Scene)
	if err != nil {
		return err
	}
	for _, v := range images {
		response.Image = append(response.Image, util.Image{Id: v.ID, Scene: v.Scene, Url: v.URL})
	}
	return nil
}
