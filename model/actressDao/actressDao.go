package actressDao

import (
	"github.com/jinzhu/gorm"
	"DawnLightGin/model"
	"fmt"
)

type Actress struct {
	gorm.Model
	ActressName string `gorm:"primary_key"`
	Portrait    string
	NoMosaicURL string
}

var db *gorm.DB

func init() {
	db = model.GetDB()
	db.AutoMigrate(&Actress{})
}
func AddActress(actress *Actress) error {
	return db.Create(actress).Error
}
func UpdateActress(actress *Actress) error {
	return db.Model(&Actress{}).Update(actress).Error
}

// 通过ID查询用户
func GetActress(id uint) (actress Actress, err error) {
	err = db.First(&actress, id).Error
	return actress, err
}

// 查询用户列表
func ListActress(start int, pageSize int) (actress []Actress, err error) {
	err = db.Limit(pageSize).Offset(start).Find(&actress).Error
	fmt.Println("数据量：",len(actress))
	return actress, err
}

// 查询用户记录数
func CountActress() (count uint, err error) {
	err = db.Model(&Actress{}).Count(&count).Error
	return count, err
}

// 通过用户ID删除用户
func DeleteActressByID(id uint) (err error) {
	actress := Actress{}
	actress.ID = id
	return DeleteActress(&actress)
}

// 删除用户
func DeleteActress(actress *Actress) (err error) {
	return db.Delete(actress).Error
}
