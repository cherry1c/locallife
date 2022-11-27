package homepage

import (
	"github.com/gin-gonic/gin"
	"locallife/internal/errcode"
	"locallife/internal/model"
)

func checkRequestParam(request *AddImageResRequest) bool {
	if len(request.Name) == 0 ||
		len(request.Scene) == 0 ||
		len(request.Url) == 0 {
		return false
	}
	return true
}

func AddImageResImpl(context *gin.Context, request *AddImageResRequest, response *AddImageResResponse) error {
	if !checkRequestParam(request) {
		return errcode.ErrCheckParam
	}
	if err := model.AddImageToDb(request.Name, request.Scene, request.Url); err != nil {
		return errcode.ErrAddImage
	}
	return nil
}
