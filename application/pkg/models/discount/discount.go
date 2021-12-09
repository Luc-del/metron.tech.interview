package discount

import (
	"sync"
)

var (
	//global variables
	d Discount
	m sync.Mutex
)

// Discount holds the discount values
// cost is the value of the items
// discountable holds the items that are discounted
// values holds the discount to apply corresponding to the number of items
type Discount struct {
	cost         int
	discountable map[int]struct{}
	values       map[int]int
}

func (d Discount) IsEmpty() bool {
	return len(d.discountable) == 0 || len(d.values) == 0
}

func (d Discount) DiscountSize() int {
	return len(d.values)
}

// Discount returns the discount to applied to a set of items
// items have to be discountable and unique
func (d Discount) Discount(items ...int) int {
	discountableCount := 0
	unique := make(map[int]struct{}, len(items))
	for _, item := range items {
		_, isUnique := unique[item]
		_, isDiscountable := d.discountable[item]
		if isDiscountable && isUnique {
			discountableCount++
		}
		unique[item] = struct{}{}
	}

	return d.values[discountableCount]
}

func (d Discount) DiscountOnUniqueAndDiscountableItems(items ...int) int {
	return d.values[len(items)]
}

func (d Discount) GetLowestAmountForDiscount() int {
	if d.IsEmpty() {
		panic("empty discount")
	}
	for amount := 0; ; amount++ {
		v, ok := d.values[amount]
		if ok {
			return v
		}
	}
}

func (d Discount) Cost() int {
	return d.cost
}

func (d Discount) IsDiscountable(item int) bool {
	_, ok := d.discountable[item]
	return ok
}

func GetDiscounts() Discount {
	m.Lock()
	defer m.Unlock()

	discountable := make(map[int]struct{}, len(d.discountable))
	for k := range d.discountable {
		discountable[k] = struct{}{}
	}
	discounts := make(map[int]int, len(d.values))
	for k, v := range d.values {
		discounts[k] = v
	}

	return Discount{
		cost:         d.cost,
		discountable: discountable,
		values:       discounts,
	}
}

// Set the discounts
func Set(p Parameters) {
	discountableMap := make(map[int]struct{}, len(p.Discountable))
	for _, v := range p.Discountable {
		discountableMap[v] = struct{}{}
	}

	discounts := make(map[int]int, len(p.AmountAndDiscounts))
	for _, item := range p.AmountAndDiscounts {
		discounts[item.Amount] = item.Discount
	}

	m.Lock()
	defer m.Unlock()
	d.cost = p.Cost
	d.discountable = discountableMap
	d.values = discounts
}
