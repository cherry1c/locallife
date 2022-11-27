package pkg

import (
	"locallife/pkg/conf"
	"locallife/pkg/log"
	"locallife/pkg/model"
)

func Init() {
	conf.Init()
	log.Init()
	model.Init()
}
