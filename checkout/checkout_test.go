package checkout

import (
	"go-checkout-kata/product"
	"testing"
)

func TestAddItem(t *testing.T) {
	c := NewCheckout()
	item := product.Item{Sku: "A", Price: 100}

	err := c.ScanItem(item)

	if err != nil {
		t.Fatalf("Error adding item: %v", err)
	}
}

func TestEmptyStringSku(t *testing.T) {
	c := NewCheckout()
	item := product.Item{Sku: "", Price: 100}

	err := c.ScanItem(item)

	if err == nil {
		t.Fatalf("Expected error adding item")
	}
}

func TestAddMultipleItems(t *testing.T) {
	checkout := NewCheckout()
	item1 := product.Item{Sku: "A", Price: 100}
	item2 := product.Item{Sku: "B", Price: 200}
	item3 := product.Item{Sku: "C", Price: 300}

	err := checkout.ScanItem(item1)
	AssertError(t, err, "Error adding item1")

	err = checkout.ScanItem(item2)
	AssertError(t, err, "Error adding item2")

	err = checkout.ScanItem(item3)
	AssertError(t, err, "Error adding item3")

	if len(checkout.basket) != 3 {
		t.Fatalf("Expected 3 items in basket, got %d", len(checkout.basket))
	}
}

func TestTotalPriceOfItems(t *testing.T) {
	checkout := NewCheckout()
	item1 := product.Item{Sku: "A", Price: 5}
	item2 := product.Item{Sku: "B", Price: 5}
	item3 := product.Item{Sku: "C", Price: 5}

	err := checkout.ScanItem(item1)
	AssertError(t, err, "Error scanning item1")

	err = checkout.ScanItem(item2)
	AssertError(t, err, "Error scanning item2")

	err = checkout.ScanItem(item3)
	AssertError(t, err, "Error scanning item3")

	price := checkout.GetPrice()

	if price != 15 {
		t.Errorf("Expected price 15, got %d", price)
	}
}

func TestTotalPriceWithSingleDiscount(t *testing.T) {
	checkout := NewCheckout()
	item1 := product.Item{Sku: "A", Price: 5}
	item2 := product.Item{Sku: "B", Price: 10}
	item3 := product.Item{Sku: "A", Price: 5}
	item4 := product.Item{Sku: "A", Price: 5}
	item5 := product.Item{Sku: "B", Price: 10}

	// scan items
	err := checkout.ScanItem(item1)
	AssertError(t, err, "Error scanning item1")

	err = checkout.ScanItem(item2)
	AssertError(t, err, "Error scanning item2")

	err = checkout.ScanItem(item3)
	AssertError(t, err, "Error scanning item3")

	err = checkout.ScanItem(item4)
	AssertError(t, err, "Error scanning item4")

	err = checkout.ScanItem(item5)
	AssertError(t, err, "Error scanning item5")

	price := checkout.GetPrice()
	if price != 32 {
		t.Fatalf("Expected price 32, got %d", price)
	}
}

func TestTotalPriceWithMultipleDiscounts(t *testing.T) {

	checkout := NewCheckout()
	item1 := product.Item{Sku: "A", Price: 5}
	item2 := product.Item{Sku: "B", Price: 10}
	item3 := product.Item{Sku: "A", Price: 5}
	item4 := product.Item{Sku: "A", Price: 5}
	item5 := product.Item{Sku: "B", Price: 10}
	item6 := product.Item{Sku: "B", Price: 10}
	item7 := product.Item{Sku: "C", Price: 15}
	item8 := product.Item{Sku: "C", Price: 15}

	// scan items
	err := checkout.ScanItem(item1)
	AssertError(t, err, "Error scanning item1")

	err = checkout.ScanItem(item2)
	AssertError(t, err, "Error scanning item2")

	err = checkout.ScanItem(item3)
	AssertError(t, err, "Error scanning item3")

	err = checkout.ScanItem(item4)
	AssertError(t, err, "Error scanning item4")

	err = checkout.ScanItem(item5)
	AssertError(t, err, "Error scanning item5")

	err = checkout.ScanItem(item6)
	AssertError(t, err, "Error scanning item6")

	err = checkout.ScanItem(item7)
	AssertError(t, err, "Error scanning item7")

	err = checkout.ScanItem(item8)
	AssertError(t, err, "Error scanning item8")

	price := checkout.GetPrice()

	if price != 64 {
		t.Fatalf("Expected price 64, got %d", price)
	}

}

func AssertError(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf("%s: %s", message, err)
	}
}
