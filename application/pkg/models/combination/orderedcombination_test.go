package combination

import (
	"github.com/stretchr/testify/assert"
	"interview/application/pkg/models/discount"
	"testing"
)

func Test_OrderedCombination(t *testing.T) {
	t.Run("single item", func(t *testing.T) {
		assert.Equal(t, Combination{[]int{1}}, NewOrderedCombination(1).Combination)
	})

	t.Run("several different items", func(t *testing.T) {
		assert.Equal(t, Combination{[]int{1, 2, 3}}, NewOrderedCombination(1, 2, 3).Combination)
	})

	t.Run("redondant items", func(t *testing.T) {
		assert.Equal(t, Combination{[]int{1, 2, 3}, {1, 3}}, NewOrderedCombination(1, 2, 1, 3, 3).Combination)
	})
}

func TestToDiscountableSets(t *testing.T) {
	discount.Set(discount.TestParameters)

	c := TestOrderedCombination()
	assert.Equal(t, TestSetCost(), c.ToDiscountableSets(discount.GetDiscounts()))
}
