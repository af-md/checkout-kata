package main

import (
	"fmt"
	"go-checkout-kata/checkout"
	"go-checkout-kata/product"
	"log"
)

func main() {
	// Single discount multiple items
	checkout := checkout.NewCheckout()
	items := createSingleDiscountItems()
	for _, item := range items {
		err := checkout.ScanItem(item)
		if err != nil {
			log.Fatalf("Error scanning item: %s", err)
		}
	}

	price := checkout.GetPrice()

	fmt.Printf("Total price: %d\n", price)

	// Multiple discount multiple items
	items = createMultipleDiscountItems()

	for _, item := range items {
		err := checkout.ScanItem(item)
		if err != nil {
			log.Fatalf("Error scanning item: %s", err)
		}
	}

	price = checkout.GetPrice()

	fmt.Printf("Total price: %d\n", price)

}

func createSingleDiscountItems() []product.Item {
	item1 := product.Item{Sku: "A", Price: 5}
	item2 := product.Item{Sku: "B", Price: 10}
	item3 := product.Item{Sku: "A", Price: 5}
	item4 := product.Item{Sku: "A", Price: 5}
	return []product.Item{item1, item2, item3, item4}
}

func createMultipleDiscountItems() []product.Item {
	items := createSingleDiscountItems()

	items = append(items, product.Item{Sku: "B", Price: 10})
	items = append(items, product.Item{Sku: "B", Price: 10})
	items = append(items, product.Item{Sku: "C", Price: 15})

	return items
}
