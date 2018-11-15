package work

import (
	"demo/src/execute"
	"demo/src/model"
)

var workJobs chan model.Task = make(chan model.Task, 1000)

func Start() {
	//TODO 监听chan 变化，当有变化时执行任务
	for{
		select {
		case task := <-workJobs:
			execute.Execute(task)
		}
	}
}

func Stop() {
	close(workJobs)
}


func AddTaskToWork(task model.Task) {
	workJobs<- task
}
