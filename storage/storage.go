package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Load Expense
// Load Budget
// Generate ID
// Save budget
// Load Budget
// Define Struct
// Expense struct
type Expense struct {
	ID          int     `json:"id"`
	Date        time.Time  `json:"time"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// A struct that holds a list of all expense
type ExpenseList struct {
	Expenses []Expense
}

// Budget struct
type Budget struct {
	Month int `json:"month"`
	Year int `json:"year"`
	Amount float64 `json:"amount"`
}

// A struct that holds a list of all budget
type BudgetList struct {
	Budgets []Budget
}


const expenseFilename = "expenses.json"
const budgetFilename = "budget.json"

// Load expense
func LoadExpenses() (*ExpenseList, error) {
	// Create an empty slice of expense
	expenseList := &ExpenseList{Expenses: []Expense{}}

	// Check if the file exists
	if _, err := os.Stat(expenseFilename); os.IsNotExist(err) {
		return expenseList, nil
	}

	// Read the file
	data, err := os.ReadFile(expenseFilename)
	if err != nil {
		return nil, fmt.Errorf("error reading expenses file: %v", err)
	}

	// Check for the data length
	if len(data) == 0 {
		return expenseList, nil
	}

	// Unmarshal the json file into go struct
	err = json.Unmarshal(data, expenseList)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse json %s", err)
	}

	return expenseList, nil
}

func SaveExpense(expense *ExpenseList) error {
	data, err := json.MarshalIndent(expense, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to load json file: %s", err)
	}

	err = os.WriteFile(expenseFilename, data, 0644)
	if err != nil {
		return fmt.Errorf("Unabel to write to file path: %s", err)
	}

	return nil
}

// Load budget
func LoadBudget() (*BudgetList, error) {
	budgetList := &BudgetList{Budgets: []Budget{}}

	// Check if file exisT
	if _, err := os.Stat(budgetFilename); os.IsNotExist(err) {
		return  budgetList, nil
	}

	// read file
	data, err := os.ReadFile(budgetFilename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file, please check file: %s", err)
	}

	// check the file length
	if len(data) == 0 {
		return  budgetList, nil
	}

	// Unmarshal from json into go struct
	err = json.Unmarshal(data, budgetList) 
	if err != nil {
		return nil, fmt.Errorf("Failed to parse json file: %v", err)
	}

	return budgetList, nil
}

func SaveBudget(budget *BudgetList) error{
	data, err := json.MarshalIndent(budget, "", " ")
	if err != nil {
		return  fmt.Errorf("Failed to unparse data: %v", err)
	}

	// Write to file
	err = os.WriteFile(budgetFilename, data, 0644) 
	if err != nil {
		return  fmt.Errorf("Failed to write file: %v", err)
	}

	return nil
}

// Generate next ID for expense
func GenNextID (expenseList *ExpenseList)  int {
	maxId := 0

	for _, expense := range expenseList.Expenses {
		if expense.ID > maxId {
			maxId = expense.ID
		}
	}
	return  maxId + 1
}


// Check budget
// checkBudget checks if current month's expenses exceed budget
func CheckBudget(expenseList *ExpenseList) {
	// Load all bugets
	budgetList, err := LoadBudget()
	if err != nil {
		return
	}

	now := time.Now()
	currentMonth := int(now.Month())
	currentYear := now.Year()

	// Find budget for current month
	var monthBudget *Budget
	for i := range budgetList.Budgets {
		if budgetList.Budgets[i].Month == currentMonth && budgetList.Budgets[i].Year == currentYear {
			monthBudget = &budgetList.Budgets[i]
			break
		}
	}

	if monthBudget == nil {
		return // No budget set for this month
	}

	// Calculate total for current month
	var total float64
	for _, expense := range expenseList.Expenses {
		if int(expense.Date.Month()) == currentMonth && expense.Date.Year() == currentYear {
			total += expense.Amount
		}
	}

	// Check if over budget
	if total > monthBudget.Amount {
		excess := total - monthBudget.Amount
		fmt.Printf("\n⚠️  WARNING: You have exceeded your %s budget by $%.2f!\n",
			now.Month().String(), excess)
	}
}