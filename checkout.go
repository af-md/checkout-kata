package main

import (
	"errors"
)

type Checkout struct {
	skus []string
}

const errorEmptySKU = "Empty SKU"

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) scanItem(sku string) error {

	if sku == "" {
		return errors.New(errorEmptySKU)
	}

	c.skus = append(c.skus, sku)
	return nil
}
