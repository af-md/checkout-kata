# Checkout kata

## Run

The main file contains runs of the checkout with different scenarios.

To run the program, follow these steps:
```
go run main.go
```

## Test
To run the tests, follow these steps:
```
cd checkout/
go test -v
```

## Assumptions made
- I assumed that the item with discounts quantity scanned at the checkout, will always match the pricing scheme quantity i.e. the checkout will only get 3 A items instead 4 A items for all my tests. This was done to limit my problem space and make the problem manageable in the recommended time frame.

## Future improvements
- Add more tests to cover more edge cases
- Support for more complex offers i.e. buy 1 get 1 for free
- Add an API and database layer
- Integration with inventory management to update stock level