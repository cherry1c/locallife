package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"locallife/internal/errcode"
	"locallife/pkg/log"
	"net/http"
	"strconv"
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

func ParseUint(s string, bitSize int) (uint64, error) {
	ans, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		log.Error("parse uint failed", log.String("value", s), log.ErrorF(err))
		return 0, err
	}
	return ans, nil
}

func ParseInt(s string, bitSize int) (int64, error) {
	ans, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		log.Error("parse int failed", log.String("value", s), log.ErrorF(err))
		return 0, err
	}
	return ans, nil
}

func StrToUint64(s string) (uint64, error) {
	return ParseUint(s, 64)
}

func StrToInt64(s string) (int64, error) {
	return ParseInt(s, 64)
}

func StrToUint32(s string) (uint32, error) {
	ui64, err := ParseUint(s, 32)
	return uint32(ui64), err
}

func StrToInt32(s string) (int32, error) {
	i64, err := ParseUint(s, 32)
	return int32(i64), err
}

func TypeToString(t interface{}) string {
	b, err := json.Marshal(t)
	if err != nil {
		log.Error("marshal failed", log.ErrorF(err))
		return ""
	}
	return string(b)
}
