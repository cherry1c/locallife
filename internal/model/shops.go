package model

import (
	"fmt"
	"locallife/generate/model/locallife"
	"locallife/internal/errcode"
	"locallife/pkg/log"
	"locallife/pkg/model"
	"strconv"
)

func AddShopToDb(typeId uint64, name, phone, address, businessHours, imageUrl string) (uint64, error) {
	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return 0, errcode.ErrNewDb
	}
	shops := &locallife.Shops{
		Type:          typeId,
		Name:          name,
		Phone:         phone,
		Address:       address,
		BusinessHours: businessHours,
		URL:           imageUrl,
	}
	result := db.Create(shops)
	if result.Error != nil {
		log.Error("add shops failed",
			log.String("typeId", fmt.Sprintf("%d", typeId)),
			log.String("name", name),
			log.String("phone", phone),
			log.String("address", address),
			log.ErrorF(result.Error))
		return 0, result.Error
	}

	log.Info("add shops success.", log.String("shop_id", strconv.FormatUint(shops.ID, 10)))
	return shops.ID, nil
}

func GetShopListFromDb(typeId uint64, offset int, limit int) ([]locallife.Shops, error) {
	db, err := model.NewDb(dbLocalLife)
	if err != nil {
		log.Error("new database failed.", log.ErrorF(err))
		return nil, errcode.ErrNewDb
	}
	var shopList []locallife.Shops
	result := db.Model(&locallife.Shops{}).Where("type = ?", typeId).Limit(limit).Offset(offset).Find(&shopList)
	if result.Error != nil {
		log.Error("get shop list failed",
			log.String("typeId", strconv.FormatUint(typeId, 64)),
			log.String("offset", strconv.Itoa(offset)),
			log.String("limit", strconv.Itoa(limit)),
			log.ErrorF(result.Error))
		return nil, result.Error
	}

	return shopList, nil
}
