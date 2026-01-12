package main

import (
	"cli-expense-tracker/budget"
	"cli-expense-tracker/helpers"
	"fmt"
	"os"
)


func main() {
	
	// Check if user provided a command
	// os.Args is like a list: [program-name, command, --flag1, value1, ...]
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Get the command (add, list, delete, etc.)
	command := os.Args[1]

	// Handle each command differently
	switch command {
	case "add":
		helpers.AddExpense()
	case "update":
		helpers.UpdateExpense()
	case "delete":
		helpers.DeleteExpense()
	case "list":
		helpers.ListExpenses()
	case "summary":
		helpers.MonthlyExpenseSummary()
	// case "export":
		helpers.ExportExpense()
	case "set-budget":
		budget.SetBudget()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}

}

// printUsage shows how to use the program
// Like an instruction manual for the piggy bank diary
func printUsage() {
	fmt.Println("Expense Tracker CLI")
	fmt.Println("\nUsage:")
	fmt.Println("  expense-tracker add --description <desc> --amount <amount> [--category <category>]")
	fmt.Println("  expense-tracker update --id <id> [--description <desc>] [--amount <amount>] [--category <category>]")
	fmt.Println("  expense-tracker delete --id <id>")
	fmt.Println("  expense-tracker list [--category <category>]")
	fmt.Println("  expense-tracker summary [--month <month>]")
	fmt.Println("  expense-tracker export --file <filename.csv>")
	fmt.Println("  expense-tracker set-budget --month <month> --amount <amount>")
	fmt.Println("\nExamples:")
	fmt.Println("  expense-tracker add --description \"Lunch\" --amount 20")
	fmt.Println("  expense-tracker add --description \"Taxi\" --amount 15 --category \"Transport\"")
	fmt.Println("  expense-tracker list --category \"Food\"")
	fmt.Println("  expense-tracker summary --month 8")
	fmt.Println("  expense-tracker set-budget --month 8 --amount 1000")
}
