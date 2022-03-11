package workerpool

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type Pool struct {
	capacity int // workerpool 大小
	active chan struct{}
	task chan Task
	wg sync.WaitGroup // 用于在pool 销毁时等待所有worker退出
	quit chan struct{} // 通知各个worker 退出的信号channel
	preAlloc bool // 是否在创建pool的时候就预创建workers 默认值 false
	block bool

}
var (
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool") // workerpool中任务已满，没有空闲goroutine用于处理新任务
	ErrWorkerPoolFreed    = errors.New("workerpool freed")       // workerpool已终止运行
)
const (
	defaultCapacity = 10
	maxCapacity     = 50
)

func NewPool(capacity int,opts ...Option) *Pool  {
	if capacity <=0{
		capacity=defaultCapacity
	}
	if capacity >maxCapacity{
		capacity=maxCapacity
	}
	p:=&Pool{
		capacity: capacity,
		task: make(chan Task),
		quit: make(chan struct{}),
		active: make(chan struct{},capacity),
	}
	for _,opt:=range opts {
		opt(p)
	}
	if p.preAlloc {
		for i := 0; i <capacity ; i++ {
			p.newWorker(i)
			p.active<- struct{}{}
		}
	}
	go p.run()
	return p
}

func (p *Pool) run()  {
	idx:=len(p.active)
	if !p.preAlloc {
	loop:
		for t:=range p.task{
			p.returnTask(t)
			select {
			case <-p.quit:
				return
			case p.active<- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}

		}
	}
	for  {
		select {
		case <-p.quit:
			return
		case p.active<- struct{}{}:
			idx++
			p.newWorker(idx)


		}

	}
}

func (p *Pool) newWorker(i int)  {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err:=recover();err!=nil{
				<-p.active
			}
			p.wg.Done()
		}()

		for  {
			select {
			case <-p.quit:
				<-p.active
				return
			case t:=<-p.task:
				t()


			}
		}
	}()
}

type Task func()


func (p *Pool) Schedule(t Task) error  {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.task <-t:
		return nil
	default:
		if p.block {
			p.task <- t
			return nil
		}
		return ErrNoIdleWorkerInPool
	}


}
func (p *Pool) returnTask(t Task) {
	go func() {
		p.task <- t
	}()
}

func (p *Pool) Free() {
	close(p.quit) // make sure all worker and p.run exit and schedule return error
	p.wg.Wait()
	fmt.Printf("workerpool freed(preAlloc=%t)\n", p.preAlloc)
}
