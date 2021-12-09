package discount

import (
	"errors"
)

type AmountAndDiscount struct {
	Amount   int `json:"nbOfBooks"`
	Discount int `json:"discountPercentage"`
}

type Parameters struct {
	Cost               int                 `json:"discountableBookCost"`
	Discountable       []int               `json:"discountableBookIds"`
	AmountAndDiscounts []AmountAndDiscount `json:"discountScaling"`
}

func (p Parameters) Validate() error {
	if len(p.Discountable) == 0 {
		return errors.New("no item eligible to discount provided")
	}

	if len(p.AmountAndDiscounts) == 0 {
		return errors.New("no discount provided")
	}

	for _, elem := range p.AmountAndDiscounts {
		if elem.Amount <= 0 || elem.Discount > 100 || elem.Discount <= 0 {
			return errors.New("invalid discount")
		}
	}

	return nil
}

var TestParameters = Parameters{
	Cost:         1000,
	Discountable: []int{1, 2, 3},
	AmountAndDiscounts: []AmountAndDiscount{
		{
			Amount:   2,
			Discount: 25,
		},
		{
			Amount:   3,
			Discount: 50,
		},
	},
}
