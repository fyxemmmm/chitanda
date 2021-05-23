package chitanda

import "sync"

type TaskFunc func(params ...interface{})

var (
	taskList chan *TaskExecutor
	once sync.Once
)

func init()  {
	chList := getTaskList() //  任务列表
	go func() {
		for t := range chList {
			doTask(t)
		}
	}()
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

func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
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
		getTaskList() <- NewTaskExecutor(f, params, callback)  // 增加任务队列
	}()
}

