package cards

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

type Card struct {
	Secret   string
	Key      string // publishable key
	Currency string
}

type Transaction struct {
	TransactionStatusID int

	/* We're not gonna use float to do our math with transactions as it suggests in stripe docs. */
	Amount   int
	Currency string

	/* The last four digits of whatever credit card transaction went though. Not the entire credit card. We don't ever want that info.*/
	LastFour string

	// whatever we get back from stripe
	BankReturnCode string
}

func (c *Card) Charge(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	return c.CreatePaymentIntent(currency, amount)
}

func (c *Card) CreatePaymentIntent(currency string, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = c.Secret

	// create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(currency),
	}

	// for reference:
	// params.AddMetadata("key", "value")

	// pi stands for payment intent
	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}

		return nil, msg, err
	}

	return pi, "", nil
}

// GetPaymentMethod gets the payment method by payment intent id
func (c *Card) GetPaymentMethod(s string) (*stripe.PaymentMethod, error) {
	stripe.Key = c.Secret

	// pm stands for payment method
	pm, err := paymentmethod.Get(s, nil)
	if err != nil {
		return nil, err
	}

	return pm, nil
}

// RetrievePaymentIntent gets an existing payment intent by id
func (c *Card) RetrievePaymentIntent(id string) (*stripe.PaymentIntent, error) {
	stripe.Key = c.Secret

	// pi stands for payment intent
	pi, err := paymentintent.Get(id, nil)
	if err != nil {
		return nil, err
	}

	return pi, nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg = ""

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your card is expired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect zip/postal code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount is too large to charge your card"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is too small to charge your card"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient balance"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Your postal code is invalid"
	default:
		msg = "Your card was declined"
	}

	return msg
}
