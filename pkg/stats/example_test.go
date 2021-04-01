package stats

import (
	"fmt"

	"github.com/Habibullo-1999/bank/pkg/types"

)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   4_000,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   5_000,
			Category: "auto",
		},
		{
			ID:       3,
			Amount:   6_000,
			Category: "auto",
		},
	}

	result := Avg(payments)
	fmt.Println(result)

	// Output: 5000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   4_000,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   5_000,
			Category: "food",
		},
		{
			ID:       3,
			Amount:   6_000,
			Category: "auto",
		},
	}

	result := TotalInCategory(payments, "auto")
	fmt.Println(result)

	// Output: 10000
}