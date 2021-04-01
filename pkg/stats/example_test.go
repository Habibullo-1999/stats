package stats

import (
	"fmt"

	"github.com/Habibullo-1999/bank/v2/pkg/types"

)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   4_000,
			Category: "auto",
			Status: types.StatusOk,
		},
		{
			ID:       2,
			Amount:   5_000,
			Category: "auto",
			Status: types.StatusOk,
		},
		{
			ID:       3,
			Amount:   6_000,
			Category: "auto",
			Status: types.StatusFail,
		},
	}

	result := Avg(payments)
	fmt.Println(result)

	// Output: 4500

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   4_000,
			Category: "auto",
			Status: types.StatusOk,
		},
		{
			ID:       2,
			Amount:   5_000,
			Category: "food",
			Status: types.StatusOk,
		},
		{
			ID:       3,
			Amount:   6_000,
			Category: "auto",
			Status: types.StatusFail,
		},
	}

	result := TotalInCategory(payments, "auto")
	fmt.Println(result)

	// Output: 4000
}