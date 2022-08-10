package main

import "fmt"

func main() {

	var donations [9]int = [9]int{1000, 5000, 4000, 2000, 6000, 3000, 2200, 6000, 8000}
	totalDonation(donations)
}

func totalDonation(donations [9]int) {
	var requiredDonation = 25000
	var collectedMoney int = 0
	for i := 0; i < 9; i++ {

		var incomingDonation int = donations[i]
		if requiredDonation == collectedMoney {
			break
		} else if (requiredDonation - collectedMoney) >= incomingDonation {
			collectedMoney += incomingDonation
			fmt.Println(collectedMoney)
		} else {
			var moneyToBeAdded int = (requiredDonation - collectedMoney)
			var moneyToBeRefunded int = (donations[i] - moneyToBeAdded)
			fmt.Println(collectedMoney + moneyToBeAdded)
			fmt.Println(moneyToBeRefunded, "lira son bagisciya geri verildi")
			break
		}

	}
}
