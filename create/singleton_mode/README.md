# 单例模式

## 定义
>一个类能且只能被允许创建一个实例，就可以成为单例，这种设计模式就可以被称为单例设计模式

## 单例的用途
* 有些数据只应该保存一份，那么就适合单例模式，如：
    * 全局唯一ID生成器：防止ID资源重复
    * 配置文件读取：配置文件被加载后理应只有一份数据
    * 长连接客户端：一个server实例化一个客户端，防止资源浪费

## 实现
* 懒惰模式加双重检查锁
```go
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
```