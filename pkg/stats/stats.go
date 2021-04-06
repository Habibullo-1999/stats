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

func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money {}	

	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount		
	}

	return categories
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money {}	
	catlen := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
		catlen[payment.Category]+=1	
	}

	for key := range categories {
		categories[key]=(categories[key]/catlen[key])
	}

	return categories
}

func PeriodsDynamic( first map[types.Category]types.Money,second map[types.Category]types.Money) map[types.Category]types.Money {
	categoriesPeriod := map[types.Category]types.Money{}
	for key := range first {
		_, ok := second[key]
		if !ok{
			categoriesPeriod[key]=first[key]*-1
			continue
		}

		for key2 := range second {
			_,ok := first[key2]
			if !ok {
				categoriesPeriod[key2]=second[key2]
				continue
			}
		}	
		categoriesPeriod[key]=second[key]-first[key]
	}
	return categoriesPeriod
   }
   
