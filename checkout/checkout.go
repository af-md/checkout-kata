package checkout

import (
	"errors"
	"go-checkout-kata/pricingscheme"
	"go-checkout-kata/product"
)

type Checkout struct {
	basket []product.Item
}

const errorEmptySKU = "Empty SKU"

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) ScanItem(item product.Item) error {
	if item.Sku == "" {
		return errors.New("Empty SKU")
	}
	c.basket = append(c.basket, item)
	return nil
}

func (c *Checkout) GetPrice() int {
	var price int
	for _, item := range c.basket {
		price += item.Price
	}

	discount := c.calculateDiscount(c.basket)

	return price - discount
}

func (c *Checkout) calculateDiscount(basket []product.Item) int {
	itemCM := make(map[string]int)

	for _, item := range basket {
		itemCM[item.Sku]++
	}

	pricingSchemes := pricingscheme.GetPricingScheme()

	totalPriceDiff := 0
	hasDiscount := false

	// loop through pricing schemes
	for _, pricingScheme := range pricingSchemes {
		for key, value := range itemCM {
			if key == pricingScheme.Sku && value == pricingScheme.Quantity {
				// sku in the basket matches the pricing scheme set flag
				hasDiscount = true
				break
			}
		}

		// if discount is applicable for the pricing scheme then calculate the total price difference
		if hasDiscount {
			item1 := product.Item{}
			for _, item := range basket {
				if item.Sku == pricingScheme.Sku {
					item1 = item
				}

			}

			totalPriceOfItems := item1.Price * pricingScheme.Quantity
			totalPriceDiff += totalPriceOfItems - pricingScheme.DiscountedPrice
			// reset flag for next pricing scheme
			hasDiscount = false
		}
	}

	return totalPriceDiff
}
