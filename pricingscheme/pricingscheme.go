package pricingscheme

type PricingScheme struct {
	Sku             string
	Quantity        int
	DiscountedPrice int
}

func GetPricingScheme() []PricingScheme {
	// this function is used to get a fresh pricing scheme from the model
	// can be replaced to a call to a database or an API

	return []PricingScheme{{
		Sku:             "A",
		Quantity:        3,
		DiscountedPrice: 12,
	}, {
		Sku:             "B",
		Quantity:        3,
		DiscountedPrice: 22,
	}}
}
