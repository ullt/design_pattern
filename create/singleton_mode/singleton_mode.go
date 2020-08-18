package singleton_mode

import "sync"

type (
	// 单例类应实现的接口
	ISingleton interface {
	}

	// 单例类（单例结构体：以面向对象分析来说，单例类更习惯一些）
	singleton struct {
	}
)

var (
	single ISingleton
	once   sync.Once
)

func initInstance() {
	once.Do(func() {
		single = new(singleton)
	})
}

func GetSingle() ISingleton {
	if single == nil {
		initInstance()
	}
	return single
}
