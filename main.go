package main

import (
	"fmt"
	m2 "personal-budget/module2"
	"time"
)

func main() {
	bu, _ := m2.CreateBudget(time.January, 1000)
	bu.AddItem("bananas", 10.0)
	fmt.Printf("Current cost for January: $%.2f \n", bu.CurrentCost())

	bu2, _ := m2.CreateBudget(time.February, 1000)
	bu2.AddItem("bananas", 10.0)
	bu2.AddItem("coffee", 3.99)
	bu2.AddItem("gym", 50.0)
	bu2.RemoveItem("coffee")
	fmt.Printf("Current cost for February: $%.2f \n", bu2.CurrentCost())

	fmt.Println("Resetting Budget Report...")
	m2.InitializeReport()

	var months = []time.Month{
		time.January,
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}

	for _, month := range months {
		_, err := m2.CreateBudget(month, 100.00)
		if err == nil {
			fmt.Println("Budget created for", month)

			fmt.Print("Budget for: ", month)
			fmt.Println(m2.GetBudget(month))
		} else {
			fmt.Println("Error creating budget:", err)
		}
	}

	_, err := m2.CreateBudget(time.December, 100.00)
	fmt.Println(err)
}
