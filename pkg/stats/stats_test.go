package stats

import (
	"reflect"
	"testing"
	"github.com/anonimous-arn/bank/v2/pkg/types"
	
)
func TestCategoriesAvg_nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0 ")
	}
}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment {}
	result := CategoriesAvg(payments)

	if len(result) != 0 {
		t.Error("result len != 0 ")
	}
}

func TestCategoriesAvg_one(t *testing.T) {
	payments := []types.Payment{
	  {
		ID:       1,
		Category: "cafe",
		Amount:   100_00,
	  },
	  {
		ID:       2,
		Category: "cafe",
		Amount:   100_00,
	  },
	}
	expected := map[types.Category]types.Money{
	  "cafe": 100_00,
	}
  
	res := CategoriesAvg(payments)
  
	if !reflect.DeepEqual(expected, res) {
	  t.Errorf("\n got > %v want > nil", res)
	}
  }