package noMosaicMovieDao

import (
	"github.com/jinzhu/gorm"
	"DawnLightGin/model"
	"log"
	"fmt"
)

type NoMosaicMovie struct {
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
	db.AutoMigrate(&NoMosaicMovie{})
}


func ListMovie(start int , pageSize int) (movie []NoMosaicMovie, err error) {

	err = db.Where("vid<>?","notFind").Limit(pageSize).Offset(start).Order("publish_date desc,vid desc").Find(&movie).Error
	return movie, err
}

func AddMovie(movie *NoMosaicMovie) error {
	var count int
	_ = db.Model(&NoMosaicMovie{}).Where("licences = ?", movie.Licences).Count(&count).Error
	if count > 0 {
		return nil
	}
	return db.Create(movie).Error
}
func UpdateMovie(licences string, vid string) error {
	return db.Model(&NoMosaicMovie{}).Where("licences = ?", licences).Update("vid", vid).Error
}

func QueryMovie(after string, before string) (movie []NoMosaicMovie) {

	err2 := db.Where("publish_date >= ? And publish_date <= ? And vid =? And hot=1", after, before, "").Order("publish_date desc").Find(&movie).Error
	//err2 := db.Where("publish_date <= ? And publish_date >= ? And vid = ?", "2018-12-14","2011-12-14","").Order("publish_date desc").Find(&movie).Error
	if err2 != nil {
		log.Println("出问题了：", err2)
	}
	fmt.Println(fmt.Sprintf("从%s到%s一共%d个电影", after, before, len(movie)))

	return movie
}
