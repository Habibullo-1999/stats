package stats

import "github.com/Habibullo-1999/bank/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var avgAmount types.Money

	for _, operation := range payments {
		avgAmount += operation.Amount
	}

	return avgAmount / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, operation := range payments {
		if category == operation.Category {
			sum += operation.Amount
		}
	}

	return sum
}