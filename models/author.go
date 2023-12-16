package models

import "time"

type Author struct {
	IDAuthor     uint       `gorm:"primary_key" json:"id"`
	NameAuthor   string     `gorm:"type:varchar(100)" json:"name_author"`
	GenderAuthor string     `gorm:"type:char(1)" json:"gender_author"`
	EmailAuthor  string     `gorm:"type:varchar(100)" json:"email_author"`
	AgeAuthor    int        `gorm:"type:tinyint(3)" json:"age_author"`
	CreateAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}