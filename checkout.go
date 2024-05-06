package main

import "errors"

type Checkout struct {
	basket []string
}

const errorEmptySKU = "Empty SKU"

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) scanItem(sku string) error {
	if sku == "" {
		return errors.New(errorEmptySKU)
	}
	c.basket = append(c.basket, sku)
	return nil
}
