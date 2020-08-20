package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	// 用户1点了一份饭一份菜
	ob := NewOrderBuilder()
	ob.SetItem(new(Rice))
	ob.SetItem(new(Vegetable))
	ob.SetUserId(1)
	order1 := ob.Build()
	order1.ShowCost()
	order1.ShowItem()

	// 用户2点了一份饭三份菜
	ob2 := NewOrderBuilder()
	ob2.SetItem(new(Rice))
	ob2.SetItem(new(Vegetable))
	ob2.SetItem(new(Vegetable))
	ob2.SetItem(new(Vegetable))
	ob2.SetUserId(2)
	order2 := ob2.Build()
	order2.ShowCost()
	order2.ShowItem()
}
