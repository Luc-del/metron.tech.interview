package solver

import (
	"interview/application/pkg/models/combination"
	"interview/application/pkg/models/discount"
)

//func bruteForce(discounts discount.Discount, items ...int) combination.PriceWithSets {
//	//maximum price if no item is items
//	lowestPriceCombination := combination.PriceWithSets{
//		Price:        discounts.Cost() * len(items),
//	}
//	// find lowest combination price for items items
//	for  comb := range makeCombinations(items...) {
//		currentCombination := comb.ToDiscountableCombinationPrice(discounts)
//		if currentCombination.Price < lowestPriceCombination.Price {
//			lowestPriceCombination = currentCombination
//		}
//	}
//
//	return lowestPriceCombination
//}

//Only works if the length of different items doesn't exceed max discount amount and with no holes
func SimpleSolve(discounts discount.Discount, items ...int) combination.PriceWithSets {
	sets := combination.NewOrderedCombination(items...)
	prices := sets.ToDiscountableSets(discounts)

	//Pending sets may be duplicates or without reduction and can be merged
	removeSeveralPendingZeroDiscount(&prices)
	return prices
}

func removeSeveralPendingZeroDiscount(prices *combination.PriceWithSets) {
	if len(prices.Combinations) > 1 {
		//last can be 0 discount
		for idx := len(prices.Combinations) - 2; prices.Combinations[idx].Discount == 0; idx = len(prices.Combinations) - 2 {

			previousLastItems := &prices.Combinations[idx].Items
			*previousLastItems = append(*previousLastItems, prices.Combinations[idx+1].Items...)

			prices.Combinations[idx].Price += prices.Combinations[idx+1].Price

			prices.Combinations = prices.Combinations[:len(prices.Combinations)-1]
		}
	}
}

func isCombinationOptimized(discounts discount.Discount, c combination.Combination) bool {
	for _, set := range c {
		if len(set) <= discounts.GetLowestAmountForDiscount() {
			return true
		}
		if discounts.DiscountOnUniqueAndDiscountableItems(set...) == 0 {
			return false
		}
	}

	return true
}

//func rebalance(sets combination.Combination) []combination.Combination {
//
//}
