package model

type User struct {
	BaseColumn
	Username string `gorm:"column:username;NOT NULL" json:"username" validate:"omitempty"`
	Uid      uint64 `gorm:"column:uid" json:"uid" validate:"omitempty"`
	Phone    string `gorm:"column:phone" json:"phone" validate:"omitempty"`
	LoginNum uint   `gorm:"column:login_num" json:"login_num" validate:"omitempty"`
}

// UserList is the whole list of all users which have been stored in stroage.
type UserList struct {
	TotalCount int64 `json:"total_count,omitempty"`

	Items []*User `json:"items"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "sec_users"
}
