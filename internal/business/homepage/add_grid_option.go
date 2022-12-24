package homepage

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/errcode"
	"locallife/internal/model"
)

func checkAddGridOptionRequestParam(request *AddGridOptionRequest) bool {
	if len(request.Name) == 0 {
		return false
	}
	return true
}

func AddGridOptionImpl(context *gin.Context, request *AddGridOptionRequest, response *AddGridOptionResponse) error {
	if !checkAddGridOptionRequestParam(request) {
		return errcode.ErrCheckParam
	}
	if err := model.AddGridOptionToDb(request.Name, request.Image); err != nil {
		return errcode.ErrGridOption
	}
	return nil
}
