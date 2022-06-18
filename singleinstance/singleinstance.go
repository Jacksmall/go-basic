// 单例模式
// 通过增加一个数字型的标志位来降低互斥锁的使用次数来提高性能
// sync.Once 就是使用一次 单例
package singleinstance

import (
	"sync"
	"sync/atomic"
)

type singlaton struct{}

var (
	instance    *singlaton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singlaton {
	// 加载标识位是否设置
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 加原子锁操作
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		// 设置标识位
		defer atomic.StoreUint32(&initialized, 1)
		instance = new(singlaton)
	}
	return instance
}
