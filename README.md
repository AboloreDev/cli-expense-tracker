# cli-expense-tracker
Expense Tracker CLI (Go)

A clean, production-style Command Line Expense Tracker built with Go that helps users record expenses, manage monthly budgets, generate summaries, and export data — all from the terminal.

This project focuses on real-world backend fundamentals, not just features.

✨ Features

➕ Add, update, delete expenses

📋 List expenses (with category filtering)

📊 Monthly & overall spending summaries

💰 Monthly budget tracking with overspend warnings

📁 Export expenses to CSV (Excel-friendly)

🕒 Timestamped records with automatic IDs

💾 Persistent storage using JSON files

🧠 Core Concepts Demonstrated

CLI argument parsing using flag

File persistence with JSON (encoding/json)

CSV generation (encoding/csv)

Struct-based data modeling

Pointer usage for shared state (*ExpenseList, *Budget)

Time-based filtering (month/year logic)

Defensive programming & error handling

Clean separation of concerns (handlers, storage, logic)

This project is intentionally designed to resemble real backend workflows, not toy examples.

🚀 Usage Examples
expense-tracker add --description "Lunch" --amount 20 --category Food
expense-tracker list --category Food
expense-tracker summary --month 8
expense-tracker set-budget --month 8 --amount 1000
expense-tracker export --file expenses.csv

📂 Data Storage

expenses.json → stores all expense records

config.json → stores monthly budgets

Both are human-readable and versionable.

🛠 Tech Stack

Go (Golang)

Standard library only (no external dependencies)

https://roadmap.sh/projects/expense-tracker
