package stats

import "github.com/Habibullo-1999/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var avgAmount types.Money
	var counter int32=0
	for _, operation := range payments {
		if operation.Status ==types.StatusFail{
			continue
		} 
		counter++
		avgAmount += operation.Amount
	}

	return avgAmount /types.Money(counter)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, operation := range payments {
		if operation.Status ==types.StatusFail{
			continue
		} 
		if category == operation.Category {
			sum += operation.Amount
		}
		
	}

	return sum
}


func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment :=range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered	
}