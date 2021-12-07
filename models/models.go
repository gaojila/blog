package models

import (
	"blog/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreateOn  int `json:"create_on`
	ModifieOn int `json:"modifie_on`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section `database`: %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
}
