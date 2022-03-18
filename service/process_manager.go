package service

import (
	"context"
	"sai/common"
	"sai/store"
)

type Manager struct {
	Processor []Processor
}


func (m *Manager) AddProcess(processor Processor)  {
	m.Processor=append(m.Processor,processor)

}

// 对责任链各个职责功能函数管理
func NewManager(store store.Factory) *Manager {
	return &Manager{Processor: []Processor{
		&PreParamCheckAction{},&AssembleAction{
			store: store,
		}, &AfterParamCheckAction{} , &SendMqAction{},
	}}
}


func (m *Manager) Run(ctx context.Context,processContext *common.ProcessContext) error {
	for _,v:=range m.Processor{
		err:=v.Process(ctx,processContext)
		if err != nil {
			return  err
		}
	}
	return nil
}
