package models

// Platform model
type Platform struct {
	ID      int    `gorm:"primary_key;auto_increment;not null"`
	Name    string `gorm:"type:varchar(255);null"`
	Profile string `gorm:"type:varchar(255);null"`
}
