package mysql

import (
	"context"
	"github.com/marmotedu/errors"
	"gorm.io/gorm"
	"sai/model"
	"sai/pkg/code"
)

type activity struct {
	db *gorm.DB
}

func newActivity(ds *datastore) *activity  {
	return &activity{db: ds.db}
}

func (a *activity)Create(ctx context.Context,activity *model.Activity) error  {
	return a.db.Create(activity).Error
}

func (a *activity) GetActivityByWhere(ctx context.Context,where map[string]interface{}) (*model.Activity,error)  {
	activityModel:=model.Activity{}
	err:=a.db.Where(where).First(&activityModel).Error
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,errors.WithCode(code.ErrActivityNotFound,err.Error())
		}
		return nil, errors.WithCode(code.ErrDatabase,err.Error())

	}
	return &activityModel,nil
}
