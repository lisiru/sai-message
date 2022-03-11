package model

import "time"

// 短信记录信息
type SmsRecord struct {
	Id                int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	MessageTemplateId int64     `gorm:"column:message_template_id" db:"message_template_id" json:"message_template_id" form:"message_template_id"` //消息模板ID
	Phone             int64     `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`                                                         //手机号
	SupplierId        int8      `gorm:"column:supplier_id" db:"supplier_id" json:"supplier_id" form:"supplier_id"`                                 //发送短信渠道商的ID
	SupplierName      string    `gorm:"column:supplier_name" db:"supplier_name" json:"supplier_name" form:"supplier_name"`                         //发送短信渠道商的名称
	MsgContent        string    `gorm:"column:msg_content" db:"msg_content" json:"msg_content" form:"msg_content"`                                 //短信发送的内容
	SeriesId          string    `gorm:"column:series_id" db:"series_id" json:"series_id" form:"series_id"`                                         //下发批次的ID
	ChargingNum       int8      `gorm:"column:charging_num" db:"charging_num" json:"charging_num" form:"charging_num"`                             //计费条数
	ReportContent     string    `gorm:"column:report_content" db:"report_content" json:"report_content" form:"report_content"`                     //回执内容
	Status            int8      `gorm:"column:status" db:"status" json:"status" form:"status"`                                                     //短信状态： 10.发送 20.成功 30.失败
	SendDate          int       `gorm:"column:send_date" db:"send_date" json:"send_date" form:"send_date"`                                         //发送日期：20211112
	UpdatedAt         time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                                     //更新时间
	CreatedAt         time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                                     //创建时间
}

func (s *SmsRecord) TableName() string  {
	return "sms_record"
}