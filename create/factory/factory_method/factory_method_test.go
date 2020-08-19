package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	barCF := NewCanvasFactory(CanvasFactoryTypeBarChart)
	lineCF := NewCanvasFactory(CanvasFactoryTypeLineChart)

	bar := barCF.Create()
	bar.Load()
	bar.Draw()

	line := lineCF.Create()
	line.Load()
	line.Draw()
}
