package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"locallife/internal/errcode"
	"locallife/pkg/log"
	"net/http"
)

func StructToMap(str interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(str)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	if err = json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func PackageSuccess(ctx *gin.Context, response interface{}) error {
	m, err := StructToMap(response)
	if err != nil {
		log.Error("PackageSuccess failed.",
			log.String("response", fmt.Sprintln(response)),
			log.ErrorF(err))
		return err
	}
	m["code"] = errcode.ErrNormal.Error()
	ctx.JSON(http.StatusOK, m)
	return nil
}

func PackageFailed(ctx *gin.Context, response interface{}) {
	m := make(map[string]string, 1)
	err, ok := response.(error)
	if !ok {
		err = errcode.ErrService
	}
	m["code"] = err.Error()
	ctx.JSON(http.StatusForbidden, m)
}

func GetRequest(reader io.Reader, request interface{}) {
	body, err := io.ReadAll(reader)
	if err != nil {
		log.Error("read request failed.", log.ErrorF(err))
		return
	}
	if len(body) == 0 {
		log.Warn("input body is empty.")
		return
	}

	if err = json.Unmarshal(body, &request); err != nil {
		log.Error("unmarshal struct failed.",
			log.String("body", string(body)),
			log.ErrorF(err))
	}
}
