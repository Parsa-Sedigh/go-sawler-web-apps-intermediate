package main

import (
	"github.com/Parsa-Sedigh/go-sawler-web-apps-intermediate/internal/models"
	"net/http"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	//stringMap := make(map[string]string)
	//stringMap["publishable_key"] = app.config.stripe.key

	//if err := app.renderTemplate(w, r, "terminal", &templateData{
	//	StringMap: stringMap,
	//}, "stripe-js"); err != nil {
	//	app.errorLog.Println(err)
	//}
	if err := app.renderTemplate(w, r, "terminal", &templateData{}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)

		// this isn't very polite(we should send back a proper err response, but it's sufficient for our purposes right now)
		return
	}

	// read POSTed data
	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})

	// we could do these on one line
	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		// in prod, do sth else
		app.errorLog.Println(err)
	}
}

// ChargeOnce displays the page to buy one widget
func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	widget := models.Widget{
		ID:             1,
		Name:           "Custom widget",
		Description:    "Great one",
		InventoryLevel: 10,
		Price:          1000,
	}

	data := make(map[string]interface{})
	data["widget"] = widget

	if err := app.renderTemplate(w, r, "buy-once", &templateData{
		Data: data,
	}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
		return
	}
}
