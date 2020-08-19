package factory_method

import "log"

type (
	// 抽象产品
	ICanvas interface {
		Load()
		Draw()
	}

	// 抽象工厂
	ICanvasFactory interface {
		Create() ICanvas
	}

	// 具体工厂
	barChartFactory struct {
	}

	lineChartFactory struct {
	}

	// 具体产品
	barChart struct {
	}

	lineChart struct {
	}

	CanvasFactoryType int
)

const (
	CanvasFactoryTypeBarChart CanvasFactoryType = iota + 1
	CanvasFactoryTypeLineChart
)

// 工厂方法的方法
func NewCanvasFactory(ct CanvasFactoryType) ICanvasFactory {
	switch ct {
	case CanvasFactoryTypeBarChart:
		return newBarChartFactory()
	case CanvasFactoryTypeLineChart:
		return newLineChartFactory()
	default:
		return nil
	}
}

// 工厂方法
func newBarChartFactory() ICanvasFactory {
	return &barChartFactory{}
}

func newLineChartFactory() ICanvasFactory {
	return &lineChartFactory{}
}

// interface implementation
func (r barChartFactory) Create() ICanvas {
	return &barChart{}
}

func (y *barChart) Load() {
	log.Println("bar chart load")
}

func (y *barChart) Draw() {
	log.Println("bar chart draw")
}

func (r lineChartFactory) Create() ICanvas {
	return &lineChart{}
}

func (y *lineChart) Load() {
	log.Println("line chart load")
}

func (y *lineChart) Draw() {
	log.Println("line chart draw")
}
