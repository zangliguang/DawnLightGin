package mosaicMovieDao

import (
	"DawnLightGin/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type MosaicMovie struct {
	gorm.Model
	Licences   string `gorm:"primary_key"`
	MovieTitle string
	Cover      string
	publishDate string
	Hot        bool
	Vid        string
}

var db *gorm.DB

func init() {
	db = model.GetDB()
	db.AutoMigrate(&MosaicMovie{})
}

func AddMovie(movie *MosaicMovie) error {
	var count int
	_ = db.Model(&MosaicMovie{}).Where("licences = ?", movie.Licences).Count(&count).Error
	if count > 0 {
		return nil
	}
	return db.Create(movie).Error
}
func UpdateMovie(licences string, vid string) error {
	return db.Model(&MosaicMovie{}).Where("licences = ?", licences).Update("vid", vid).Error
}

func ListMovie(start int, pageSize int) (movie []MosaicMovie, err error) {

	err = db.Where("vid<>? ", "notFind").Order("publish_date desc,vid desc").Offset(start).Limit(pageSize).Find(&movie).Error
	return movie, err
}
func QueryMovie(after string, before string) (movie []MosaicMovie) {

	err2 := db.Where("publish_date >= ? And publish_date <= ? And (vid is null or vid=?)", after, before, "").Order("publish_date desc").Find(&movie).Error
	if err2 != nil {
		log.Println("出问题了：", err2)
	}
	fmt.Println(fmt.Sprintf("从%s到%s一共%d个电影", after, before, len(movie)))
	return movie
}
