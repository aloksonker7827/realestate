package models

import "time"

type BookKeeping struct {
	ID            uint `gorm:"primary_key;AUTO_INCREMENT"`
	ClientID      uint `sql:"REFERENCES api_clients(id)"`
	Version       int
	RequestID     string `gorm:"not null;index:idx_request_id"`
	Status        string `gorm:"type:varchar(50)"`
	RequestData   string `gorm:"type:text"`
	FailureReason string `gorm:"type:text"`
	FailureAction string `gorm:"type:text"`
	UploadedBy    string `gorm:"type:varchar(50)"`
	Action        string `gorm:"type:varchar(50)"`
	Result        string `gorm:"type:varchar(160)"`
	IsRetryable   bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}
