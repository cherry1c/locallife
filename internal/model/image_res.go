package model

import (
	"fmt"
	"locallife/generate/model/locallife"
	"locallife/internal/errcode"
	"locallife/pkg/log"
	"locallife/pkg/model"
)

const dbName = "locallife"

func AddImageToDb(imageName, scene, url string) error {
	db, err := model.NewDb(dbName)
	if err != nil {
		return errcode.ErrNewDb
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
		return result.Error
	}
	fmt.Println(result)
	return nil
}

func GetImageResFromDb(scene string) ([]locallife.ImageRes, error) {
	db, err := model.NewDb(dbName)
	if err != nil {
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
