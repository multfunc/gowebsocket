package models_x

import (
	"gorm.io/gorm"
	"log"
	"time"
)

/**
* @Author junfenghe
* @Description 直播群
* @Date 2021-03-21 13:26
* @Param
* @return
**/
type GroupLive struct {
	ID        string         `json:"id"gorm:"primary_key;comment:group_live_id"`
	Number    string         `json:"number"gorm:"unique;comment:房间号"`
	Name      string         `json:"name"gorm:"comment:房间名"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Updated   int64          `json:"updated"gorm:"autoUpdateTime:milli"`
	Created   int64          `json:"created"gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	IAccounts []*IAccount    `json:"i_accounts"gorm:"many2many:i_account__group_live;constraint:OnDelete:CASCADE;comment:这个群里有哪些用户"`
}

func (groupLive *GroupLive) BeforeCreate(tx *gorm.DB) error {
	if groupLive.ID == "" {
		return BeforeCreateUpdateID(tx)
	}
	return nil
}

func InitGroupLive() {
	err := db.AutoMigrate(&GroupLive{})
	if err != nil {
		log.Fatal(err)
	}
}
