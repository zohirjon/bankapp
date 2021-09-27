package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     2,
			Amount: 500,
		},
		{
			ID:     555,
			Amount: 5000,
		},
		{
			ID:     4,
			Amount: 500,
		},
	}

	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 555

}
