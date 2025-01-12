/*
	Switch Case
	- in multiple languages -> for every case, we need a break
	- in Go, we don't need
		-> just need return at default case


*/

package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	for {
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Println("Your balance is", accountBalance)
			case 2:
				fmt.Print("Your deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				accountBalance += depositAmount 
				fmt.Println("Balance updated! New amount:", accountBalance)
			case 3:
				fmt.Print("Withdrawal amount: ")
				var withdrawalAmount float64
				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					continue
				}

				accountBalance -= withdrawalAmount 
				fmt.Println("Balance updated! New amount:", accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for choosing our bank")
				return // no break -> since we want to get out of the switch statement
			}
	}
}
