package main

import "testing"

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

	if err != nil {
		t.Errorf("Error adding item1: %s", err)
	}

	err = checkout.scanItem(item2)

	if err != nil {
		t.Errorf("Error adding item2: %s", err)
	}

	err = checkout.scanItem(item3)

	if err != nil {
		t.Errorf("Error adding item3: %s", err)
	}

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
	err = checkout.scanItem(item2)
	err = checkout.scanItem(item3)

	if err != nil {
		t.Errorf("Error scanning items: %s", err)
	}

	price := checkout.getPrice()

	if price != 15 {
		t.Errorf("Expected price 15, got %d", price)
	}
}
