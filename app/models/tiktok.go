package models

import "time"

// Tiktok model
type Tiktok struct {
	ID        int       `gorm:"primary_key;auto_increment;not null"`
	Profile   string    `gorm:"type:varchar(255);null"`
	Username  string    `gorm:"type:varchar(255);null"`
	Followers string    `gorm:"type:varchar(255);null"`
	Following string    `gorm:"type:varchar(255);null"`
	Likes     string    `gorm:"type:varchar(255);null"`
	PostCount string    `gorm:"type:varchar(255);null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
