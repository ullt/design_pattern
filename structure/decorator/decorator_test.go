package decorator

func Example() {
	var a IDecorate = &DecorateA{}
	a = NewDecorateFunc(a)
	a.Do()
	// Output:
	//   进入装饰方法
	//   原始方法
	//   结束装饰方法
}
