package main

import (
	"errors"
	"go-checkout-kata/pricingscheme"
)

type Checkout struct {
	basket []Item
}

const errorEmptySKU = "Empty SKU"

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) scanItem(item Item) error {
	if item.sku == "" {
		return errors.New("Empty SKU")
	}
	c.basket = append(c.basket, item)
	return nil
}

func (c *Checkout) getPrice(pricingScheme pricingscheme.PricingScheme) int {
	var price int
	for _, item := range c.basket {
		price += item.price
	}

	hasDiscount, discount := checkDiscount(c.basket, pricingScheme)
	if hasDiscount {
		return price - discount
	}

	return price
}

func checkDiscount(basket []Item, pricingScheme pricingscheme.PricingScheme) (bool, int) {
	// create a map to count the number of items per sku
	itemMap := make(map[string]int)

	for _, item := range basket {
		itemMap[item.sku]++
	}

	hasDiscount := false
	for key, value := range itemMap {
		// check if the sku and quantity matches the pricing scheme
		if key == pricingScheme.Sku && value == pricingScheme.Quantity {
			hasDiscount = true
			break
		}
	}

	// if there is a discount, calculate the total price of the item and return the difference between the total price and the discounted price
	if hasDiscount {
		it := Item{}
		for _, item := range basket {
			if item.sku == pricingScheme.Sku {
				it = item
			}
		}

		totalPriceOfItem := it.price * pricingScheme.Quantity
		return true, totalPriceOfItem - pricingScheme.DiscountedPrice
	}

	return false, 0
}
