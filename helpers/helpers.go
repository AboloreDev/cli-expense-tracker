package helpers

import (
	"cli-expense-tracker/storage"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Adding expenses
// Updating Expenses
// Deleting expense
// List all expenses

func AddExpense() {
	// Create a falg that parses command line arguments
	// flags are like --descriptions etc
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// define what flags to expect
	// flag.string returns a pointer to the string value
	desciption := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0, "Expense Amount")
	category := addCmd.String("category", "Genenral", "Expense category")

	// Parse the flags from the command line
	addCmd.Parse(os.Args[2:])

	// access the value of the description using pointer
	if *desciption == "" {
		fmt.Println("Error: Description is required")
		fmt.Println("Usage: expense-tracker add --description <desc> --amount <amount>")
		os.Exit(1)
	}

	// Access the value of amount using pointer
	if *amount == 0 {
		fmt.Println("Error: Amount is required")
		fmt.Println("Usage: expense-tracker add --description <desc> --amount <amount>")
		os.Exit(1)
	}

	// Category is set to default "Genenral" but you can set a new category

	// Load expenses
	expenseList, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Error loading response: %v\n", err)
	}

	// Add a new expense using the go struct
	newExpense := storage.Expense{
		ID:          storage.GenNextID(expenseList),
		Description: *desciption,
		Amount:      *amount,
		Category:    *category,
		Date:        time.Now(),
	}

	// Append the new expense to the expense list
	expenseList.Expenses = append(expenseList.Expenses, newExpense)

	// Save the expense
	err = storage.SaveExpense(expenseList)
	if err != nil {
		fmt.Printf("Failed to save expense: %v\n", err)
	}

	// Print a message if successfull
	fmt.Println("Expense successfully added", newExpense.ID)
}

func UpdateExpense() {
	// Create a comand flag
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	// find the expense to update using the Id
	id := updateCmd.Int("id", 0, "Expense Id")
	desciption := updateCmd.String("description", "", "Expense description")
	amount := updateCmd.Float64("amount", 0, "Expense Amount")
	category := updateCmd.String("category", "Genenral", "Expense category")

	//Parse the flag command
	updateCmd.Parse(os.Args[2:])

	if *id == 0 {
		fmt.Println("Error: Id is required")
		fmt.Println("Usage: expense-tracker update --id <id> [--description <desc>] [--amount <amount>]")
		os.Exit(1)
	}

	// Load expenses
	expenseList, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Failed to load expenses: %v", err)
		os.Exit(1)
	}

	// Loop over each expense and update if it matches the id
	found := false
	for i := range expenseList.Expenses {
		if expenseList.Expenses[i].ID == *id {
			// Upate only the field thats required
			if *desciption == "" {
				expenseList.Expenses[i].Description = *desciption
			}
			if *amount > 0 {
				expenseList.Expenses[i].Amount = *amount
			}
			if *category != "" {
				expenseList.Expenses[i].Category = *category
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Error: Expense with ID %d not found\n", *id)
		os.Exit(1)
	}

	err = storage.SaveExpense(expenseList)
	if err != nil {
		fmt.Printf("Failed to save expense: %v\n", err)
	}

	// Print on success
	fmt.Println("Sucecssfully updated the expenses")
}

func DeleteExpense() {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	id := deleteCmd.Int("ID", 0, "Expense ID")

	if *id == 0 {
		fmt.Println("Error: Id is required")
		fmt.Println("Usage: expense-tracker delete --id <id>")
		os.Exit(1)
	}

	expenseList, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Failed to load expense: %v", err)
		os.Exit(1)
	}

	// To delete we need to create a copy
	found := false
	var newExpense = []storage.Expense{}

	for _, expense := range expenseList.Expenses {
		if expense.ID == *id {
			found = true
			continue
		}
		newExpense = append(newExpense, expense)
	}

	if !found {
		fmt.Printf("Error: Expense with ID %d not found\n", *id)
		os.Exit(1)
	}

	expenseList.Expenses = newExpense

	err = storage.SaveExpense(expenseList)
	if err != nil {
		fmt.Printf("Failed to save expense: %v\n", err)
	}

	// Print on success
	fmt.Println("Sucecssfully deleted the expenses")

}

// handleList processes the "list" command
// Like reading all expense cards in the diary
func ListExpenses() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// use the category to list
	category := listCmd.String("category", "", "Filter by Category")

	listCmd.Parse(os.Args[2:])

	// load expense
	expenseList, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		os.Exit(1)
	}

	if len(expenseList.Expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	fmt.Printf("%-5s %-12s %-20s %-15s %-10s\n", "ID", "Date", "Description", "Category", "Amount")
	fmt.Println(strings.Repeat("-", 70))

	for _, expense := range expenseList.Expenses {
		// If category filter is set, skip expenses that don't match
		if *category != "" && expense.Category != *category {
			continue
		}

		// Format date as YYYY-MM-DD
		dateStr := expense.Date.Format("2026-01-01")

		// Truncate description if too long
		desc := expense.Description
		if len(desc) > 20 {
			desc = desc[:17] + "..."
		}

		// Print the expense row
		fmt.Printf("%-5d %-12s %-20s %-15s $%-9.2f\n",
			expense.ID,
			dateStr,
			desc,
			expense.Category,
			expense.Amount)
	}
}

// handleSummary processes the "summary" command
// Like adding up all the money you spent for a month
func MonthlyExpenseSummary() {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	// Summary command for filterng based on month
	month := summaryCmd.Int("month", 0, "filter by month (1-12)")

	summaryCmd.Parse(os.Args[2:])

	expenseList, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		os.Exit(1)
	}

	// Create a variable to hold the value of total
	var total float64 = 0
	// get the current year
	currentYear := time.Now().Year()

	// Add up all expense
	for _, expense := range expenseList.Expenses {
		// If month filter is set, only count expenses from that month
		if *month != 0 {
			// Check if the expenses is from the specified month and the current year
			if expense.Date.Month() == time.Month(*month) && expense.Date.Year() == currentYear {
				total += expense.Amount
			} 
		}else {
				// No filter, add all expenses
				total += expense.Amount
			}
	}

	// Print the reuslt
	if *month != 0 {
		monthName := time.Month(*month).String()
		fmt.Printf("Total expenses for %s: $%.2f\n", monthName, total)

		// Check if the month selected has a budget
		budgetList, err := storage.LoadBudget()
		if err != nil {
			fmt.Printf("Error loading budget: %v\n", err)
			os.Exit(1)
		}

		// lOOP OVER EACH BUDGET
		for _, budget := range budgetList.Budgets {
			// Check if the budget of the month equals month passed into command and
			// also the year matches the current year of the budget year
			if budget.Month == *month && budget.Year == currentYear {
				// calculate the remaining if its above
				remaining := budget.Amount - total
				if remaining < 0 {
					fmt.Printf("⚠️ WARNING: You have exceeded your budget by $%.2f!\n", -remaining)
				} else {
					fmt.Printf("✓ Remaining budget: $%.2f\n", remaining)
				}
				break
			}
		}
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}

}
