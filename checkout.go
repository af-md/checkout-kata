package main

import "errors"

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

func (c *Checkout) getPrice() int {
	var price int
	for _, item := range c.basket {
		price += item.price
	}
	return price
}
