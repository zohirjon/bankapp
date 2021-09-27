package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	var max = payments[0].Amount
	var idmaxpayment int

	for id, payment := range payments {
		if max < payment.Amount {
			idmaxpayment = id
			max = payment.Amount
		}
	}

	return payments[idmaxpayment]
}
