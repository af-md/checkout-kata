package main

import (
	"go-checkout-kata/pricingscheme"
	"testing"
)

func TestAddItem(t *testing.T) {
	c := NewCheckout()
	item := Item{sku: "A", price: 100}

	err := c.scanItem(item)

	if err != nil {
		t.Fatalf("Error adding item: %v", err)
	}
}

func TestEmptyStringSku(t *testing.T) {
	c := NewCheckout()
	item := Item{sku: "", price: 100}

	err := c.scanItem(item)

	if err == nil {
		t.Fatalf("Expected error adding item")
	}
}

func TestAddMultipleItems(t *testing.T) {
	checkout := NewCheckout()
	item1 := Item{sku: "A", price: 100}
	item2 := Item{sku: "B", price: 200}
	item3 := Item{sku: "C", price: 300}

	err := checkout.scanItem(item1)
	AssertError(t, err, "Error adding item1")

	err = checkout.scanItem(item2)
	AssertError(t, err, "Error adding item2")

	err = checkout.scanItem(item3)
	AssertError(t, err, "Error adding item3")

	if len(checkout.basket) != 3 {
		t.Fatalf("Expected 3 items in basket, got %d", len(checkout.basket))
	}
}

func TestTotalPriceOfItems(t *testing.T) {
	checkout := NewCheckout()
	item1 := Item{sku: "A", price: 5}
	item2 := Item{sku: "B", price: 5}
	item3 := Item{sku: "C", price: 5}

	err := checkout.scanItem(item1)
	AssertError(t, err, "Error scanning item1")

	err = checkout.scanItem(item2)
	AssertError(t, err, "Error scanning item2")

	err = checkout.scanItem(item3)
	AssertError(t, err, "Error scanning item3")

	price := checkout.getPrice(pricingscheme.PricingScheme{})

	if price != 15 {
		t.Errorf("Expected price 15, got %d", price)
	}
}

func TestTotalPriceWithSingleDiscount(t *testing.T) {
	checkout := NewCheckout()
	item1 := Item{sku: "A", price: 5}
	item2 := Item{sku: "B", price: 10}
	item3 := Item{sku: "A", price: 5}
	item4 := Item{sku: "A", price: 5}
	item5 := Item{sku: "B", price: 10}

	pricingScheme := pricingscheme.PricingScheme{Sku: "A", Quantity: 3, DiscountedPrice: 12}

	// scan items
	err := checkout.scanItem(item1)
	AssertError(t, err, "Error scanning item1")

	err = checkout.scanItem(item2)
	AssertError(t, err, "Error scanning item2")

	err = checkout.scanItem(item3)
	AssertError(t, err, "Error scanning item3")

	err = checkout.scanItem(item4)
	AssertError(t, err, "Error scanning item4")

	err = checkout.scanItem(item5)
	AssertError(t, err, "Error scanning item5")

	price := checkout.getPrice(pricingScheme)
	if price != 32 {
		t.Fatalf("Expected price 32, got %d", price)
	}
}

func AssertError(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("%s: %s", message, err)
	}
}
