package stats

import (
	"github.com/anonimous-arn/bank/v2/pkg/types"
)



func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
    sum := types.Money(0)
    for _, payment := range payments {
		if payment.Category == category {
            if payment.Status != types.StatusFail {
                sum += payment.Amount
            }
        }
	
    }
    return sum
}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
    moneyInCategory := map[types.Category]types.Money {}
    amountInCategory := map[types.Category]types.Money {}

    for _, payment := range payments {
        moneyInCategory [payment.Category] += payment.Amount
        amountInCategory [payment.Category] += 1
    }

    for key := range moneyInCategory {
        moneyInCategory [key] /= amountInCategory [key]
    }

    return moneyInCategory
}