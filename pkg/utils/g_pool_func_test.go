package utils

import (
	"fmt"
	"testing"
)

func TestGFuncPool(t *testing.T) {
	goPool := NewGFuncPool(4)
	goPool.Start()
	for i := 0; i < 10; i++ {
		goPool.Schedule(func() error {
			printNum(i)
			return nil
		})
	}
	goPool.WaitAndStop()
}

func printNum(num int) {
	fmt.Printf("%d ", num)
}
