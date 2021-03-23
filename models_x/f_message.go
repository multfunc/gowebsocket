package models_x

import "time"

/**
* @Author junfenghe
* @Description friend message
* @Date 2021-03-23 7:45
* @Param
* @return
**/
type Message struct {
	ID           string    `json:"id"`
	FromID       string    `json:"from_id"`
	FromNickname string    `json:"from_nickname"`
	FromAvatar   string    `json:"from_avatar"` // 发送消息人的头像
	FromOrgID    string    `json:"from_org_id"`
	FromOrgName  string    `json:"from_org_name"`
	FromJob      string    `json:"from_job"`
	ToID         string    `json:"to_id"`
	ToNickname   string    `json:"to_nickname"`
	ToAvatar     string    `json:"to_avatar"`
	ToOrgID      string    `json:"to_org_id"`
	ToOrgName    string    `json:"to_org_name"`
	ToJob        string    `json:"to_job"`
	Type         string    `json:"type"`    // img|text|emu:message type
	IsRead       bool      `json:"is_read"` // read or not
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Updated      int64     `json:"updated"`
	Created      int64     `json:"created"`
	//Type        string `json:"type"` // group_live,group,point2point|直播群聊，普通群聊，点对点
}
