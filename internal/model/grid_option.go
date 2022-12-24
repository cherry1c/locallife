package model

import (
	"locallife/generate/model/locallife"
	"locallife/internal/errcode"
	"locallife/pkg/log"
	"locallife/pkg/model"
	"strconv"
)

func AddGridOptionToDb(name string, image string) error {
	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return errcode.ErrNewDb
	}
	const scene = "homepage"
	gridOption := &locallife.GridOption{
		Name:  name,
		URL:   image,
		Scene: scene,
	}
	result := db.Create(gridOption)
	if result.Error != nil {
		log.Error("add grid option failed.",
			log.String("name", name),
			log.String("image", image),
			log.ErrorF(result.Error))
		return result.Error
	}
	log.Info("add image success.", log.String("gridOptionId", strconv.FormatUint(gridOption.ID, 10)))
	return nil
}

func GetGridOptionFromDb(scene string) ([]locallife.GridOption, error) {
	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return nil, errcode.ErrNewDb
	}
	var gridOptions []locallife.GridOption
	result := db.Model(&locallife.GridOption{}).Where("scene = ?", scene).Find(&gridOptions)
	if result.Error != nil {
		log.Error("get grid option database failed",
			log.String("scene", scene),
			log.ErrorF(result.Error))
		return nil, result.Error
	}
	return gridOptions, nil
}
