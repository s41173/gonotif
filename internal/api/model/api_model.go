package model

import "time"

// Message merepresentasikan satu record pada tabel messages.
type Notif struct {
	ID       int        `gorm:"column:id;primaryKey" json:"id"`
	Customer int        `gorm:"column:customer"             json:"customer"`
	CustName *string    `gorm:"column:custname"             json:"custname"`
	SentTo   string     `gorm:"column:sentto"               json:"sentto"`
	Subject  string     `gorm:"column:subject"              json:"subject"`
	Content  string     `gorm:"column:content"              json:"content"`
	Type     int16      `gorm:"column:type"                 json:"type"`
	Reading  int16      `gorm:"column:reading"              json:"reading"`
	Modul    string     `gorm:"column:modul"                json:"modul"`
	Status   int16      `gorm:"column:status"               json:"status"`
	Campaign int16      `gorm:"column:campaign"             json:"campaign"`
	Created  *time.Time `gorm:"column:created"              json:"created_at"`
	Deleted  *time.Time `gorm:"column:deleted"              json:"deleted_at"`
	Updated  *time.Time `gorm:"column:updated"              json:"updated_at"`
}

func (Notif) TableName() string {
	return "notif"
}
