package helpers

import (
	"cli-expense-tracker/storage"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func ExportExpense() {
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)

	filename := exportCmd.String("file", "expenses.csv", "Export Expenses as CSV file")

	exportCmd.Parse(os.Args[2:])

	expenseList, err := storage.LoadExpenses()
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		os.Exit(1)
	}
	if len(expenseList.Expenses) == 0 {
		fmt.Println("No expenses to export")
		return
	}

	// Create a CSV file
	file, err := os.Create(*filename)
	if err != nil {
		fmt.Printf("Failed to create export file: %v", err)
		os.Exit(1)
	}
	// Close when don
	defer file.Close()

	// Write to the csv file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Set the file structure
	header := []string{"ID", "Date", "Description", "Category", "Amount"}
	if err := writer.Write(header) ; err != nil {
		fmt.Printf("Error writing header: %v\n", err)
		os.Exit(1)
	}

	// Write each expense
	for _, expense := range expenseList.Expenses {
		record := []string{
			strconv.Itoa(expense.ID),
			expense.Date.Format("2006-01-02"),
			expense.Description,
			expense.Category,
			fmt.Sprintf("%.2f", expense.Amount),
		}
		if err := writer.Write(record); err != nil {
			fmt.Printf("Error writing record: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Expenses exported successfully to %s\n", *filename)
}