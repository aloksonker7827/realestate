package models

import (
	"time"
)

// ApiClient struct denotes the table structure in database.
// This is used to fetch the provider and sub-tag given the api-key.
// Note : We can use gorm.Model to embed default fields, but the ordering
// is not well maintained. so, using all the fields in gorm.Model with some additions here.
type ApiClient struct {
	ID         uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Client     string `gorm:"type:varchar(50)"`
	ApiKey     string `gorm:"type:varchar(100);index:idx_api_key;unique;not null;"`
	Provider   string `gorm:"type:varchar(50)"`
	SubTag     string `gorm:"type:varchar(50)"`
	Enabled    bool
	MaxRetries int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}
