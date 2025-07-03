package utils

import (
	"fmt"
	"testing"
)

func TestGTaskPool(t *testing.T) {
	goPool := NewGTaskPool(4)
	goPool.Start()
	for i := 0; i < 10; i++ {
		goPool.Schedule(&tTack{num: i})
	}
	goPool.WaitAndStop()
}

type tTack struct {
	num int
}

func (t *tTack) Run() {
	fmt.Println(t.num)
}
