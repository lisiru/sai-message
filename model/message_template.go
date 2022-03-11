package model

import "time"

// 消息模板信息
type MessageTemplate struct {
	Id             int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name           string    `gorm:"column:name" db:"name" json:"name" form:"name"`                                                 //标题
	AuditStatus    int8      `gorm:"column:audit_status" db:"audit_status" json:"audit_status" form:"audit_status"`                 //当前消息审核状态： 10.待审核 20.审核成功 30.被拒绝
	FlowId         string    `gorm:"column:flow_id" db:"flow_id" json:"flow_id" form:"flow_id"`                                     //工单ID
	MsgStatus      int8      `gorm:"column:msg_status" db:"msg_status" json:"msg_status" form:"msg_status"`                         //当前消息状态：10.新建 20.停用 30.启用 40.等待发送 50.发送中 60.发送成功 70.发送失败
	CronTaskId     int64     `gorm:"column:cron_task_id" db:"cron_task_id" json:"cron_task_id" form:"cron_task_id"`                 //定时任务Id (xxl-job-admin返回)
	CronCrowdPath  string    `gorm:"column:cron_crowd_path" db:"cron_crowd_path" json:"cron_crowd_path" form:"cron_crowd_path"`     //定时发送人群的文件路径
	ExpectPushTime string    `gorm:"column:expect_push_time" db:"expect_push_time" json:"expect_push_time" form:"expect_push_time"` //期望发送时间：0:立即发送 定时任务以及周期任务:cron表达式
	IdType         int       `gorm:"column:id_type" db:"id_type" json:"id_type" form:"id_type"`                                     //消息的发送ID类型：10. userId 20.did 30.手机号 40.openId 50.email
	SendChannel    int       `gorm:"column:send_channel" db:"send_channel" json:"send_channel" form:"send_channel"`                 //消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序
	TemplateType   int       `gorm:"column:template_type" db:"template_type" json:"template_type" form:"template_type"`             //10.运营类 20.技术类接口调用
	MsgType        int       `gorm:"column:msg_type" db:"msg_type" json:"msg_type" form:"msg_type"`                                 //10.通知类消息 20.营销类消息 30.验证码类消息
	MsgContent     string    `gorm:"column:msg_content" db:"msg_content" json:"msg_content" form:"msg_content"`                     //消息内容 占位符用{$var}表示
	SendAccount    int       `gorm:"column:send_account" db:"send_account" json:"send_account" form:"send_account"`                 //发送账号 一个渠道下可存在多个账号
	Creator        string    `gorm:"column:creator" db:"creator" json:"creator" form:"creator"`                                     //创建者
	Updator        string    `gorm:"column:updator" db:"updator" json:"updator" form:"updator"`                                     //更新者
	Auditor        string    `gorm:"column:auditor" db:"auditor" json:"auditor" form:"auditor"`                                     //审核人
	Team           string    `gorm:"column:team" db:"team" json:"team" form:"team"`                                                 //业务方团队
	Proposer       string    `gorm:"column:proposer" db:"proposer" json:"proposer" form:"proposer"`                                 //业务方
	IsDeleted      int8      `gorm:"column:is_deleted" db:"is_deleted" json:"is_deleted" form:"is_deleted"`                         //是否删除：0.不删除 1.删除
	UpdatedAt      time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`                         //更新时间
	CreatedAt      time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`                         //创建时间
}

func (m *MessageTemplate) TableName() string {
	return "message_template"
}
