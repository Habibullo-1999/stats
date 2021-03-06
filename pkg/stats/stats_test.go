package stats

import (
	"reflect"
	"testing"

	"github.com/Habibullo-1999/bank/v2/pkg/types"

)


func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}
	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}


func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}
	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}


func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}

	expected := map[types.Category]types.Money{
		"auto": 8_000_00,
		"food": 2_000_00,
		"fun":  5_000_00,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}


func TestCategoryAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 2_000_000},
		{ID: 2, Category: "food", Amount: 2_000_000},
		{ID: 3, Category: "auto", Amount: 3_000_000},
		{ID: 4, Category: "auto", Amount: 4_000_000},
		{ID: 5, Category: "fun", Amount: 5_000_000},
	}

	expected := map[types.Category]types.Money{
		"auto":	3_000_000,
		"food":	2_000_000,
		"fun":	5_000_000,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected,result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}


func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money {
		"auto":	1_000_000,
		"food":	2_000_000,
		"fun":	5_000_000,
	}
	second := map[types.Category]types.Money {
		"auto":	2_000_000,
		"food":	3_000_000,
		"fun":	3_000_000,
		"mobile": 1_000_000,
	}

	expected := map[types.Category]types.Money{
		"auto":	1_000_000,
		"food":	1_000_000,
		"fun":	-2_000_000,
		"mobile": 1_000_000,
		
	}

	result := PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected,result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}