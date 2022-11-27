package model

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"locallife/pkg/conf"
	"locallife/pkg/log"
)

type dBConf struct {
	DB       *gorm.DB
	connName string
}

var dbMap map[string]*dBConf

func NewDb(name string) (*gorm.DB, error) {
	if d, ok := dbMap[name]; ok {
		return d.DB, nil
	}
	fmt.Println(dbMap)
	log.Error("get database failed.", log.String("db_name", name))
	return nil, errors.New("db name no such")
}

func (d *dBConf) getDbParams() string {
	var preConnName = "local_life.mysql." + d.connName + "."
	address := conf.GetString(preConnName + "address")
	username := conf.GetString(preConnName + "username")
	passwd := conf.GetString(preConnName + "password")
	dbName := conf.GetString(preConnName + "db_name")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true", username, passwd, address, dbName)
}

func (d *dBConf) setDbParams() {

}

func (d *dBConf) InitDb() {
	dbParams := d.getDbParams()

	mysqlConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	var err error
	fmt.Println(dbParams)
	d.DB, err = gorm.Open(mysql.Open(dbParams), mysqlConfig)
	if err != nil {
		log.Fatal("open database failed.", log.String("db_name", d.connName))
	}
	d.setDbParams()
	log.Info("open database success.")
}

func Init() {
	dbMap = make(map[string]*dBConf)
	instances, ok := conf.Get("local_life.mysql").(map[string]interface{})
	if !ok {
		panic("mysql conf failed.")
	}
	for dbName, _ := range instances {
		d := &dBConf{connName: dbName}
		d.InitDb()
		dbMap[dbName] = d
	}
}
