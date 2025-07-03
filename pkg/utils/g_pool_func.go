package utils

import (
	"gin-scaffold/pkg/logger"
	"sync"
)

type IGoFuncPool interface {
	Start()
	Schedule(task func() error)
	WaitAndStop()
}

type gFuncPool struct {
	workers int
	tasks   chan func() error
	wg      sync.WaitGroup
}

func NewGFuncPool(workers int) IGoFuncPool {
	return &gFuncPool{
		workers: workers,
		tasks:   make(chan func() error, workers),
		wg:      sync.WaitGroup{},
	}
}

func (g *gFuncPool) Start() {
	for i := 0; i < g.workers; i++ {
		g.wg.Add(1)
		go func() {
			defer g.wg.Done()
			for task := range g.tasks {
				err := task()
				if err != nil {
					logger.Logger.Error(err.Error())
					continue
				}
			}
		}()
	}
}

func (g *gFuncPool) Schedule(task func() error) {
	g.tasks <- task
}

func (g *gFuncPool) WaitAndStop() {
	close(g.tasks)
	g.wg.Wait()
}
