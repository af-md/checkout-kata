package main

import "testing"

func TestAddItem(t *testing.T) {
	c := NewCheckout()
	err := c.scanItem("A")

	if err != nil {
		t.Fatalf("Error adding item: %v", err)
	}
}
