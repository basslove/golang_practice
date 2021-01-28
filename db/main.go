package main

import (
	//"encoding/csv"
	//"fmt"
	//_ "github.com/lib/pq"
	//"github.com/jinzhu/gorm"
	//"io"
	//"log"
	//"golang_sample/db/models"
	//"os"
	//"strconv"
)

func main() {
	//db, err := gorm.Open("postgres", dataResourceName)
	//if err != nil {
	//	log.Fatalln("接続失敗", err)	dataResourceName := "postgres://postgres:ppassword@golang_sample_database_1:5432/golang_sample_development?sslmode=disable"
	//
	//}
	//defer db.Close()
	//
	//f, err := os.Open("db/seeds/csv/prefecture.csv")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//r := csv.NewReader(f)
	//r.TrimLeadingSpace = true
	//
	//for i := 0; ; i++ {
	//	record, err := r.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if i == 0 {
	//		continue
	//	}
	//
	//	toInt64, _ := strconv.ParseInt(record[0], 10, 64)
	//	toBool, _ := strconv.ParseBool(record[2])
	//
	//	var prefecture models.Prefecture
	//	if err := db.First(&prefecture, toInt64).Error; err != nil {
	//		prefecture = models.Prefecture{ID: toInt64, Name: record[1], IsActive: toBool}
	//		db.Create(&prefecture)
	//	}
	//}
	//
	//var prefectures []models.Prefecture
	//if err = db.Where("id >= ?", 1).Find(&prefectures).Error; err != nil {
	//	log.Fatalln("取得失敗", err)
	//}
	//fmt.Println(prefectures)
}