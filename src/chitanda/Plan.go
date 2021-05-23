package chitanda

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type TaskFunc func(params ...interface{})

var (
	taskList chan *TaskExecutor
	once sync.Once

	onceCron sync.Once
	taskCron *cron.Cron  // 定时任务
)

func init()  {
	chList := getTaskList()
	go func() {
		for t := range chList {
			doTask(t)
		}
	}()
}

func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}

func doTask(t *TaskExecutor)  {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}

func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

type TaskExecutor struct {
	f TaskFunc
	p []interface{}
	callback func()
}

func newTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}

func (this *TaskExecutor) Exec() {
	this.f(this.p...)
}


func Task(f TaskFunc, callback func(), params ...interface{})  {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- newTaskExecutor(f, params, callback)  // 增加任务队列
	}()
}

