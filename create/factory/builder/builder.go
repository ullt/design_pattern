package builder

import (
	"log"
)

type (
	// 抽象餐品
	IItem interface {
		Name() string
		Price() int // 精确到分
	}

	// 具体餐品
	Rice struct {
		name  string
		price int
	}

	Vegetable struct {
		name  string
		price int
	}

	// 抽象订单
	IOrder interface {
		ShowCost()
		ShowItem()
	}

	// 具体订单
	Order struct {
		UserId   uint64
		itemList []IItem
	}

	// 订单构建类
	OrderBuilder struct {
		order *Order
	}
)

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{
		order: &Order{
			UserId:   0,
			itemList: []IItem{},
		},
	}
}

func (ob *OrderBuilder) SetItem(item IItem) *OrderBuilder {
	// verify item
	ob.order.itemList = append(ob.order.itemList, item)
	return ob
}

func (ob *OrderBuilder) SetUserId(userId uint64) *OrderBuilder {
	// verify parameter
	ob.order.UserId = userId
	return ob
}

func (ob *OrderBuilder) Build() IOrder {
	// verify order required and legitimacy
	return ob.order
}

func (r *Rice) Name() string {
	return "rice"
}
func (r *Rice) Price() int {
	return 1
}

func (v *Vegetable) Name() string {
	return "vegetable"
}
func (v *Vegetable) Price() int {
	return 5
}

func (o *Order) ShowCost() {
	total := 0
	for _, item := range o.itemList {
		total += item.Price()
	}
	log.Printf("total cost: %d", total)
}
func (o *Order) ShowItem() {
	for _, item := range o.itemList {
		log.Printf(">>> order item: %s", item.Name())
	}
}
