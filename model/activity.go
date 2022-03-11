package model

import "time"

type Activity struct {
	BaseColumn
	ActivityName      string    `gorm:"column:activity_name" json:"activity_name" form:"activity_name"`                   //活动名称
	ActivityGoodId    uint      `gorm:"column:activity_good_id" json:"activity_good_id" form:"activity_good_id"`          //活动商品id
	ActivityGoodStock uint      `gorm:"column:activity_good_stock" json:"activity_good_stock" form:"activity_good_stock"` //活动库存
	LimitBuy          uint      `gorm:"column:limit_buy" json:"limit_buy" form:"limit_buy"`
	StartTime         time.Time `gorm:"column:start_time" json:"start_time" form:"start_time"` //活动开始时间
	EndTime           time.Time `gorm:"column:end_time" json:"end_time" form:"end_time"`       //活动结束时间
	Status            uint      `gorm:"column:status" json:"status" form:"status"`
}

func (a *Activity) TableName() string {
	return "sec_activity"
}
