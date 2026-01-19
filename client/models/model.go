package models

import "time"

type MsgModel struct {
	UserName  string
	Content   string
	TimeStamp time.Time
	Type      string
}
