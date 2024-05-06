package main

import "testing"

func TestAddItem(t *testing.T) {
	c := NewCheckout()
	err := c.scanItem("A")

	if err != nil {
		t.Fatalf("Error adding item: %v", err)
	}
}

func TestEmptyStringSku(t *testing.T) {
	c := NewCheckout()
	err := c.scanItem("")

	if err == nil {
		t.Fatalf("Expected error adding item")
	}
}

func TestAddMultipleItems(t *testing.T) {
	checkout := NewCheckout()

	err := checkout.scanItem("sku1")

	if err != nil {
		t.Errorf("Error scanning item sku1: %s", err)
	}

	err = checkout.scanItem("sku2")

	if err != nil {
		t.Errorf("Error scanning item sku2: %s", err)
	}

	err = checkout.scanItem("sku3")

	if err != nil {
		t.Errorf("Error scanning item sku3: %s", err)
	}

	if len(checkout.basket) != 3 {
		t.Fatalf("Expected 3 items in basket, got %d", len(checkout.basket))
	}
}
