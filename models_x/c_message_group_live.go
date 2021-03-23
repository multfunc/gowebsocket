package models_x

import (
	"time"
)

/**
* @Author junfenghe
* @Description group message
* @Date 2021-03-23 7:45
* @Param
* @return
**/
type MessageGroupLive struct {
	ID            string    `json:"id"`
	GroupLiveID   string    `json:"group_live_id"`
	GroupLiveName string    `json:"group_live_name"`
	FromID        string    `json:"from_id"`
	FromNickname  string    `json:"from_nickname"`
	FromOrgID     string    `json:"from_org_id"`
	FromOrgName   string    `json:"from_org_name"`
	FromJob       string    `json:"from_job"`
	Type          string    `json:"type"`    // img|text|emu:message type
	IsRead        bool      `json:"is_read"` // read or not
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Updated       int64     `json:"updated"`
	Created       int64     `json:"created"`
	//Type        string `json:"type"` // group_live,group,point2point|直播群聊，普通群聊，点对点
}
