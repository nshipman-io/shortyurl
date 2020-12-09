package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

type ShortURL struct {
	gorm.Model
	id int
	OriginalUrl string
}

var Db *gorm.DB
func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=shortyurl dbname=shortyurl password=notsosafe sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&ShortURL{})
}

func CreateUrl(url *ShortURL)(err error){
	err = Db.Create(&url).Error
	if err != nil {
		return err
	}
	return
}

func GetUrl(id int) (*ShortURL, error) {
	url := ShortURL{}
	err := Db.First(&url, id).Error

	return &url, err
}