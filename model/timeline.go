package model

import (
	"github.com/cauakath/timeline-server/enum"
	"gorm.io/gorm"
)

type Timeline struct {
	Id       int               `gorm:"type:int;primaryKey;autoIncrement;not null" json:"id"`
	Title    string            `gorm:"type:varchar(255);not null" json:"title"`
	Location string            `gorm:"type:varchar(255);not null" json:"location"`
	Start    string            `gorm:"type:varchar(255);not null" json:"start"`
	End      string            `gorm:"type:varchar(255)" json:"end"`
	Type     enum.TimelineType `gorm:"type:varchar(255);not null" json:"type"`
	*gorm.Model
}
