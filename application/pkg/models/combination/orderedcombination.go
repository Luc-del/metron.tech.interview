package combination

import "interview/application/pkg/models/discount"

type OrderedUniqueCombination struct {
	Combination `json:",inline"`
}

func NewOrderedCombination(items ...int) OrderedUniqueCombination {
	res := OrderedUniqueCombination{}
	for _, item := range items {
		res.addIfUniqueOrPropagate(0, item)
	}

	return res
}

func (c *OrderedUniqueCombination) addIfUniqueOrPropagate(idx int, elem int) {
	if idx == len(c.Combination) {
		c.Combination = append(c.Combination, []int{elem})
		return
	}

	if !contains(c.Combination[idx], elem) {
		c.Combination[idx] = append(c.Combination[idx], elem)
		return
	}

	c.addIfUniqueOrPropagate(idx+1, elem)
}

func contains(t []int, elem int) bool {
	for _, k := range t {
		if k == elem {
			return true
		}
	}

	return false
}

// Only valid for discountable Combination of unique items
func (c OrderedUniqueCombination) ToDiscountableSets(d discount.Discount) PriceWithSets {
	totalPrice := 0

	sets := make([]SetCost, 0, len(c.Combination))
	for _, set := range c.Combination {
		setDiscount := d.DiscountOnUniqueAndDiscountableItems(set...)
		sc := SetCost{
			Items:    set,
			Discount: setDiscount,
			Price:    int(float64(100-setDiscount) / 100 * float64(len(set)*d.Cost())),
		}
		sets = append(sets, sc)

		totalPrice += sc.Price
	}

	return PriceWithSets{
		Price:        totalPrice,
		Combinations: sets,
	}
}
