package processtransaction

import "testing"

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {

	input := TransactionDtoInput{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "40000000000000000",
		CreditCardName:            "Renato Alves",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             123,
		Amount:                    200,
	}

	expectedOutPut := TransactionDtoOutPut{
		ID:           "1",
		Status:       "rejected",
		ErrorMessage: "",
	}

}
