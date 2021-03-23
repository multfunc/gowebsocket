package models_x

import (
	"gorm.io/gorm"
	"log"
	"time"
)

/**
* @Author junfenghe
* @Description friends 该数据需要定期更新昵称和职位和单位
* @Date 2021-03-23 13:49
* @Param
* @return
**/
type Friend struct {
	ID             string    `json:"id"`
	FriendID       string    `json:"friend_id"gorm:"comment:好友 ID"`
	FriendNickname string    `json:"friend_nickname"gorm:"comment:好友昵称"`
	FriendJob      string    `json:"friend_job"gorm:"comment:好友职位"`
	FriendOrg      string    `json:"friend_org"gorm:"comment:好友所在机构"`
	Status         string    `json:"status"gorm:"comment:好友|黑名单|特别关注|临时"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Updated        int64     `json:"updated"`
	Created        int64     `json:"created"`
}

func (friend *Friend) BeforeCreate(tx *gorm.DB) error {
	if friend.ID == "" {
		return BeforeCreateUpdateID(tx)
	}
	return nil
}

func InitFriend() {
	err := db.AutoMigrate(&Friend{})
	if err != nil {
		log.Fatal(err)
	}
}
