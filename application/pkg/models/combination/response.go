package combination

type PriceWithSets struct {
	Price        int       `json:"totalCost"`
	Combinations []SetCost `json:"sets"`
}

type SetCost struct {
	Items    []int `json:"ids"`
	Discount int   `json:"discountPercentage"`
	Price    int   `json:"setTotalCost"`
}
