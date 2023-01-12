package entity

import (
	// "golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("history_student.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = database

	// password, err := bcrypt.GenerateFromPassword([]byte("123"), 14)
	// if err != nil {
	// 	panic("failed to generate password")
	// }

	// db.Model(&Admin{}).Create(&Admin{
	// 	Aname:     "นางสาววันวิสา เถาโท",
	// 	Aemail:    "w@gmail.com",
	// 	Apassword: string(password),
	// })
	// db.Model(&Admin{}).Create(&Admin{
	// 	Aname:     "นายอภิสิทธิ์ วงศ์วิศิษฐ์",
	// 	Aemail:    "a@gmail.com",
	// 	Apassword: string(password),
	// })
	// db.Model(&Admin{}).Create(&Admin{
	// 	Aname:     "นายธีรวัฒน์ กูดกิ่ง",
	// 	Aemail:    "t@gmail.com",
	// 	Apassword: string(password),
	// })
	// db.Model(&Admin{}).Create(&Admin{
	// 	Aname:     "นายพีรพล นนทคําจันทร์",
	// 	Aemail:    "p@gmail.com",
	// 	Apassword: string(password),
	// })
	// db.Model(&Admin{}).Create(&Admin{
	// 	Aname:     "นายธนเดช เชิดในเมือง",
	// 	Aemail:    "k@gmail.com",
	// 	Apassword: string(password),
	// })
	// db.Model(&Admin{}).Create(&Admin{
	// 	Aname:     "ณัฐพล จงเกษกรรม",
	// 	Aemail:    "n@gmail.com",
	// 	Apassword: string(password),
	// })
}
