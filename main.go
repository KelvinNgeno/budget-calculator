package main

import "fmt"

func main() {
	//==========================================
	// 1. Ask for Income
	//==========================================
	var income float64
	fmt.Print("Enter your monthly income(KES): ")
	fmt.Scan(&income)

	//==========================================
	// 2. Ask for Expenses
	//==========================================
	var rent, food, transport, other float64
	fmt.Print("Enter rent/house expense: ")
	fmt.Scan(&rent)

	fmt.Print("Enter food expense: ")
	fmt.Scan(&food)

	fmt.Print("Enter transport expense: ")
	fmt.Scan(&transport)

	fmt.Print("Enter other expenses: ")
	fmt.Scan(&other)

	//=========================================
	// 3. Calculate totals
	//=========================================
	totalExpenses := rent + food + transport + other
	remaining := income - totalExpenses

	//=========================================
	// 4. Show the results
	//=========================================
	fmt.Println("\n==========BUDGET SUMMARY==========")
	fmt.Printf("Monthly Income: %.2f KES\n", income)
	fmt.Printf("Total Expenses: %.2f KES\n", totalExpenses)
	fmt.Printf("Money Remaining: %.2f KES\n", remaining)
	fmt.Println("===================================")

	//==========================================
	// 5. Give advice based on the result
	//==========================================
	if remaining > 0 {
		fmt.Printf("\nGood job! You can save %.2f KES this month.\n", remaining)
	} else if remaining == 0 {
		fmt.Println("\nYou broke even. No savings this month.")
	} else {
		fmt.Printf("\nWarning! You are over budget by %.2f KES.\n", -remaining)
	}
}

