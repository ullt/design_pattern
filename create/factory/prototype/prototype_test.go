package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrototype(t *testing.T) {
	as := assert.New(t)
	cc := NewCarCenter()

	cc.Set(CarTypeBMWCar, &BMWCar{Name: "BMW"})
	cc.Set(CarTypeLRCar, &LRCar{Name: "LR"})

	lr := cc.Get(CarTypeLRCar)
	lr1 := lr.Clone()
	as.NotNil(lr)
	as.NotNil(lr1)
	as.NotSame(lr, lr1)

	bmw := cc.Get(CarTypeBMWCar)
	bmw1 := bmw.Clone()
	as.NotNil(bmw)
	as.NotNil(bmw1)
	as.NotSame(bmw, bmw1)
}
