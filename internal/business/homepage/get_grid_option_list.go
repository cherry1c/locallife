package homepage

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/model"
)

func GetGridOptionListImpl(context *gin.Context, request *GetGridOptionListRequest, response *GetGridOptionListResponse) error {
	const scene = "homepage"
	gridOptions, err := model.GetGridOptionFromDb(scene)
	if err != nil {
		return err
	}
	for _, v := range gridOptions {
		response.GridOptionList = append(response.GridOptionList, GridOption{Id: v.ID, Name: v.Name, Image: v.URL})
	}

	// sort grid option

	return nil
}
