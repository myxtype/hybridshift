package worker

import (
	"frame/pkg/grace"
	"frame/pkg/worker"
)

func StartWorker() {
	m := worker.NewWorkerManager()

	m.AddWorker(NewTaskCheckWorker())

	grace.WorkerRun(m)
}
