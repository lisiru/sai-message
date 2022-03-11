package mysql

import (
	"sai/model"
	"sai/pkg/code"
	"sai/pkg/util/gormutil"
	"context"
	"github.com/marmotedu/errors"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{db: ds.db}
}
func (u *users) Create(ctx context.Context, user *model.User) error {
	return u.db.Create(&user).Error
}

// Update updates an user account information.
func (u *users) Update(ctx context.Context, user *model.User) error {
	return u.db.Save(user).Error
}

func (u *users) Get(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("name = ? and status = 1", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
		}

		return nil, errors.WithCode(code.ErrDatabase, err.Error())
	}

	return user, nil
}

func (u *users) GetUserByWhere(ctx context.Context,where map[string]interface{}) (*model.User,error)  {
	user:=&model.User{}
	err:=u.db.Where(where).First(&user).Error
	if err !=nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(code.ErrUserNotFound,err.Error())
		}
		return nil,errors.WithCode(code.ErrDatabase,err.Error())
	}
	return user,nil
}

func (u *users) List(ctx context.Context, limit int64,offset int64) (*model.UserList, error) {
	ret := &model.UserList{}
	ol := gormutil.Unpointer(&offset, &limit)
	where :=model.User{}
	//whereNot:=model.User{
	//	IsAdmin :0,
	//}
	d := u.db.Where(where).
		//Not(whereNot).
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
