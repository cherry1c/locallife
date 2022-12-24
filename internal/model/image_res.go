package model

import (
	"fmt"
	"locallife/generate/model/locallife"
	"locallife/internal/errcode"
	"locallife/pkg/log"
	"locallife/pkg/model"
	"strconv"
)

const dbLocalLife = "locallife"

func AddImageToDb(imageName, scene, url string) (uint64, error) {
	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return 0, errcode.ErrNewDb
	}
	image := &locallife.ImageRes{
		Name:  imageName,
		Scene: scene,
		URL:   url,
	}
	result := db.Create(image)
	if result.Error != nil {
		log.Error("add image failed.",
			log.String("imageName", imageName),
			log.String("scene", scene),
			log.String("url", url),
			log.ErrorF(result.Error))
		return 0, result.Error
	}
	log.Info("add image success.", log.String("imageId", strconv.FormatUint(image.ID, 10)))
	return image.ID, nil
}

func GetImageResFromDb(scene string) ([]locallife.ImageRes, error) {
	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return nil, errcode.ErrNewDb
	}
	var images []locallife.ImageRes
	result := db.Model(&locallife.ImageRes{}).Where("scene = ?", scene).Find(&images)
	if result.Error != nil {
		log.Error("get image res database failed",
			log.String("scene", scene),
			log.ErrorF(result.Error))
		return nil, result.Error
	}
	return images, nil
}

func GetImageResByIds(ids []uint64) ([]locallife.ImageRes, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return nil, errcode.ErrNewDb
	}
	var images []locallife.ImageRes
	result := db.Model(&locallife.ImageRes{}).Where("id in ?", ids).Find(&images)
	if result.Error != nil {
		log.Error("get image res database failed",
			log.String("ids", fmt.Sprintln(ids)),
			log.ErrorF(result.Error))
		return nil, result.Error
	}
	return images, nil
}
