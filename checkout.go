package main

type Checkout struct {
	skus []string
}

func NewCheckout() *Checkout {
	return &Checkout{}
}

func (c *Checkout) scanItem(sku string) error {
	c.skus = append(c.skus, sku)
	return nil
}
