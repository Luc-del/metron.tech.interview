package cashier

import (
	"interview/application/pkg/models/combination"
	"interview/application/pkg/models/discount"
	"interview/application/pkg/services/solver"
	"math"
)

func ComputeCost(discounts discount.Discount, items ...int) combination.PriceWithSets {
	discountable, notDiscountable := splitDiscountable(discounts, items...)

	res := combination.PriceWithSets{}
	if !discounts.IsEmpty() {
		res = solver.SimpleSolve(discounts, discountable...)
	}

	//make a distinct category on the receipt for non discountable items
	NotDiscountableSet := combination.SetCost{
		Items:    notDiscountable,
		Discount: 0,
		Price:    len(notDiscountable) * discounts.Cost(),
	}
	res.Combinations = append(res.Combinations, NotDiscountableSet)
	res.Price += NotDiscountableSet.Price

	return res
}

func splitDiscountable(d discount.Discount, items ...int) ([]int, []int) {
	l := int(math.Ceil(float64(len(items) / 2)))
	isDiscountable := make(map[bool][]int, 2)
	isDiscountable[true] = make([]int, 0, l)
	isDiscountable[false] = make([]int, 0, l)

	for _, item := range items {
		isDiscountable[d.IsDiscountable(item)] = append(isDiscountable[d.IsDiscountable(item)], item)
	}

	return isDiscountable[true], isDiscountable[false]
}
