package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rubbercable/mysql-book-mgmt/pkg/config"
	_ "github.com/rubbercable/mysql-book-mgmt/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book())
}

func (b *Book) createBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func getBookByID(I int64)
