package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

	/*
		finalCost = bulkMessageCost
		if isPremiumUser {
			finalCost -= (finalCost * discountRate)
		}
	*/
	// GoLang should allow for ternary operators "(conditional) ? true : false", closest aproximation
	// output = (map[bool]int{true: trueResult, false: falseResult})[condtion]
	finalCost = (map[bool]float64{true: bulkMessageCost - (bulkMessageCost * discountRate), false: bulkMessageCost})[isPremiumUser]
	if accountBalance >= finalCost {
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}
