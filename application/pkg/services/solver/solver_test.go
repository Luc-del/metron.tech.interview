package solver

import (
	"github.com/stretchr/testify/assert"
	"interview/application/pkg/models/combination"
	"testing"
)

func Test_RemoveSeveralPendingZeroDiscount(t *testing.T) {
	t.Run("nothing to do", func(t *testing.T) {
		c := combination.TestSetCostNoDuplicate()
		removeSeveralPendingZeroDiscount(&c)
		assert.Equal(t, combination.TestSetCostNoDuplicate(), c)
	})

	t.Run("remove last", func(t *testing.T) {
		c := combination.TestSetCost()
		removeSeveralPendingZeroDiscount(&c)
		assert.Equal(t, combination.TestSetCostNoDuplicate(), c)
	})

	t.Run("remove several", func(t *testing.T) {
		c := combination.TestSetCostTwiceDuplicated()
		removeSeveralPendingZeroDiscount(&c)
		assert.Equal(t, []int{1, 1, 1}, c.Combinations[len(c.Combinations)-1].Items)
	})
}
