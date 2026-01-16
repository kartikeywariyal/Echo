package model

import "time"

type MsgModel struct {
	UserName  string
	content   string
	TimeStamp time.Time
	Type      string
}
