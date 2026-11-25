package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// readAmount asks until the user enters a valid non-negative number.
func readAmount(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		text, err := reader.ReadString('\n')
		if err != nil && text == "" {
			fmt.Println("\nInput closed. Exiting.")
			os.Exit(1)
		}

		text = strings.ReplaceAll(strings.TrimSpace(text), ",", "")
		value, parseErr := strconv.ParseFloat(text, 64)
		if parseErr != nil || value < 0 {
			fmt.Println("Please enter a valid number (0 or more).")
			continue
		}
		return value
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	//==========================================
	// 1. Ask for Income
	//==========================================
	income := readAmount(reader, "Enter your monthly income (KES): ")

	//==========================================
	// 2. Ask for Expenses
	//==========================================
	rent := readAmount(reader, "Enter rent/house expense: ")
	food := readAmount(reader, "Enter food expense: ")
	transport := readAmount(reader, "Enter transport expense: ")
	other := readAmount(reader, "Enter other expenses: ")

	//==========================================
	// 3. Calculate totals (rounded to cents)
	//==========================================
	totalExpenses := math.Round((rent+food+transport+other)*100) / 100
	remaining := math.Round((income-totalExpenses)*100) / 100

	//==========================================
	// 4. Show the results
	//==========================================
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
