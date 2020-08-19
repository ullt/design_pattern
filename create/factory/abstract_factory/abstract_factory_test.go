package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	lineChartFactory := NewCanvasFactory(FactoryTypeLineChart)
	barChartFactory := NewCanvasFactory(FactoryTypeBarChart)

	lcCanvas := lineChartFactory.CreateCanvas()
	lcCanvas.Load()
	lcCanvas.Draw()

	lcSvg := lineChartFactory.CreateSvg()
	lcSvg.Init()
	lcSvg.Show()

	bcCanvas := barChartFactory.CreateCanvas()
	bcCanvas.Load()
	bcCanvas.Draw()

	bcSvg := barChartFactory.CreateSvg()
	bcSvg.Init()
	bcSvg.Show()
}
