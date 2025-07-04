package g_pool

import "sync"

type ITask interface {
	Run()
}

type IGoTaskPool interface {
	Start()
	Schedule(task ITask)
	WaitAndStop()
}

type gTaskPool struct {
	workers int
	tasks   chan ITask
	wg      sync.WaitGroup
}

func NewGTaskPool(workers int) IGoTaskPool {
	return &gTaskPool{
		workers: workers,
		tasks:   make(chan ITask, workers),
		wg:      sync.WaitGroup{},
	}
}

func (g *gTaskPool) Start() {
	for i := 0; i < g.workers; i++ {
		g.wg.Add(1)
		go func() {
			defer g.wg.Done()
			for task := range g.tasks {
				task.Run()
			}
		}()
	}
}

func (g *gTaskPool) Schedule(task ITask) {
	g.tasks <- task
}

func (g *gTaskPool) WaitAndStop() {
	close(g.tasks)
	g.wg.Wait()
}
