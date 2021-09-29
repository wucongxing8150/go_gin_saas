package Models

import (
	"go/Cores/mysql"
	"time"
)
var db = mysql.DB
type Table struct{
	Id          int `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	UserName    string `json:"user_name"`
	UserPhone   string `json:"user_phone"`
	CreateTime time.Time `json:"create_time" gorm:"type:datetime;column:create_time;default:null" description:"创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"type:datetime;column:update_time;default:null" description:"修改时间"`
}

func Get(maps interface {}) (table []Table) {
	db.Model(&Table{}).Where(maps).Find(&table)
	return
}
func Search(maps interface {}) (table []Table) {
	db.Model(&Table{}).Where(maps).Find(&table)
	return
}
func SearchPage(pageNum int, pageSize int, maps interface {}) (table []Table) {
	db.Model(&Table{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&table)
	return
}
func GetTotal(maps interface {}) (count int){
	db.Model(&Table{}).Where(maps).Count(&count)
	return
}
func Add(Data map[string]interface{}) bool{
	table := Table {
		UserName : Data["user_name"].(string),
		UserPhone : Data["user_phone"].(string),
	}
	db.Model(&Table{}).Create(&table)
	return !db.NewRecord(table)
}
func EditId(id int, data interface {}) bool {
	db.Model(&Table{}).Where("id = ?", id).Updates(data)
	return true
}
func EditMap(maps interface {}, data interface {}) bool {
	db.Model(&Table{}).Where(maps).Updates(data)
	return true
}