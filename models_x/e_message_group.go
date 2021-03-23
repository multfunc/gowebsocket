package models_x

import "time"

/**
* @Author junfenghe
* @Description group message
* @Date 2021-03-23 7:45
* @Param
* @return
**/
type MessageGroup struct {
	ID           string    `json:"id"`
	GroupID      string    `json:"group_id"`
	GroupName    string    `json:"group_name"`
	FromID       string    `json:"from_id"`
	FromNickname string    `json:"from_nickname"`
	FromAvatar   string    `json:"from_avatar"` // 发送消息人的头像
	Type         string    `json:"type"`        // img|text|emu:message type
	IsRead       bool      `json:"is_read"`     // read or not
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Updated      int64     `json:"updated"`
	Created      int64     `json:"created"`
	//Type        string `json:"type"` // group_live,group,point2point|直播群聊，普通群聊，点对点
}
