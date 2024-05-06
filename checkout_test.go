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
