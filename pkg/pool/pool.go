package pool

import (
	"github.com/panjf2000/ants/v2"
	"sync"
)

var pool *ants.PoolWithFunc
var wg sync.WaitGroup

func init() {
	defer ants.Release()
	p, err := ants.NewPoolWithFunc(20, func(i interface{}) {
		task := i.(func()) // 将传入的任务转换成函数
		task()             // 执行任务函数
	})
	if err != nil {
		panic(err)
	}
	pool = p
}

func SubmitFunc(task func()) error {
	wg.Add(1)
	err := pool.Invoke(task)
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}
