# Section 10 Refunds

## 125-001 Refunds from the Stripe Dashboard
In stripe dashboard, go to `payments` page, you can choose a payment and click the `refund` button. Doing it this way, but on our own dashboard,
we don't know if a payment is refunded or not.

There are number of ways to fix this.

- You can set up a stripe hook which is just a webhook and anytime a refund or a payment takes place, we fire off a req to your own API and handle it that way
- implement refunds on our own application(we'll do this)

## 126-002 Adding a refund function to our cards package
Anytime we're working with stripe, we need to set the secret key, like in the `Refund` method.

It's more secure that in `/refund` to accept only the orderID and look up in DB and get the paymentIntent, than sending the paymentIntent as part of
payload. But we didn't go with this approach because we want to write a more readable code in this course. So we created some hidden fields to store
the payload data for `/refund` in `sale.page.gohtml`.

## 127-003 Creating an API handler to process refunds

## 128-004 Update the front end for refunds

## 129-005 Improving the front end

## 130-006 Adding UI components to the sales page

## 131-007 Updating status to refunded in the database

---

126 Refund-documentation-on-Stripe
https://stripe.com/docs/refunds

127 The-Stripe-refund-object
https://stripe.com/docs/api/refunds/object

128 SweetAlert
https://sweetalert2.github.io/
