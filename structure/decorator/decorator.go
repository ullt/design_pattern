package decorator

import "fmt"

type IDecorate interface {
	Do()
}

type Decorate struct {
	decorate IDecorate
}

func NewDecorateFunc(decorate IDecorate) IDecorate {
	return &Decorate{
		decorate: decorate,
	}
}

func (d *Decorate) Do() {
	fmt.Println("进入装饰方法")
	d.decorate.Do()
	fmt.Println("结束装饰方法")
}

type DecorateA struct{}

func (d *DecorateA) Do() {
	fmt.Println("原始方法")
}
