package models

import "time"

type Article struct {
	ID       int    `json:"id" binding:"-" gorm:"primarykey;autoIncrement"`
	Title    string    `json:"title" binding:"required"`
	CateID   string    `json:"cate_id" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	CoverImg string    `json:"cover_img"`
	State    string    `json:"state" binding:"required"`
	PubDate  time.Time `json:"pub_time"`
}
