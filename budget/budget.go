package budget

import (
	"cli-expense-tracker/storage"
	"flag"
	"fmt"
	"os"
	"time"
)

// Set a monthly budget
func SetBudget() {
	budgetCmd := flag.NewFlagSet("set-budget", flag.ExitOnError)

	month := budgetCmd.Int("month", 0, "Monthly Budget")
	amount := budgetCmd.Float64("amount", 0, "Budget Amount")

	budgetCmd.Parse(os.Args[2:])

	if *month < 1 || *month > 12 {
		fmt.Println("Error: Month must be between 1 and 12")
		os.Exit(1)
	}

	if *amount <= 0 {
		fmt.Println("Error: Amount must be greater than 0")
		os.Exit(1)
	}

	// Load budget
	budgetList, err := storage.LoadBudget()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Current Year
	currentYear := time.Now().Year()

	// Update existing budget or add new budget
	found := false 
	for i := range budgetList.Budgets {
		if budgetList.Budgets[i].Month == *month && budgetList.Budgets[i].Year == currentYear {
			budgetList.Budgets[i].Amount = *amount
			found = true
			break
		}
	}

	// If not found create
	if !found {
		newBudget := storage.Budget{
			Amount: *amount,
			Month: *month,
			Year: currentYear,
		}

		budgetList.Budgets = append(budgetList.Budgets, newBudget)
	}

	// Save budget
	err = storage.SaveBudget(budgetList)
	if err != nil {
		fmt.Printf("Failed to save budget: %v", err)
		os.Exit(1)
	}

	// Print on success
	monthName := time.Month(*month).String()
	fmt.Printf("Budget for %s set to $%.2f\n", monthName, *amount)
}

