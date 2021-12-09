package combination

func TestOrderedCombination() OrderedUniqueCombination {
	return OrderedUniqueCombination{Combination{[]int{1, 2, 3}, []int{1, 3}, []int{1}, []int{1}}}
}

func TestSetCost() PriceWithSets {
	return PriceWithSets{
		Price: (15 + 15 + 10 + 10) * 100,
		Combinations: []SetCost{
			{
				Items:    []int{1, 2, 3},
				Discount: 50,
				Price:    1500,
			},
			{
				Items:    []int{1, 3},
				Discount: 25,
				Price:    1500,
			},
			{
				Items:    []int{1},
				Discount: 0,
				Price:    1000,
			},
			{
				Items:    []int{1},
				Discount: 0,
				Price:    1000,
			},
		},
	}
}

func TestSetCostTwiceDuplicated() PriceWithSets {
	return PriceWithSets{
		Price: (15 + 15 + 10 + 10 + 10) * 100,
		Combinations: []SetCost{
			{
				Items:    []int{1, 2, 3},
				Discount: 50,
				Price:    1500,
			},
			{
				Items:    []int{1, 3},
				Discount: 25,
				Price:    1500,
			},
			{
				Items:    []int{1},
				Discount: 0,
				Price:    1000,
			},
			{
				Items:    []int{1},
				Discount: 0,
				Price:    1000,
			},
			{
				Items:    []int{1},
				Discount: 0,
				Price:    1000,
			},
		},
	}
}

func TestSetCostNoDuplicate() PriceWithSets {
	return PriceWithSets{
		Price: (15 + 15 + 20) * 100,
		Combinations: []SetCost{
			{
				Items:    []int{1, 2, 3},
				Discount: 50,
				Price:    1500,
			},
			{
				Items:    []int{1, 3},
				Discount: 25,
				Price:    1500,
			},
			{
				Items:    []int{1, 1},
				Discount: 0,
				Price:    2000,
			},
		},
	}
}
