package prototype

type (
	ICloneable interface {
		Clone() ICloneable
	}

	CarCenter struct {
		carMap map[CarType]ICloneable
	}

	CarType int

	BMWCar struct{
		Name string
	}
	LRCar  struct{
		Name string
	}
)

const (
	CarTypeBMWCar CarType = iota + 1
	CarTypeLRCar
)

func NewCarCenter() *CarCenter {
	return &CarCenter{
		carMap: make(map[CarType]ICloneable),
	}
}

func (c *CarCenter) Get(ct CarType) ICloneable {
	return c.carMap[ct]
}

func (c *CarCenter) Set(ct CarType, car ICloneable) {
	c.carMap[ct] = car
}

func (c *BMWCar) Clone() ICloneable {
	car := *c
	return &car
}

func (c *LRCar) Clone() ICloneable {
	car := *c
	return &car
}
