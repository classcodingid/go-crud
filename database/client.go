package database

import (
	"go-crud/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

// konek ke database mysql
func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Tidak bisa terhubung ke Database!")
	}
	log.Println("Berhasil terhubung kedatabase...")
}

// migrasi
func Migrate() {
	Instance.AutoMigrate(&models.Product{})
	log.Println("Migration Database Selesai...")
}
