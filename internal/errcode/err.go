package errcode

import "errors"

var (
	ErrNormal     = errors.New("OK")
	ErrService    = errors.New("服务错误")
	ErrNewDb      = errors.New("获取数据库失败")
	ErrGetImaData = errors.New("获取图片数据失败")
	ErrCheckParam = errors.New("检查参数失败")
	ErrAddImage   = errors.New("添加图片失败")
	ErrGridOption = errors.New("添加网格选项失败")
	ErrAddShop    = errors.New("添加商铺失败")
)
