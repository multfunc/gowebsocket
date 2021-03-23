package models_x

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type IAccount struct {
	ID         string         `json:"id"gorm:"primary_key;comment:user_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Updated    int64          `json:"updated"gorm:"autoUpdateTime:milli"`
	Created    int64          `json:"created"gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Friends    []*Friend      `json:"friends"gorm:"many2many:i_account__group_live;constraint:OnDelete:CASCADE;comment:用户的好友"`
	GroupLives []*GroupLive   `json:"group_lives"gorm:"many2many:i_account__group_live;constraint:OnDelete:CASCADE;comment:这个用户有哪些群"`
}

func (iAccount *IAccount) BeforeCreate(tx *gorm.DB) error {
	if iAccount.ID == "" {
		return BeforeCreateUpdateID(tx)
	}
	return nil
}

func InitIAccount() {
	err := db.AutoMigrate(&IAccount{})
	if err != nil {
		log.Fatal(err)
	}
}
