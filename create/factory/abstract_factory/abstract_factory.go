package abstract_factory

import "log"

type (
	// 抽象产品
	ICanvas interface {
		Load()
		Draw()
	}

	ISvg interface {
		Init()
		Show()
	}

	// 抽象工厂
	ICanvasFactory interface {
		CreateCanvas() ICanvas
		CreateSvg() ISvg
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

	FactoryType int
)

const (
	FactoryTypeBarChart FactoryType = iota + 1
	FactoryTypeLineChart
)

// 工厂方法的方法
func NewCanvasFactory(ct FactoryType) ICanvasFactory {
	switch ct {
	case FactoryTypeBarChart:
		return newBarChartFactory()
	case FactoryTypeLineChart:
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
func (r barChartFactory) CreateCanvas() ICanvas {
	return &barChart{}
}

func (r barChartFactory) CreateSvg() ISvg {
	return &barChart{}
}

func (r lineChartFactory) CreateCanvas() ICanvas {
	return &lineChart{}
}

func (r lineChartFactory) CreateSvg() ISvg {
	return &lineChart{}
}

// barChart
func (y *barChart) Load() {
	log.Println("canvas: bar chart load")
}

func (y *barChart) Draw() {
	log.Println("canvas: bar chart draw")
}

func (y *barChart) Init() {
	log.Println("svg: bar chart init")
}

func (y *barChart) Show() {
	log.Println("svg: bar chart show")
}

// lineChart
func (y *lineChart) Load() {
	log.Println("canvas: line chart load")
}

func (y *lineChart) Draw() {
	log.Println("canvas: line chart Draw")
}

func (y *lineChart) Init() {
	log.Println("svg: line chart init")
}

func (y *lineChart) Show() {
	log.Println("svg: line chart show")
}
